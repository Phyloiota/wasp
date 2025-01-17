//nolint:dupl
package test

import (
	"testing"

	"github.com/iotaledger/wasp/contracts/rust/testcore"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/stretchr/testify/require"
)

func TestOffLedgerFailNoAccount(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := deployTestCore(t, w)
		chainAccountBalances(ctx, w, 2, 2)

		user := ctx.NewSoloAgent()
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))

		// no deposit yet, so account is unverified

		f := testcore.ScFuncs.SetInt(ctx.OffLedger(user))
		f.Params.Name().SetValue("ppp")
		f.Params.IntValue().SetValue(314)
		f.Func.Post()
		require.Error(t, ctx.Err)
		require.Contains(t, ctx.Err.Error(), "unverified account")

		t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 2, 2)
	})
}

func TestOffLedgerNoFeeNoTransfer(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := deployTestCore(t, w)
		chainAccountBalances(ctx, w, 2, 2)

		user := ctx.NewSoloAgent()
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))

		deposit(t, ctx, user, nil, 10)
		require.EqualValues(t, solo.Saldo-10, user.Balance())
		require.EqualValues(t, 10, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 2, 2+10)

		// Look, Ma! No .TransferIotas() necessary when doing off-ledger request!
		// we're using setInt() here to be able to verify the state update was done
		f := testcore.ScFuncs.SetInt(ctx.OffLedger(user))
		f.Params.Name().SetValue("ppp")
		f.Params.IntValue().SetValue(314)
		f.Func.Post()
		require.NoError(t, ctx.Err)

		t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
		require.EqualValues(t, solo.Saldo-10, user.Balance())
		require.EqualValues(t, 10, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 2, 2+10)

		// verify state update
		v := testcore.ScFuncs.GetInt(ctx)
		v.Params.Name().SetValue("ppp")
		v.Func.Call()
		require.NoError(t, ctx.Err)
		require.EqualValues(t, 314, v.Results.Values().GetInt64("ppp").Value())
	})
}

func TestOffLedgerFeesEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := deployTestCore(t, w)
		chainAccountBalances(ctx, w, 2, 2)

		user := ctx.NewSoloAgent()
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))

		setOwnerFee(t, ctx, 10)
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3)

		deposit(t, ctx, user, nil, 10)
		require.EqualValues(t, solo.Saldo-10, user.Balance())
		require.EqualValues(t, 10, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3+10)

		// pay enough fees for the request
		nop := testcore.ScFuncs.DoNothing(ctx.OffLedger(user))
		nop.Func.TransferIotas(10).Post()
		require.NoError(t, ctx.Err)

		t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
		require.EqualValues(t, solo.Saldo-10, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3+10, 3+10)
	})
}

func TestOffLedgerFeesNotEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := deployTestCore(t, w)
		chainAccountBalances(ctx, w, 2, 2)

		user := ctx.NewSoloAgent()
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))

		setOwnerFee(t, ctx, 10)
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3)

		deposit(t, ctx, user, nil, 9)
		require.EqualValues(t, solo.Saldo-9, user.Balance())
		require.EqualValues(t, 9, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3+9)

		// try to pay enough fees for the request
		nop := testcore.ScFuncs.DoNothing(ctx.OffLedger(user))
		nop.Func.TransferIotas(10).Post()
		require.Error(t, ctx.Err)
		require.Contains(t, ctx.Err.Error(), "not enough fees")

		t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
		require.EqualValues(t, solo.Saldo-9, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3+9, 3+9)
	})
}

func TestOffLedgerFeesExtra(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := deployTestCore(t, w)
		chainAccountBalances(ctx, w, 2, 2)

		user := ctx.NewSoloAgent()
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))

		setOwnerFee(t, ctx, 10)
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3)

		deposit(t, ctx, user, nil, 11)
		require.EqualValues(t, solo.Saldo-11, user.Balance())
		require.EqualValues(t, 11, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3+11)

		// we have enough fees for the request
		nop := testcore.ScFuncs.DoNothing(ctx.OffLedger(user))
		nop.Func.TransferIotas(10).Post()
		require.NoError(t, ctx.Err)

		t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
		require.EqualValues(t, solo.Saldo-11, user.Balance())
		require.EqualValues(t, 1, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3+10, 3+10+1)
	})
}

func TestOffLedgerTransferWithFeesEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := deployTestCore(t, w)
		chainAccountBalances(ctx, w, 2, 2)

		user := ctx.NewSoloAgent()
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))

		setOwnerFee(t, ctx, 10)
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3)

		deposit(t, ctx, user, nil, 10+42)
		require.EqualValues(t, solo.Saldo-10-42, user.Balance())
		require.EqualValues(t, 10+42, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3+10+42)

		// we have enough fees for the request plus transfer
		nop := testcore.ScFuncs.DoNothing(ctx.OffLedger(user))
		nop.Func.TransferIotas(10 + 42).Post()
		require.NoError(t, ctx.Err)

		t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
		require.EqualValues(t, solo.Saldo-10-42, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 42, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3+10, 3+10+42)
	})
}

func TestOffLedgerTransferWithFeesNotEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := deployTestCore(t, w)
		chainAccountBalances(ctx, w, 2, 2)

		user := ctx.NewSoloAgent()
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))

		setOwnerFee(t, ctx, 10)
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3)

		deposit(t, ctx, user, nil, 10+41)
		require.EqualValues(t, solo.Saldo-10-41, user.Balance())
		require.EqualValues(t, 10+41, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3+10+41)

		// not enough in account for fee + transfer
		nop := testcore.ScFuncs.DoNothing(ctx.OffLedger(user))
		nop.Func.TransferIotas(10 + 42).Post()
		require.NoError(t, ctx.Err)

		t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
		require.EqualValues(t, solo.Saldo-10-41, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 41, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3+10, 3+10+41)
	})
}

func TestOffLedgerTransferWithFeesExtra(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := deployTestCore(t, w)
		chainAccountBalances(ctx, w, 2, 2)

		user := ctx.NewSoloAgent()
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))

		setOwnerFee(t, ctx, 10)
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3)

		deposit(t, ctx, user, nil, 10+43)
		require.EqualValues(t, solo.Saldo-10-43, user.Balance())
		require.EqualValues(t, 10+43, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3, 3+10+43)

		// more than enough in account for fee + transfer
		nop := testcore.ScFuncs.DoNothing(ctx.OffLedger(user))
		nop.Func.TransferIotas(10 + 42).Post()
		require.NoError(t, ctx.Err)

		t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
		require.EqualValues(t, solo.Saldo-10-43, user.Balance())
		require.EqualValues(t, 1, ctx.Balance(user))
		require.EqualValues(t, 42, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 3+10, 3+10+42+1)
	})
}

func TestOffLedgerTransferEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := deployTestCore(t, w)
		chainAccountBalances(ctx, w, 2, 2)

		user := ctx.NewSoloAgent()
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))

		deposit(t, ctx, user, nil, 42)
		require.EqualValues(t, solo.Saldo-42, user.Balance())
		require.EqualValues(t, 42, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 2, 2+42)

		// pay enough fees for the request
		nop := testcore.ScFuncs.DoNothing(ctx.OffLedger(user))
		nop.Func.TransferIotas(42).Post()
		require.NoError(t, ctx.Err)

		t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
		require.EqualValues(t, solo.Saldo-42, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 42, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 2, 2+42)
	})
}

func TestOffLedgerTransferNotEnough(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := deployTestCore(t, w)
		chainAccountBalances(ctx, w, 2, 2)

		user := ctx.NewSoloAgent()
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))

		deposit(t, ctx, user, nil, 41)
		require.EqualValues(t, solo.Saldo-41, user.Balance())
		require.EqualValues(t, 41, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 2, 2+41)

		// pay enough fees for the request
		nop := testcore.ScFuncs.DoNothing(ctx.OffLedger(user))
		nop.Func.TransferIotas(42).Post()
		require.NoError(t, ctx.Err)

		t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
		require.EqualValues(t, solo.Saldo-41, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 41, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 2, 2+41)
	})
}

func TestOffLedgerTransferExtra(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		ctx := deployTestCore(t, w)
		chainAccountBalances(ctx, w, 2, 2)

		user := ctx.NewSoloAgent()
		require.EqualValues(t, solo.Saldo, user.Balance())
		require.EqualValues(t, 0, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))

		deposit(t, ctx, user, nil, 43)
		require.EqualValues(t, solo.Saldo-43, user.Balance())
		require.EqualValues(t, 43, ctx.Balance(user))
		require.EqualValues(t, 0, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 2, 2+43)

		// pay enough fees for the request
		nop := testcore.ScFuncs.DoNothing(ctx.OffLedger(user))
		nop.Func.TransferIotas(42).Post()
		require.NoError(t, ctx.Err)

		t.Logf("dump accounts:\n%s", ctx.Chain.DumpAccounts())
		require.EqualValues(t, solo.Saldo-43, user.Balance())
		require.EqualValues(t, 1, ctx.Balance(user))
		require.EqualValues(t, 42, ctx.Balance(ctx.Account()))
		chainAccountBalances(ctx, w, 2, 2+43)
	})
}

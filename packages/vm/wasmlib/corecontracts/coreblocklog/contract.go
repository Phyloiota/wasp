// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreblocklog

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type ControlAddressesCall struct {
	Func    *wasmlib.ScView
	Results ImmutableControlAddressesResults
}

type GetBlockInfoCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetBlockInfoParams
	Results ImmutableGetBlockInfoResults
}

type GetEventsForBlockCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetEventsForBlockParams
	Results ImmutableGetEventsForBlockResults
}

type GetEventsForContractCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetEventsForContractParams
	Results ImmutableGetEventsForContractResults
}

type GetEventsForRequestCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetEventsForRequestParams
	Results ImmutableGetEventsForRequestResults
}

type GetLatestBlockInfoCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetLatestBlockInfoResults
}

type GetRequestIDsForBlockCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetRequestIDsForBlockParams
	Results ImmutableGetRequestIDsForBlockResults
}

type GetRequestReceiptCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetRequestReceiptParams
	Results ImmutableGetRequestReceiptResults
}

type GetRequestReceiptsForBlockCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetRequestReceiptsForBlockParams
	Results ImmutableGetRequestReceiptsForBlockResults
}

type IsRequestProcessedCall struct {
	Func    *wasmlib.ScView
	Params  MutableIsRequestProcessedParams
	Results ImmutableIsRequestProcessedResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) ControlAddresses(ctx wasmlib.ScViewCallContext) *ControlAddressesCall {
	f := &ControlAddressesCall{Func: wasmlib.NewScView(ctx, HScName, HViewControlAddresses)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) GetBlockInfo(ctx wasmlib.ScViewCallContext) *GetBlockInfoCall {
	f := &GetBlockInfoCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetBlockInfo)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) GetEventsForBlock(ctx wasmlib.ScViewCallContext) *GetEventsForBlockCall {
	f := &GetEventsForBlockCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetEventsForBlock)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) GetEventsForContract(ctx wasmlib.ScViewCallContext) *GetEventsForContractCall {
	f := &GetEventsForContractCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetEventsForContract)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) GetEventsForRequest(ctx wasmlib.ScViewCallContext) *GetEventsForRequestCall {
	f := &GetEventsForRequestCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetEventsForRequest)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) GetLatestBlockInfo(ctx wasmlib.ScViewCallContext) *GetLatestBlockInfoCall {
	f := &GetLatestBlockInfoCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetLatestBlockInfo)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) GetRequestIDsForBlock(ctx wasmlib.ScViewCallContext) *GetRequestIDsForBlockCall {
	f := &GetRequestIDsForBlockCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetRequestIDsForBlock)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) GetRequestReceipt(ctx wasmlib.ScViewCallContext) *GetRequestReceiptCall {
	f := &GetRequestReceiptCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetRequestReceipt)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) GetRequestReceiptsForBlock(ctx wasmlib.ScViewCallContext) *GetRequestReceiptsForBlockCall {
	f := &GetRequestReceiptsForBlockCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetRequestReceiptsForBlock)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) IsRequestProcessed(ctx wasmlib.ScViewCallContext) *IsRequestProcessedCall {
	f := &IsRequestProcessedCall{Func: wasmlib.NewScView(ctx, HScName, HViewIsRequestProcessed)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddView(ViewControlAddresses, wasmlib.ViewError)
	exports.AddView(ViewGetBlockInfo, wasmlib.ViewError)
	exports.AddView(ViewGetEventsForBlock, wasmlib.ViewError)
	exports.AddView(ViewGetEventsForContract, wasmlib.ViewError)
	exports.AddView(ViewGetEventsForRequest, wasmlib.ViewError)
	exports.AddView(ViewGetLatestBlockInfo, wasmlib.ViewError)
	exports.AddView(ViewGetRequestIDsForBlock, wasmlib.ViewError)
	exports.AddView(ViewGetRequestReceipt, wasmlib.ViewError)
	exports.AddView(ViewGetRequestReceiptsForBlock, wasmlib.ViewError)
	exports.AddView(ViewIsRequestProcessed, wasmlib.ViewError)
}

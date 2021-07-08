// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

package test

import "github.com/iotaledger/wasp/packages/coretypes"

const (
	ScName        = "erc20"
	ScDescription = "ERC-20 PoC for IOTA Smart Contracts"
	HScName       = coretypes.Hname(0x200e3733)
)

const (
	ParamAccount    = "ac"
	ParamAmount     = "am"
	ParamCreator    = "c"
	ParamDelegation = "d"
	ParamRecipient  = "r"
	ParamSupply     = "s"
)

const (
	ResultAmount = "am"
	ResultSupply = "s"
)

const (
	StateAllAllowances = "a"
	StateBalances      = "b"
	StateSupply        = "s"
)

const (
	FuncApprove      = "approve"
	FuncInit         = "init"
	FuncTransfer     = "transfer"
	FuncTransferFrom = "transferFrom"
	ViewAllowance    = "allowance"
	ViewBalanceOf    = "balanceOf"
	ViewTotalSupply  = "totalSupply"
)

const (
	HFuncApprove      = coretypes.Hname(0xa0661268)
	HFuncInit         = coretypes.Hname(0x1f44d644)
	HFuncTransfer     = coretypes.Hname(0xa15da184)
	HFuncTransferFrom = coretypes.Hname(0xd5e0a602)
	HViewAllowance    = coretypes.Hname(0x5e16006a)
	HViewBalanceOf    = coretypes.Hname(0x67ef8df4)
	HViewTotalSupply  = coretypes.Hname(0x9505e6ca)
)
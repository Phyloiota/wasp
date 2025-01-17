// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package inccounter

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type CallIncrementCall struct {
	Func *wasmlib.ScFunc
}

type CallIncrementRecurse5xCall struct {
	Func *wasmlib.ScFunc
}

type EndlessLoopCall struct {
	Func *wasmlib.ScFunc
}

type IncrementCall struct {
	Func *wasmlib.ScFunc
}

type IncrementWithDelayCall struct {
	Func   *wasmlib.ScFunc
	Params MutableIncrementWithDelayParams
}

type InitCall struct {
	Func   *wasmlib.ScInitFunc
	Params MutableInitParams
}

type LocalStateInternalCallCall struct {
	Func *wasmlib.ScFunc
}

type LocalStatePostCall struct {
	Func *wasmlib.ScFunc
}

type LocalStateSandboxCallCall struct {
	Func *wasmlib.ScFunc
}

type PostIncrementCall struct {
	Func *wasmlib.ScFunc
}

type RepeatManyCall struct {
	Func   *wasmlib.ScFunc
	Params MutableRepeatManyParams
}

type TestLeb128Call struct {
	Func *wasmlib.ScFunc
}

type WhenMustIncrementCall struct {
	Func   *wasmlib.ScFunc
	Params MutableWhenMustIncrementParams
}

type GetCounterCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetCounterResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) CallIncrement(ctx wasmlib.ScFuncCallContext) *CallIncrementCall {
	return &CallIncrementCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncCallIncrement)}
}

func (sc Funcs) CallIncrementRecurse5x(ctx wasmlib.ScFuncCallContext) *CallIncrementRecurse5xCall {
	return &CallIncrementRecurse5xCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncCallIncrementRecurse5x)}
}

func (sc Funcs) EndlessLoop(ctx wasmlib.ScFuncCallContext) *EndlessLoopCall {
	return &EndlessLoopCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncEndlessLoop)}
}

func (sc Funcs) Increment(ctx wasmlib.ScFuncCallContext) *IncrementCall {
	return &IncrementCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncIncrement)}
}

func (sc Funcs) IncrementWithDelay(ctx wasmlib.ScFuncCallContext) *IncrementWithDelayCall {
	f := &IncrementWithDelayCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncIncrementWithDelay)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) Init(ctx wasmlib.ScFuncCallContext) *InitCall {
	f := &InitCall{Func: wasmlib.NewScInitFunc(ctx, HScName, HFuncInit, keyMap[:], idxMap[:])}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) LocalStateInternalCall(ctx wasmlib.ScFuncCallContext) *LocalStateInternalCallCall {
	return &LocalStateInternalCallCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncLocalStateInternalCall)}
}

func (sc Funcs) LocalStatePost(ctx wasmlib.ScFuncCallContext) *LocalStatePostCall {
	return &LocalStatePostCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncLocalStatePost)}
}

func (sc Funcs) LocalStateSandboxCall(ctx wasmlib.ScFuncCallContext) *LocalStateSandboxCallCall {
	return &LocalStateSandboxCallCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncLocalStateSandboxCall)}
}

func (sc Funcs) PostIncrement(ctx wasmlib.ScFuncCallContext) *PostIncrementCall {
	return &PostIncrementCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncPostIncrement)}
}

func (sc Funcs) RepeatMany(ctx wasmlib.ScFuncCallContext) *RepeatManyCall {
	f := &RepeatManyCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncRepeatMany)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) TestLeb128(ctx wasmlib.ScFuncCallContext) *TestLeb128Call {
	return &TestLeb128Call{Func: wasmlib.NewScFunc(ctx, HScName, HFuncTestLeb128)}
}

func (sc Funcs) WhenMustIncrement(ctx wasmlib.ScFuncCallContext) *WhenMustIncrementCall {
	f := &WhenMustIncrementCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncWhenMustIncrement)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) GetCounter(ctx wasmlib.ScViewCallContext) *GetCounterCall {
	f := &GetCounterCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetCounter)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

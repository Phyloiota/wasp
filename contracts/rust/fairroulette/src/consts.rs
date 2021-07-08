// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

//@formatter:off

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME:  &str = "fairroulette";
pub const HSC_NAME: ScHname = ScHname(0xdf79d138);

pub const PARAM_NUMBER:      &str = "number";
pub const PARAM_PLAY_PERIOD: &str = "playPeriod";

pub const RESULT_LAST_WINNING_NUMBER: &str = "lastWinningNumber";

pub const STATE_BETS:                &str = "bets";
pub const STATE_LAST_WINNING_NUMBER: &str = "lastWinningNumber";
pub const STATE_LOCKED_BETS:         &str = "lockedBets";
pub const STATE_PLAY_PERIOD:         &str = "playPeriod";

pub const FUNC_LOCK_BETS:           &str = "lockBets";
pub const FUNC_PAY_WINNERS:         &str = "payWinners";
pub const FUNC_PLACE_BET:           &str = "placeBet";
pub const FUNC_PLAY_PERIOD:         &str = "playPeriod";
pub const VIEW_LAST_WINNING_NUMBER: &str = "lastWinningNumber";

pub const HFUNC_LOCK_BETS:           ScHname = ScHname(0xe163b43c);
pub const HFUNC_PAY_WINNERS:         ScHname = ScHname(0xfb2b0144);
pub const HFUNC_PLACE_BET:           ScHname = ScHname(0xdfba7d1b);
pub const HFUNC_PLAY_PERIOD:         ScHname = ScHname(0xcb94b293);
pub const HVIEW_LAST_WINNING_NUMBER: ScHname = ScHname(0x2f5f09fe);

//@formatter:on
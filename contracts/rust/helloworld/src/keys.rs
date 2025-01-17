// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// @formatter:off

#![allow(dead_code)]

use wasmlib::*;

use crate::*;

pub(crate) const IDX_RESULT_HELLO_WORLD: usize = 0;

pub const KEY_MAP_LEN: usize = 1;

pub const KEY_MAP: [&str; KEY_MAP_LEN] = [
    RESULT_HELLO_WORLD,
];

pub static mut IDX_MAP: [Key32; KEY_MAP_LEN] = [Key32(0); KEY_MAP_LEN];

pub fn idx_map(idx: usize) -> Key32 {
    unsafe {
        IDX_MAP[idx]
    }
}

// @formatter:on

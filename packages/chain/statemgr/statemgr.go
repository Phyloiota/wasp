// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// statemgr package implements object which is responsible for the smart contract
// ledger state to be synchronized and validated
package statemgr

import (
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	txstream "github.com/iotaledger/goshimmer/packages/txstream/client"
	"time"

	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/state"
)

type stateManager struct {
	chain                              chain.Chain
	peers                              chain.PeerGroupProvider
	nodeConn                           *txstream.Client
	pingPong                           []bool
	deadlineForPongQuorum              time.Time
	pullStateDeadline                  time.Time
	blockCandidates                    map[hashing.HashValue]*pendingBlock
	solidState                         state.VirtualState
	stateOutput                        *ledgerstate.AliasOutput
	stateOutputTimestamp               time.Time
	syncingBlocks                      map[uint32]*syncingBlock
	consensusNotifiedOnStateTransition bool
	log                                *logger.Logger

	// Channels for accepting external events.
	eventStateIndexPingPongMsgCh chan *chain.BlockIndexPingPongMsg
	eventGetBlockMsgCh           chan *chain.GetBlockMsg
	eventBlockHeaderMsgCh        chan *chain.BlockHeaderMsg
	eventStateUpdateMsgCh        chan *chain.StateUpdateMsg
	eventStateOutputMsgCh        chan *chain.StateMsg
	eventPendingBlockMsgCh       chan chain.BlockCandidateMsg
	eventTimerMsgCh              chan chain.TimerTick
	closeCh                      chan bool
}

const (
	pullStateTimeout          = 5 * time.Second
	periodBetweenSyncMessages = 1 * time.Second
)

type syncingBlock struct {
	msgCounter    uint16
	stateUpdates  []state.StateUpdate
	stateOutputID ledgerstate.OutputID
	pullDeadline  time.Time
	block         state.Block
}

type pendingBlock struct {
	// block of state updates, not validated yet
	block state.Block
	// resulting variable state after applied the block to the solidState
	nextState state.VirtualState
}

func New(c chain.Chain, peers chain.PeerGroupProvider, nodeconn *txstream.Client, log *logger.Logger) chain.StateManager {
	ret := &stateManager{
		chain:                        c,
		nodeConn:                     nodeconn,
		syncingBlocks:                make(map[uint32]*syncingBlock),
		blockCandidates:              make(map[hashing.HashValue]*pendingBlock),
		log:                          log.Named("s"),
		eventStateIndexPingPongMsgCh: make(chan *chain.BlockIndexPingPongMsg),
		eventGetBlockMsgCh:           make(chan *chain.GetBlockMsg),
		eventBlockHeaderMsgCh:        make(chan *chain.BlockHeaderMsg),
		eventStateUpdateMsgCh:        make(chan *chain.StateUpdateMsg),
		eventStateOutputMsgCh:        make(chan *chain.StateMsg),
		eventPendingBlockMsgCh:       make(chan chain.BlockCandidateMsg),
		eventTimerMsgCh:              make(chan chain.TimerTick),
		closeCh:                      make(chan bool),
	}
	ret.SetPeers(peers)
	go ret.initLoadState()

	return ret
}

func (sm *stateManager) SetPeers(p chain.PeerGroupProvider) {
	if p != nil {
		sm.log.Debugf("SetPeers: num = %d", p.NumPeers())
	}
	sm.peers = p
	sm.pingPong = make([]bool, p.NumPeers())
}

func (sm *stateManager) Close() {
	close(sm.closeCh)
}

// initial loading of the solid state
func (sm *stateManager) initLoadState() {
	var err error
	var batch state.Block
	var stateExists bool

	sm.solidState, batch, stateExists, err = state.LoadSolidState(sm.chain.ID())
	if err != nil {
		sm.log.Errorf("initLoadState: %v", err)
		sm.chain.Dismiss()
		return
	}
	if stateExists {
		h := sm.solidState.Hash()
		txh := batch.ApprovingOutputID()
		sm.log.Infof("solid state has been loaded. Block index: $%d, State hash: %s, ancor tx: %s",
			sm.solidState.BlockIndex(), h.String(), txh.String())
	} else {
		sm.solidState = nil
		sm.addBlockCandidate(state.MustNewOriginBlock(ledgerstate.OutputID{}))
		sm.log.Info("solid state does not exist: WAITING FOR THE ORIGIN TRANSACTION")
	}
	sm.recvLoop() // Start to process external events.
}

func (sm *stateManager) recvLoop() {
	for {
		select {
		case msg, ok := <-sm.eventStateIndexPingPongMsgCh:
			if ok {
				sm.eventStateIndexPingPongMsg(msg)
			}
		case msg, ok := <-sm.eventGetBlockMsgCh:
			if ok {
				sm.eventGetBlockMsg(msg)
			}
		case msg, ok := <-sm.eventBlockHeaderMsgCh:
			if ok {
				sm.eventBlockHeaderMsg(msg)
			}
		case msg, ok := <-sm.eventStateUpdateMsgCh:
			if ok {
				sm.eventStateUpdateMsg(msg)
			}
		case msg, ok := <-sm.eventStateOutputMsgCh:
			if ok {
				sm.eventStateMsg(msg)
			}
		case msg, ok := <-sm.eventPendingBlockMsgCh:
			if ok {
				sm.eventBlockCandidateMsg(msg)
			}
		case msg, ok := <-sm.eventTimerMsgCh:
			if ok {
				sm.eventTimerMsg(msg)
			}
		case <-sm.closeCh:
			return
		}
	}
}
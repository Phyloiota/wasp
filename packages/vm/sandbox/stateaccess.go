package sandbox

import (
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/state"
)

type stateWrapper struct {
	virtualState state.VirtualState
	stateUpdate  state.StateUpdate
}

func (s *stateWrapper) Codec() kv.Codec {
	return kv.NewCodec(s)
}

func (s *stateWrapper) Get(name kv.Key) ([]byte, error) {
	// FIXME: this is O(N) with N = amount of accumulated mutations
	// it could be improved by caching the latest mutation for evey key
	muts := s.stateUpdate.Mutations()
	for i := muts.Len() - 1; i >= 0; i-- {
		m := muts.At(i)
		if (*m).Key() == name {
			// The key-value pair has been modified during the current request
			// return the latest assigned value
			return (*m).Value(), nil
		}
	}

	// The key-value pair has not been modified
	// Fetch its value from the virtual state
	return s.virtualState.Variables().Get(name)
}

func (s *stateWrapper) Del(name kv.Key) {
	s.stateUpdate.Mutations().Add(kv.NewMutationDel(name))
}

func (s *stateWrapper) Set(name kv.Key, value []byte) {
	s.stateUpdate.Mutations().Add(kv.NewMutationSet(name, value))
}
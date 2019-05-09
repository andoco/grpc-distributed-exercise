package session

import (
	"math/rand"
	"sync"

	"github.com/pkg/errors"
)

type Store interface {
	Add(state *State) error
	Update(state *State) error
	Get(clientId string) (*State, error)
}

type State struct {
	ClientId     string
	MaxNumbers   int
	Rand         *rand.Rand
	NumGenerated int
}

func NewAStore() *AStore {
	return &AStore{
		clientState: make(map[string]State),
		mutex:       sync.RWMutex{},
	}
}

type AStore struct {
	clientState map[string]State
	mutex       sync.RWMutex
}

func (s *AStore) Add(state *State) error {
	s.mutex.Lock()
	s.clientState[state.ClientId] = *state
	s.mutex.Unlock()
	return nil
}

func (s *AStore) Update(state *State) error {
	s.mutex.Lock()
	s.clientState[state.ClientId] = *state
	s.mutex.Unlock()
	return nil
}

func (s *AStore) Get(clientId string) (*State, error) {
	s.mutex.RLock()
	state, ok := s.clientState[clientId]
	s.mutex.RUnlock()
	if !ok {
		return nil, errors.Errorf("state not found for client %q", clientId)
	}
	return &state, nil
}

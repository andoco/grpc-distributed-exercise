package session

import "math/rand"

type Store interface {
	Add(state *State) error
	Update(state *State) error
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
	}
}

type AStore struct {
	clientState map[string]State
}

func (s *AStore) Add(state *State) error {
	s.clientState[state.ClientId] = *state
	return nil
}

func (s *AStore) Update(state *State) error {
	s.clientState[state.ClientId] = *state
	return nil
}

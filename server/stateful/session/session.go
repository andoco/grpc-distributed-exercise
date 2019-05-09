package session

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
	"sync"

	"github.com/pkg/errors"
)

type Store interface {
	Add(state *State) error
	Update(state *State) error
	Get(clientId string) (*State, error)
	Delete(clientId string) error
}

type State struct {
	ClientId     string
	MaxNumbers   int
	Seed         int64
	NumGenerated int
}

func NewAStore() *AStore {
	store := &AStore{
		clientState: make(map[string]State),
		mutex:       sync.RWMutex{},
	}

	if err := store.loadFromFile(); err != nil {
		panic(err)
	}

	return store
}

type AStore struct {
	clientState map[string]State
	mutex       sync.RWMutex
}

func (s *AStore) Add(state *State) error {
	s.mutex.Lock()
	s.clientState[state.ClientId] = *state
	if err := s.saveToFile(); err != nil {
		return err
	}
	s.mutex.Unlock()
	return nil
}

func (s *AStore) Update(state *State) error {
	s.mutex.Lock()
	s.clientState[state.ClientId] = *state
	if err := s.saveToFile(); err != nil {
		return err
	}
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

func (s *AStore) Delete(clientId string) error {
	s.mutex.Lock()
	delete(s.clientState, clientId)
	if err := s.saveToFile(); err != nil {
		return err
	}
	s.mutex.Unlock()
	return nil
}

const dataFileName = "data"

func (s *AStore) saveToFile() error {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)

	// Encoding the map
	err := e.Encode(s.clientState)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(dataFileName, b.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func (s *AStore) loadFromFile() error {
	if _, err := os.Stat(dataFileName); os.IsNotExist(err) {
		return nil
	}

	data, err := ioutil.ReadFile(dataFileName)
	if err != nil {
		return err
	}

	b := bytes.NewBuffer(data)
	d := gob.NewDecoder(b)
	if err := d.Decode(&s.clientState); err != nil {
		return err
	}

	return nil
}

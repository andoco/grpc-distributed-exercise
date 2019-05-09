package session

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

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
	LastActive   time.Time
}

func NewAStore() *AStore {
	store := &AStore{
		clientState: make(map[string]State),
		mutex:       sync.RWMutex{},
	}

	if err := store.loadFromFile(); err != nil {
		panic(err)
	}

	store.startCleanupRoutine()

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

func (s *AStore) startCleanupRoutine() {
	go func() {
		for {
			time.Sleep(1 * time.Second)

			s.mutex.Lock()

			expired := []string{}
			for id, state := range s.clientState {
				age := time.Since(state.LastActive)
				if age >= 30*time.Second {
					log.Printf("Found expired client state for %q\n", id)
					expired = append(expired, id)
				}
			}

			for _, id := range expired {
				log.Printf("Deleting expired client state for %q\n", id)
				delete(s.clientState, id)
			}

			if err := s.saveToFile(); err != nil {
				// Should ideally return this somehow instead of exiting
				log.Fatal(err)
			}

			s.mutex.Unlock()
		}
	}()
}

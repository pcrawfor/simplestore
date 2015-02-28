package simplestore

import (
	"bufio"
	"encoding/gob"
	"os"
)

type entry struct {
	Key   string
	Value interface{}
}

// SimpleStore is a basic key value store which writes to a gob file on local disk on save
type SimpleStore struct {
	entries  []*entry
	filePath string
}

// New instantiates a new store with the given filepath either creating or updating the given file
func New(filePath string) *SimpleStore {
	s := SimpleStore{entries: []*entry{}, filePath: filePath}
	s.loadEntries()
	return &s
}

// Get retrieves the value for a given key or returns nil if the key does not exist
func (s *SimpleStore) Get(key []byte) interface{} {
	for _, v := range s.entries {
		if string(key) == v.Key {
			return v.Value
		}
	}

	return nil
}

// Set sets the value of the given key in the store - only accepts string values
func (s *SimpleStore) Set(key []byte, value interface{}) {
	gob.Register(value)
	for _, v := range s.entries {
		if string(key) == v.Key {
			v.Value = value
			return
		}
	}

	s.entries = append(s.entries, &entry{Key: string(key), Value: value})
}

// Save writes the current state of the store to the file on disk
func (s *SimpleStore) Save() error {
	f := s.getFile()
	w := bufio.NewWriter(f)
	enc := gob.NewEncoder(w)

	err := enc.Encode(&s.entries)
	if err != nil {
		return err
	}
	w.Flush()
	f.Close()
	return nil
}

func (s *SimpleStore) loadEntries() error {
	r := bufio.NewReader(s.getFile())
	dec := gob.NewDecoder(r)
	err := dec.Decode(&s.entries)
	if err != nil {
		if err.Error() == "EOF" {
			s.entries = []*entry{}
			return nil
		}
		return err
	}

	return nil
}

func (s *SimpleStore) getFile() *os.File {
	f, err := os.OpenFile(s.filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err.Error())
	}

	return f
}

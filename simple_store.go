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

// Store is a basic key value store which writes to a gob file on local disk on save
type Store struct {
	filePath string
	entries  []*entry
}

// New instantiates a new store with the given filepath either creating or updating the given file
func New(filePath string, types []interface{}) *Store {
	s := Store{entries: []*entry{}, filePath: filePath}
	if len(types) > 0 {
		s.registerTypes(types)
	}
	s.loadEntries()
	return &s
}

func (s *Store) registerTypes(types []interface{}) {
	for _, t := range types {
		gob.Register(t)
	}
}

// Get retrieves the value for a given key or returns nil if the key does not exist
func (s *Store) Get(key string) interface{} {
	for _, v := range s.entries {
		if key == v.Key {
			return v.Value
		}
	}

	return nil
}

// Set sets the value of the given key in the store - only accepts string values
func (s *Store) Set(key string, value interface{}) {
	for i := range s.entries {
		v := s.entries[i]
		if key == v.Key {
			v.Value = value
			return
		}

	}

	s.entries = append(s.entries, &entry{Key: key, Value: value})
}

// Exists checks the existance of a key in the store and returns a bool
func (s *Store) Exists(key string) bool {
	for i := range s.entries {
		v := s.entries[i]
		if key == v.Key {
			return true
		}
	}
	return false
}

// Remove deletes and entry from the store list
func (s *Store) Remove(key string) {
	for i, v := range s.entries {
		if key == v.Key {
			s.entries = append(s.entries[:i], s.entries[i+1:]...)
		}
	}
}

// Values returns the slice of values currently stored
func (s *Store) Values() []interface{} {
	vals := []interface{}{}
	for _, v := range s.entries {
		vals = append(vals, v.Value)
	}
	return vals
}

// Keys returns the slice of keys currently stored
func (s *Store) Keys() []interface{} {
	keys := []interface{}{}
	for _, v := range s.entries {
		keys = append(keys, v.Key)
	}
	return keys
}

// Save writes the current state of the store to the file on disk
func (s *Store) Save() error {
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

func (s *Store) loadEntries() error {
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

func (s *Store) getFile() *os.File {
	f, err := os.OpenFile(s.filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err.Error())
	}

	return f
}

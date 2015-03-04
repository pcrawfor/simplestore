package simplestore

import (
	"os"
	"testing"
)

func initStore() *Store {
	return New("./test_store.gob")
}

func cleanupStore(t *testing.T) {
	rerr := os.Remove("./test_store.gob")
	if rerr != nil {
		t.Error("Failed to cleanup test store")
	}
}

func TestSimpleStoreSaveString(t *testing.T) {
	store := initStore()

	e := store.Save()
	if e != nil {
		t.Error("Error saving store " + e.Error())
	}

	store.Set("foo", "bar")
	v := store.Get("foo")
	if v != "bar" {
		t.Error("Value not found")
	}

	v = store.Get("bar")
	if v != nil {
		t.Error("Value should be nil")
	}

	e = store.Save()
	if e != nil {
		t.Error("Error saving store after write " + e.Error())
	}

	cleanupStore(t)
}

func TestSimpleStoreSaveOthers(t *testing.T) {
	store := initStore()

	store.Set("big", map[string]int{"hi": 1})
	v := store.Get("big")
	if v == nil {
		t.Error("Value should not be nil")
	}

	m := v.(map[string]int)
	if m["hi"] != 1 {
		t.Error("stored map value is incorrect")
	}

	store.Set("small", []bool{true, false, true})

	v = store.Get("small")
	if v == nil {
		t.Error("Value should not be nil")
	}

	a := v.([]bool)
	if a[1] != false {
		t.Error("Array value incorrect")
	}

	cleanupStore(t)
}

func TestSimpleStoreSaveAndLoad(t *testing.T) {
	store := initStore()

	store.Set("a", map[string]string{"one": "two", "three": "four"})
	e := store.Save()
	if e != nil {
		t.Error("Error saving ", e.Error())
	}

	store2 := initStore()
	v := store2.Get("a")

	if v == nil {
		t.Error("Value should not be nil")
	}

	m := v.(map[string]string)

	if m["three"] != "four" {
		t.Error("Error values loaded in store do not match")
	}

	cleanupStore(t)
}

func TestSimpleStoreAddRemove(t *testing.T) {
	store := initStore()

	store.Set("foo", "test")
	v := store.Get("foo")
	if v == nil {
		t.Error("Value should not be nil")
	}

	store.Remove("foo")
	v = store.Get("foo")
	if v != nil {
		t.Error("Value should be nil")
	}

	cleanupStore(t)
}

package simplestore

import (
	"os"
	"testing"
)

func initStore(types []interface{}) *Store {
	return New("./test_store.gob", types)
}

func cleanupStore(t *testing.T) {
	rerr := os.Remove("./test_store.gob")
	if rerr != nil {
		t.Error("Failed to cleanup test store")
	}
}

func TestSimpleStoreSaveString(t *testing.T) {
	store := initStore(nil)

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
	store := initStore([]interface{}{[]bool{}, map[string]int{}})

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

	store.Set("small", "b")
	v = store.Get("small")
	if v == nil {
		t.Error("Value should not be nil")
	}

	cleanupStore(t)
}

func TestSimpleStoreSaveAndLoad(t *testing.T) {
	store := initStore([]interface{}{map[string]string{}})

	store.Set("a", map[string]string{"one": "two", "three": "four"})
	e := store.Save()
	if e != nil {
		t.Error("Error saving ", e.Error())
	}

	store2 := initStore([]interface{}{map[string]string{}})
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

func TestComplexStoreSaveReload(t *testing.T) {
	store := initStore([]interface{}{[]int{}, map[int]string{}})

	store.Set("one", []int{1, 2, 3})
	store.Set("two", map[int]string{1: "one", 2: "two"})
	store.Save()

	store2 := initStore([]interface{}{[]int{}, map[int]string{}})
	v := store2.Get("one")
	b := v.([]int)
	if b == nil {
		t.Error("Value should not be nil")
	}
	if len(b) != 3 || (len(b) == 2 && b[1] != 2) {
		t.Error("Value is invalid")
	}

	v = store2.Get("two")
	c := v.(map[int]string)
	if c == nil {
		t.Error("Value should not be nil")
	}
	if c[1] != "one" || c[2] != "two" {
		t.Error("Map value is invalid")
	}

	cleanupStore(t)
}

func TestSimpleStoreAddRemove(t *testing.T) {
	store := initStore(nil)

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

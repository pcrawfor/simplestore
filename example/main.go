package main

import (
	"fmt"
	"os"

	"github.com/pcrawfor/simplestore"
)

func main() {

	// create the store
	wd, _ := os.Getwd()
	store := simplestore.New(wd+"/example.gob", []interface{}{map[string]int{}})

	store.Set("foo", "bar")
	err := store.Save()
	if err != nil {
		fmt.Println("Ah an error: ", err)
	}

	val := store.Get("foo")
	if val != nil {
		fmt.Println("Value: ", val)
	}

	store.Set("a", map[string]int{"one": 1, "two": 2, "three": 3})
	mapVal := store.Get("a").(map[string]int)
	fmt.Println("Map: ", mapVal)
	// value for key "a" is in the store in memory but if we don't call Save it won't be written to disk
	store.Save()
}

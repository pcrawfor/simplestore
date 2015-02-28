# SimpleStore

SimpleStore is a dead simple key/value store that writes to disk.

It doesn't do anything fancy and has not yet been vetted as goroutine safe so fair warning to you :)

# Usage

    import (
        "github.com/pcrawfor/simplestore"
        "os"
        "fmt"
    )

    // create the store
    wd, _ := os.Getwd()
    store := simplestore.New(wd + "/hulustore.gob")

    store.Set([]byte("foo"), "bar")
    err := store.Save() // save writes the current contents of our key/value store to disk
    if err != nil {
        fmt.Println("Ah an error: ", err)
    }

    val := store.Get([]byte("foo"))
    if val != nil {
        fmt.Println("Value: ", val.(string))
    }

# Notes

The keys used in the store must be of type []byte but the value can be any type, the type will be registered with gob when the value is set and encoded by the underlying gob encoder appropriately.  When reading the value back out you must assert the correct type for the stored value.

All values are stored as interface{} types when written to disk.
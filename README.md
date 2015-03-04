# SimpleStore

SimpleStore is a dead simple key/value store that writes to disk.

It allows you to store arbitrary objects to disk via a key/value interface, the keys are always strings but the values can be any common type or any complex/custom type.

If you are using complex types you must pass in a slice of representative values for the types you will be using to register them with the store.  This allows the store to encode and decode the values properly.

# Usage

    import (
        "github.com/pcrawfor/simplestore"
        "os"
        "fmt"
    )

    // create the store
    wd, _ := os.Getwd()
    // the second param is a set of custom types to support in the store in this case we are not passing any
    store := simplestore.New(wd + "/hulustore.gob", nil)

    store.Set("foo", "bar")
    err := store.Save() // save writes the current contents of our key/value store to disk
    if err != nil {
        fmt.Println("Ah an error: ", err)
    }

    val := store.Get("foo")
    if val != nil {
        fmt.Println("Value: ", val.(string))
    }

    // ------------------------------------

    // store custom types
    // if we have custom types we need to pass them in an array of values of the given type to register with them on the store
    store := simplestore.New(wd + "/hulustore.gob", []interface{}{map[string]string{}})

    m := map[string]string{
        "foo" : "bar"
    }

    store.Set("more", v)
    lm := store.Get("more").(map[string]string)
    if lm != nil {
        fmt.Println("Map value: ", lm)
    }

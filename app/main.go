package main

import (
	"fmt"
	"plugin"

	"app/storage"
)

func main() {
	pl, err := plugin.Open("../lib/lib.so")
	if err != nil {
		panic(err)
	}

	sm, err := pl.Lookup("NewRecordStorage")
	if err != nil {
		panic(err)
	}

	newRecordStorage := sm.(func() storage.RecordStorage)
	st := newRecordStorage()

	rc, err := st.FindOne(7)
	if err == storage.ErrRecordNotFound {
		panic("NOT FOUND")
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", rc)
}

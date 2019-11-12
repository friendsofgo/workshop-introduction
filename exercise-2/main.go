package main

import (
	"encoding/json"
	"fmt"
	"github.com/friendsofgo/workshop-introduction/exercise-2/cache"
)

type StupidCache interface {
	Get(key string) []byte
	Set(key string, data []byte)
}

func main() {
	doSomething(cache.New())
}

func doSomething(cache StupidCache) {
	kv := map[string]string {
		"Company": "Promofarma",
		"Community": "Friends of Go",
		"Language": "Go",
		"Level": "Introduction",
	}

	fmt.Printf("Inserting some keys into the given stupid cache...\n")
	fmt.Printf("%value \n", kv)

	for key, value := range kv {
		data, _ := json.Marshal(value)
		cache.Set(key, data)
	}

	fmt.Printf("Getting some values from the given stupid cache...\n")
	for key, value := range kv {

		cached := cache.Get(key)
		fmt.Printf("The cached value for the key %s is: %s; and the expected was: %s\n", key, value, string(cached))
	}
}


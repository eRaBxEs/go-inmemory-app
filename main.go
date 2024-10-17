package main

import (
	"fmt"
)

var kv = map[string]interface{}{}

type KeyValueStore interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	Del(key string)
}

type InMemoryStore struct {
	key   string
	value interface{}
}

func (i *InMemoryStore) Set(key string, value interface{}) {
	// implementation

	kv[key] = value

}

func (i *InMemoryStore) Get(key string) (interface{}, bool) {
	// implementation
	value, ok := kv[key]

	return value, ok

}

func (i *InMemoryStore) Del(key string) {
	// implementation
	delete(kv, key)

}

func main() {
	fmt.Println("Hello inMemory app")
}

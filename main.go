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

func main() {
	fmt.Println("Hello inMemory app")
}

package main

import (
	"log"
	"net/http"
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
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/del", delHandler)

	log.Print("Server running on :8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Printf("Server shutting down: %v", err)
	}
}

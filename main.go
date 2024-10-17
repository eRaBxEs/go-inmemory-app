package main

import (
	"encoding/json"
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

func setHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST Allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO: Finish me
	// process the request from the json body
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Ensure that the request contains both key and value
	key, ok := data["key"].(string)
	if !ok || key == "" {
		http.Error(w, "Key must be a string", http.StatusBadRequest)
		return
	}

	value, ok := data["value"]
	if !ok {
		http.Error(w, "Value is missing", http.StatusBadRequest)
		return
	}

	// Store the key-value pair in memory
	store := &InMemoryStore{}
	store.Set(key, value)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Key-value pair set successfully"))

}

// getHandler takes a `key` value as a Query parameter and returns the associated
// value in JSON format. If the key does not exist, it instead returns 404.
//
// Input: /get?key=abc
// Output:
//
//	{
//	   "value": any data type
//	}
func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET Allowed", http.StatusMethodNotAllowed)
		return
	}
	// TODO: Finish me

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is missing", http.StatusBadRequest)
		return
	}

	store := &InMemoryStore{}
	value, ok := store.Get(key)
	if !ok {
		http.Error(w, "It does not exist", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{"value": value}

	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}

// delHandler takes a `key` value as a Query parameter and removes it from the
// KeyValueStore. It should return 200 regardless of it the key existed or not.
//
// Input: /del?key=abc
func delHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET Allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO: Finish me

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is missing", http.StatusBadRequest)
		return
	}

	// if query param key is not empty, go ahead and delete whether or not key exists
	store := &InMemoryStore{}
	store.Del(key)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Key deleted successfully"))

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

package main

import (
	"fmt"
	"sync"
	"time"
)

// KeyValueStore represents the key-value storage system
type KeyValueStore struct {
	mu     sync.RWMutex
	values map[string]ValueWithTTL
}

// ValueWithTTL represents a value with its associated TTL
type ValueWithTTL struct {
	Value     interface{}
	ExpiredAt time.Time
}

// NewKeyValueStore creates a new KeyValueStore instance
func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		values: make(map[string]ValueWithTTL),
	}
}

// Set adds a key-value pair to the store with an optional TTL
// Set adds a key-value pair to the store with an optional TTL
func (kv *KeyValueStore) Set(key string, value interface{}, ttl time.Duration) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	var expiredAt time.Time
	expiredAt = time.Now().Add(ttl)

	if ttl == 0 {
		expiredAt = time.Now().Add(time.Hour * 1000 * 1000)
	}

	kv.values[key] = ValueWithTTL{
		Value:     value,
		ExpiredAt: expiredAt,
	}
}

// Get retrieves the value associated with the given key
func (kv *KeyValueStore) Get(key string) interface{} {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	entry, exists := kv.values[key]
	if !exists || entry.ExpiredAt.Before(time.Now()) {
		return nil // Key not found or expired
	}
	return entry.Value
}

// Delete removes the key-value pair associated with the given key
func (kv *KeyValueStore) Delete(key string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	delete(kv.values, key)
}

func main() {
	// Example usage
	store := NewKeyValueStore()

	// Set key-value pairs with TTL
	store.Set("key1", "value1", time.Second*10)
	store.Set("key2", "value2", 0) // No TTL (will never expire)

	// Get values
	fmt.Println("Get key1:", store.Get("key1")) // Output: value1
	fmt.Println("Get key2:", store.Get("key2")) // Output: value2

	// Wait for some time to simulate expiration
	time.Sleep(time.Second * 15)

	// Get expired value
	fmt.Println("Get key1 after expiration:", store.Get("key1")) // Output: <nil>

	// Delete a key
	store.Delete("key2")
	fmt.Println("Get key2 after deletion:", store.Get("key2")) // Output: <nil>
}

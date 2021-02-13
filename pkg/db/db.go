package db

import "sync"

// MapSliceDB is a simple in Memory DB to store key/values, where values are
// explicitly slices of strings.
type MapSliceDB struct {
	// Maps are not r/w safe - so add lock
	lock sync.RWMutex

	data map[string][]string
}

// NewMapSliceDB creates a memory DB with empty trigrams.
func NewMapSliceDB() *MapSliceDB {
	return &MapSliceDB{
		data: map[string][]string{},
	}
}

// Append will append a given value to the slice of a given key.
// If the key doesn't exist, it will create one.
func (db *MapSliceDB) Append(key, value string) {
	db.lock.Lock()
	defer db.lock.Unlock()

	if vals, ok := db.data[key]; ok {
		db.data[key] = append(vals, value)
	} else {
		db.data[key] = []string{value}
	}
}

// Get will retrieve the values for a given key.
func (db *MapSliceDB) Get(key string) []string {
	db.lock.Lock()
	defer db.lock.Unlock()

	suffixes, ok := db.data[key]
	if !ok {
		return nil
	}
	return suffixes
}

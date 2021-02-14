package db_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tgmendes/trigramerator/pkg/db"
)

func TestSetAndGet(t *testing.T) {
	db := db.NewMapSliceDB()

	db.Append("one two", "three")

	suffixes := db.Get("one two")

	expSuffixes := []string{"three"}
	assert.Equal(t, expSuffixes, suffixes)
}

func TestMultipleSetAndGet(t *testing.T) {
	db := db.NewMapSliceDB()

	db.Append("one two", "three")
	db.Append("one two", "four")
	db.Append("one two", "four")
	db.Append("one two", "three")
	db.Append("one two", "three")
	db.Append("one two", "five")
	db.Append("one two", "four")
	db.Append("one two", "four")

	suffixes := db.Get("one two")

	expSuffixes := []string{"three", "four", "four", "three", "three", "five", "four", "four"}
	assert.Equal(t, expSuffixes, suffixes)
}

func TestGetNonexistant(t *testing.T) {
	db := db.NewMapSliceDB()

	suffixes := db.Get("one four")

	assert.Nil(t, suffixes)
}

func TestConcurrentReadWrite(t *testing.T) {
	var wg sync.WaitGroup
	db := db.NewMapSliceDB()

	for i := 0; i < 100; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()
			db.Append("one two", "three")
		}()

		go func() {
			defer wg.Done()
			_ = db.Get("one two")
		}()
	}

	wg.Wait()

	suffixes := db.Get("one two")

	// small sanity check - all added suffixes are in DB
	assert.Len(t, suffixes, 100)
}

func TestEmptyRandomKey(t *testing.T) {
	db := db.NewMapSliceDB()
	key := db.RandomKey()

	assert.Equal(t, "", key)
}

func TestNonEmptyRandomKey(t *testing.T) {
	db := db.NewMapSliceDB()
	db.Append("one two", "three")
	key := db.RandomKey()

	assert.Equal(t, "one two", key)
}

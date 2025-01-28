package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestCacheOverwrite(t *testing.T) {
	cache := NewCache(time.Second)

	// Add initial value
	cache.Add("test", []byte("first"))

	// Overwrite with new value
	cache.Add("test", []byte("second"))

	// Check if it was updated
	val, ok := cache.Get("test")
	if !ok || string(val) != "second" {
		t.Error("Cache didn't properly overwrite value")
	}
}

func TestGetNonExistent(t *testing.T) {
	cache := NewCache(time.Second)
	val, exists := cache.Get("doesnotexist")
	if exists {
		t.Error("Found a key that shouldn't exist")
	}
	if val != nil {
		t.Error("Got a value that should be nil")
	}
}

func TestReplacingKey(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("myValue", []byte("oldValue"))
	cache.Add("myValue", []byte("newValue"))

	val, ok := cache.Get("myValue")
	if !ok || string(val) != "newValue" {
		t.Error("the value wasn't replaced")
	}
}

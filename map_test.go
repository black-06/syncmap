package syncmap

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntMap(t *testing.T) {
	var m Map[int, int]

	m.Store(1, 2)
	_, ok := m.Load(1)
	assert.True(t, ok, "value should be existed")

	m.Delete(1)
	_, ok = m.Load(1)
	assert.False(t, ok, "value should not be existed")

	r, loaded := m.LoadOrStore(1, 2)
	assert.False(t, loaded, "value should not be loaded")

	lr, loaded := m.LoadOrStore(1, r)
	assert.True(t, loaded, "value should not be loaded")
	assert.Equal(t, r, lr, "loaded value should be the same")

	s, _ := m.LoadOrStore(2, 3)
	kv := map[int]int{1: r, 2: s}
	m.Range(func(key, value int) bool {
		v, ok := kv[key]
		assert.True(t, ok, "keys do not match")
		assert.Equal(t, v, value, "values do not match")
		delete(kv, key)
		return true
	})
}

func TestRequests(t *testing.T) {
	var m Map[string, *http.Request]

	m.Store("r", &http.Request{})
	_, ok := m.Load("r")
	assert.True(t, ok, "value should be existed")

	v, ok := m.LoadAndDelete("r")
	assert.True(t, ok, "value should be existed")
	assert.NotNil(t, v, "value should be existed")

	_, ok = m.Load("r")
	assert.False(t, ok, "value should not be existed")

	r, loaded := m.LoadOrStore("r", &http.Request{})
	assert.False(t, loaded, "value should not be loaded")

	lr, loaded := m.LoadOrStore("r", r)
	assert.True(t, loaded, "value should not be loaded")
	assert.Equal(t, r, lr, "loaded value should be the same")

	s, _ := m.LoadOrStore("s", &http.Request{})
	kv := map[string]*http.Request{"r": r, "s": s}
	m.Range(func(key string, value *http.Request) bool {
		v, ok := kv[key]
		assert.True(t, ok, "keys do not match")
		assert.Equal(t, v, value, "values do not match")
		delete(kv, key)
		return true
	})
}

func TestStringIntChan(t *testing.T) {
	var m Map[string, chan int]

	m.Store("r", make(chan int))
	_, ok := m.Load("r")
	assert.True(t, ok, "value should be existed")

	m.Delete("r")
	_, ok = m.Load("r")
	assert.False(t, ok, "value should not be existed")

	r, loaded := m.LoadOrStore("r", make(chan int))
	assert.False(t, loaded, "value should not be loaded")

	lr, loaded := m.LoadOrStore("r", r)
	assert.True(t, loaded, "value should not be loaded")
	assert.Equal(t, r, lr, "loaded value should be the same")

	s, _ := m.LoadOrStore("s", make(chan int))
	kv := map[string]chan int{"r": r, "s": s}
	m.Range(func(key string, value chan int) bool {
		v, ok := kv[key]
		assert.True(t, ok, "keys do not match")
		assert.Equal(t, v, value, "values do not match")
		delete(kv, key)
		return true
	})
}

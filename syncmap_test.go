package main

import (
	"testing"
)

func TestSyncMap1(t *testing.T) {
	m := NewSyncMap(make(map[string]int))

	m.Put("pig", 10)

	v, _ := m.Get("pig")
	if v != 10 {
		t.Failed()
	}
}

func TestSyncMap2(t *testing.T) {
	m := NewSyncMap(make(map[string]int))
	m.Put("dog", 20)
	if useMap(m) != 20 {
		t.Failed()
	}
}

func TestSyncMap3(t *testing.T) {
	m := NewSyncMap(make(map[string]int))
	m.Put("cow", 30)
	wm := &WrapMap{M: m}
	v, _ := wm.M.Get("cow")
	if v != 20 {
		t.Failed()
	}
}

func useMap(m *SyncMap[string, int]) int {
	v, _ := m.Get("cow")
	return v
}

type WrapMap struct {
	M *SyncMap[string, int]
}

package syncmap

import (
	"sort"
	"strings"
	"testing"
)

func TestSyncMap1(t *testing.T) {
	m := NewSyncMap(make(map[string]int))

	m.Put("pig", 10)
	m.Put("dog", 20)

	v, _ := m.Get("pig")
	if v != 10 {
		t.Log("TestSyncMap1: Get")
		t.Fail()
	}

	m.Delete("pig")
	if m.Len() != 1 {
		t.Log("TestSyncMap1: Delete")
		t.Fail()
	}
}

func TestSyncMap2(t *testing.T) {
	m := NewSyncMap(make(map[string]int))
	m.Put("dog", 20)
	if useMap(m, "dog") != 20 {
		t.Log("TestSyncMap2")
		t.Fail()
	}
}

func TestSyncMap3(t *testing.T) {
	m := NewSyncMap(make(map[string]int))
	m.Put("cow", 30)
	wm := &WrapMap{M: m}
	v, _ := wm.M.Get("cow")
	if v != 30 {
		t.Log("TestSyncMap3")
		t.Fail()
	}
}
func TestSyncMap4(t *testing.T) {
	m := NewSyncMap(make(map[string]int))
	m.Put("pig", 10)
	m.Put("dog", 20)
	m.Put("cow", 30)

	keys := []string{}
	val := 0
	for k, v := range m.ForRange() {
		keys = append(keys, k)
		val += v
	}
	sort.Strings(keys)
	if strings.Join(keys, "-") != "cow-dog-pig" {
		t.Log("TestSyncMap4, error 1:", strings.Join(keys, "-"))
		t.Fail()
	}
	if val != 60 {
		t.Log("TestSyncMap4, error 2")
		t.Fail()
	}

	if m.Len() != 3 {
		t.Log("TestSyncMap4, error .Len()")
		t.Fail()
	}
}

func useMap(m *SyncMap[string, int], kw string) int {
	v, _ := m.Get(kw)
	return v
}

type WrapMap struct {
	M *SyncMap[string, int]
}

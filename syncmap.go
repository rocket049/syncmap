// generics sync map
package syncmap

import (
	"sync"
)

type SyncMap[K comparable, V any] struct {
	locker *sync.Mutex
	m      map[K]V
}

// usage:
//
//	var m SyncMap[int,string] = NewSyncMap( make( map[int]string ) )
//	m := NewSyncMap( make( map[int]string ) )
//
// These 2 work same.
func NewSyncMap[K comparable, V any](m map[K]V) *SyncMap[K, V] {
	return &SyncMap[K, V]{locker: &sync.Mutex{}, m: m}
}

// usage:
//
//	value,ok := m.Get( keyword )
//
// If keyword is not exists, ok will be false
func (s *SyncMap[K, V]) Get(kw K) (V, bool) {
	s.locker.Lock()
	v, ok := s.m[kw]
	s.locker.Unlock()
	return v, ok
}

// Store data key/value
func (s *SyncMap[K, V]) Put(kw K, val V) {
	s.locker.Lock()
	s.m[kw] = val
	s.locker.Unlock()
}

// Delete key data
func (s *SyncMap[K, V]) Delete(kw K) {
	s.locker.Lock()
	delete(s.m, kw)
	s.locker.Unlock()
}

// Clear all data
func (s *SyncMap[K, V]) Clear() {
	s.locker.Lock()
	for k, _ := range s.m {
		delete(s.m, k)
	}
	s.locker.Unlock()
}

// usage:
//
//	for key,value := range m.ForRange()
//		// operate key/value
func (s *SyncMap[K, V]) ForRange() map[K]V {
	res := make(map[K]V)
	s.locker.Lock()
	for k, v := range s.m {
		res[k] = v
	}
	s.locker.Unlock()
	return res
}

func (s *SyncMap[K, V]) Len() int {
	var ret int
	s.locker.Lock()
	ret = len(s.m)
	s.locker.Unlock()
	return ret
}

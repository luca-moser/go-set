package set

import (
	"crypto/md5"
	"fmt"
	"sync"
)

func NewSet(entries ...interface{}) Set {
	set := &genericset{}
	set.data = map[interface{}]interface{}{}
	set.Add(entries...)
	return set
}

type Set interface {
	Add(entries ...interface{}) int
	Remove(entries ...interface{}) int
	Has(entry interface{}) bool
	Length() int
	Clear()
	Entries() []interface{}
	Strings() []string
	Ints() []int
	Floats32() []float32
	Floats64() []float64
	Join(set Set)
}

type genericset struct {
	mu   sync.Mutex
	data map[interface{}]interface{}
}

func (s *genericset) Clear() {
	s.mu.Lock()
	s.data = map[interface{}]interface{}{}
	s.mu.Unlock()
}

func (s *genericset) Length() int {
	return len(s.data)
}

func (s *genericset) Entries() []interface{} {
	entries := []interface{}{}
	s.mu.Lock()
	for _, v := range s.data {
		entries = append(entries, v)
	}
	s.mu.Unlock()
	return entries
}

func (s *genericset) key(entry interface{}) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%T(%#v)", entry, entry))))
}

func (s *genericset) Add(entries ...interface{}) int {
	var added int
	s.mu.Lock()
	for x := range entries {
		key := s.key(entries[x])
		if _, has := s.data[key]; has {
			continue
		}
		s.data[key] = entries[x]
		added++
	}
	s.mu.Unlock()
	return added
}

func (s *genericset) Remove(entries ...interface{}) int {
	var removed int
	s.mu.Lock()
	for x := range entries {
		key := s.key(entries[x])
		if _, has := s.data[key]; !has {
			continue
		}
		delete(s.data, key)
		removed++
	}
	s.mu.Unlock()
	return removed
}

func (s *genericset) Has(entry interface{}) bool {
	s.mu.Lock()
	_, has := s.data[s.key(entry)]
	s.mu.Unlock()
	return has
}

func (s *genericset) Join(set Set) {
	s.mu.Lock()
	for _, entry := range set.Entries() {
		s.data[s.key(entry)] = entry
	}
	s.mu.Unlock()
}

func (s *genericset) Strings() []string {
	values := []string{}
	s.mu.Lock()
	for _, v := range s.data {
		switch x := v.(type) {
		case string:
			values = append(values, x)
		}
	}
	s.mu.Unlock()
	return values
}

func (s *genericset) Ints() []int {
	values := []int{}
	s.mu.Lock()
	for _, v := range s.data {
		switch x := v.(type) {
		case int:
			values = append(values, x)
		}
	}
	s.mu.Unlock()
	return values
}

func (s *genericset) Floats32() []float32 {
	values := []float32{}
	s.mu.Lock()
	for _, v := range s.data {
		switch x := v.(type) {
		case float32:
			values = append(values, x)
		}
	}
	s.mu.Unlock()
	return values
}

func (s *genericset) Floats64() []float64 {
	values := []float64{}
	s.mu.Lock()
	for _, v := range s.data {
		switch x := v.(type) {
		case float64:
			values = append(values, x)
		}
	}
	s.mu.Unlock()
	return values
}

package main

import "sync"

type Set struct {
	sync.Mutex
	m map[interface{}]bool
}

func New() *Set {
	return &Set{
		m: make(map[interface{}]bool),
	}
}

func (s *Set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Remove(item interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *Set) Has(item interface{}) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) List() []interface{}  {
	s.Lock()
	defer s.Unlock()
	list := make([]interface{}, 0)
	for value := range s.m {
		list = append(list, value)
	}
	return list
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]bool{}
}

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

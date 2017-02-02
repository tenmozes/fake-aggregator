package aggregator

import "sync"

var mappings map[string]AggregatorInterface

func init() {
	mappings = map[string]AggregatorInterface{
		"/odd": Odd{},
		"/primes": Prime{},
		"/fibo": Fibo{},
		"/rand": Rand{},
	}
}

type MapperInterface interface {
	Mappings() map[string]AggregatorInterface
}

func NewMapper() MapperInterface {
	m := &mapper{
		mappings: make(map[string]AggregatorInterface, len(mappings)),
	}
	for k, v := range mappings {
		m.add(k, v)
	}
	return m
}

type mapper struct {
	mappings map[string]AggregatorInterface
	m        sync.Mutex
}

func (m *mapper) add(path string, aggregator AggregatorInterface) {
	m.m.Lock()
	m.mappings[path] = aggregator
	m.m.Unlock()
}

func (m *mapper) Mappings() map[string]AggregatorInterface {
	m.m.Lock()
	defer m.m.Unlock()
	return m.mappings
}

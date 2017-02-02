package server

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/tenmozes/fake-aggregator/aggregator"
)

type MockAggregator struct{}

func (MockAggregator) Numbers(_ int) ([]int, error) { return []int{3, 2, 1}, nil }

type MockMapping struct{}

func (MockMapping) Mappings() map[string]aggregator.AggregatorInterface {
	return map[string]aggregator.AggregatorInterface{"/stub": MockAggregator{}}
}

func TestServerAggregator(t *testing.T) {
	s := getAndRunServer(7556)
	resp, err := http.Get(s.baseURL + "/stub")
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Error("Exptected 200 status")
		return
	}
	n := numberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&n); err != nil {
		t.Error(err)
		return
	}
	if len(n.Numbers) != 3 {
		t.Error("Exptected 3 items in result")
	}
	for i, v := range []int{3, 2, 1} {
		if n.Numbers[i] != v {
			t.Errorf("Exptected [3,2,1] got %#v", n.Numbers)
			return
		}
	}
}

func TestServerNumbersEmpty(t *testing.T) {
	s := getAndRunServer(7556)
	resp, err := http.Get(s.baseURL + "/numbers")
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Error("Exptected 200 status")
		return
	}
	n := numberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&n); err != nil {
		t.Error(err)
		return
	}
	if len(n.Numbers) != 0 {
		t.Error("Exptected 0 items in result")
	}
}

func TestServerNumbers(t *testing.T) {
	s := getAndRunServer(7556)
	resp, err := http.Get(s.baseURL + "/numbers?u=http://example.com/stub")
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Error("Exptected 200 status")
		return
	}
	n := numberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&n); err != nil {
		t.Error(err)
		return
	}
	if len(n.Numbers) != 3 {
		t.Error("Exptected 3 items in result")
	}
	for i, v := range []int{1, 2, 3} {
		if n.Numbers[i] != v {
			t.Errorf("Exptected [1,2,3] got %#v", n.Numbers)
			return
		}
	}
}

func TestServerNumbersDelayed(t *testing.T) {
	s := getAndRunSlowServer(7557)
	resp, err := http.Get(s.baseURL + "/numbers?u=http://example.com/stub")
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Error("Exptected 200 status")
		return
	}
	n := numberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&n); err != nil {
		t.Error(err)
		return
	}
	if len(n.Numbers) != 0 {
		t.Error("Exptected 0 items in result")
	}
}

func getAndRunServer(port int) *Server {
	s := NewServer(MockMapping{})
	s.maxDelay = 1
	s.deadline = 1 * time.Minute
	s.randomFactor = 1
	go s.Run(port)
	time.Sleep(100 * time.Millisecond)
	return s
}

func getAndRunSlowServer(port int) *Server {
	s := NewServer(MockMapping{})
	s.maxDelay = 100
	s.deadline = 50 * time.Millisecond
	s.minDelay = 100
	s.randomFactor = 1
	go s.Run(port)
	time.Sleep(100 * time.Millisecond)
	return s
}

package aggregator

import "testing"

type nopAggregator struct{}

func (nopAggregator) Numbers(_ int) ([]int, error) { return nil, nil }

var mp *mapper

func init() {
	mp = NewMapper().(*mapper)
	mp.add("foo", nopAggregator{})
}

func TestMapper1(t *testing.T) {
	t.Parallel()
	mp.add("bar", nopAggregator{})
	if _, ok := mp.Mappings()["bar"]; !ok {
		t.Error("bar not found")
	}
	if _, ok := mp.Mappings()["foo"]; !ok {
		t.Error("foo not found")
	}
}

func TestMapper2(t *testing.T) {
	t.Parallel()
	mp.add("baz", nopAggregator{})
	if _, ok := mp.Mappings()["baz"]; !ok {
		t.Error("baz not found")
	}
	if _, ok := mp.Mappings()["foo"]; !ok {
		t.Error("foo not found")
	}
}

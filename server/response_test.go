package server

import "testing"

func TestNumberResponse(t *testing.T) {
	expected := []int{1, 2, 3, 4}
	n := numberResponse{
		Numbers: []int{1, 1, 2, 2, 2, 3, 4, 4, 4, 4, 4, 4},
	}
	n.unique()

	if len(n.Numbers) != len(expected) {
		t.Errorf("not equal %#v - %#v", expected, n.Numbers)
		return
	}
	for i := 0; i < len(n.Numbers); i++ {
		if n.Numbers[i] != n.Numbers[i] {
			t.Errorf("not equal %#v - %#v", expected, n.Numbers)
			return
		}
	}
}

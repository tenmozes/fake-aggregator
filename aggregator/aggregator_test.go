package aggregator

import "testing"

func TestFibo(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}
	actual, err := (Fibo{}).Numbers(len(expected))
	if err != nil {
		t.Error(err)
		return
	}
	if !equal(expected, actual) {
		t.Errorf("not equal %#v - %#v", expected, actual)
	}
}

func TestOdd(t *testing.T) {
	odd := Odd{}
	if _, err := odd.Numbers(0); err == nil {
		t.Error("expected error")
	}
	expected := []int{1, 3, 5, 7, 9, 11, 13, 15}
	actual, err := odd.Numbers(15)
	if err != nil {
		t.Error(err)
		return
	}
	if !equal(expected, actual) {
		t.Errorf("not equal %#v - %#v", expected, actual)
	}
}

func TestPrime(t *testing.T) {
	primes := Prime{}
	if _, err := primes.Numbers(1); err == nil {
		t.Error("expected error")
	}
	expected := []int{2, 3, 5, 7, 11, 13, 17}
	actual, err := primes.Numbers(18)
	if err != nil {
		t.Error(err)
		return
	}
	if !equal(expected, actual) {
		t.Errorf("not equal %#v - %#v", expected, actual)
	}
}

func TestRand(t *testing.T) {
	r := Rand{}
	if _, err := r.Numbers(0); err == nil {
		t.Error("expected error")
	}
	expected := []int{81, 87, 47, 59, 81}
	actual, err := r.Numbers(5)
	if err != nil {
		t.Error(err)
		return
	}
	if !equal(expected, actual) {
		t.Errorf("not equal %#v - %#v", expected, actual)
	}
}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

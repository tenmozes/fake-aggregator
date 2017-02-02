package aggregator

import (
	"errors"
	"math"
	"math/rand"
	"sort"
)

var ErrInvalidInput = errors.New("Invalid input")

type AggregatorInterface interface {
	Numbers(int) ([]int, error)
}

type Odd struct{}

func (Odd) Numbers(n int) ([]int, error) {
	if n < 1 {
		return nil, ErrInvalidInput
	}
	var data []int
	for i := 1; i <= n; i = i + 2 {
		data = append(data, i)
	}
	return data, nil
}

type Prime struct{}

func (Prime) Numbers(n int) ([]int, error) {
	if n < 2 {
		return nil, ErrInvalidInput
	}
	n++
	sieve := make(map[int]bool)
	for i := 2; i < n; i++ {
		sieve[i] = true
	}
	for i := 2; i <= int(math.Pow(float64(n), 0.5)); i++ {
		if sieve[i] {
			for j := i * i; j < n; j = j + i {
				sieve[j] = false
			}
		}
	}
	primes := make([]int, 0, n)
	for prime := range sieve {
		if sieve[prime] {
			primes = append(primes, prime)
		}
	}
	sort.Ints(primes)
	return primes, nil
}

type Fibo struct{}

func (Fibo) Numbers(n int) ([]int, error) {
	fibo := []int{0, 1}
	if n < len(fibo) {
		return fibo[:n], nil
	}
	for i := len(fibo); i < n; i++ {
		fibo = append(fibo, fibo[i-2]+fibo[i-1])
	}
	return fibo, nil
}

type Rand struct {}

func (Rand) Numbers(n int) ([]int, error) {
	if n < 1 {
		return nil, ErrInvalidInput
	}
	r := make([]int, 0, n)
	for i:=0; i<n; i++ {
		r = append(r, rand.Intn(100))
	}
	return r, nil
}
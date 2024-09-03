package main

import "testing"

func Fib(n int, recursive bool) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		if recursive {
			return Fib(n-1, true) + Fib(n-2, true)
		}

		a, b := 0, 1

		for i := 1; i < n; i++ {
			a, b = b, a+b
		}
		return b
	}
}

func BenchmarkFib20T(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		Fib(20, true)
	}
}

func BenchmarkFib30T(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		Fib(30, true)
	}
}

func BenchmarkFib40T(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		Fib(40, true)
	}
}


func BenchmarkFib20F(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		Fib(20, false)
	}
}

func BenchmarkFib30F(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		Fib(30, false)
	}
}

func BenchmarkFib40F(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		Fib(40, false)
	}
}


type node struct {
	v *int
	t *node
}

func insert(i int, h *node) *node {
	t := &node{&i, nil}

	if h != nil {
		h.t = t
	}
	return t
}


func mkList(n int) *node {

	var h, t *node

	h = insert(0, h)
	t = insert(1, h)

	for i:= 2; i< n; i++ {
		t = insert(i,t)
	}

	return h
}
func sumList(h *node) (i int) {

	for n := h; n != nil; n = n.t {
		i += *h.v
	}

	return i
}

func mkSlice(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = i
	}

	return r
}

func sumSlice(l []int) (i int) {
	for _, v := range l  {
		i += v
	}
	return
}

func Benchmark1List(b *testing.B) {


	for n := 0; n < b.N; n++ {
		l := mkList(1200)
		sumList(l)
	}
 }

 func Benchmark1Slice(b *testing.B) {



	for n := 0; n < b.N; n++ {
		l := mkSlice(1200)
		sumSlice(l)
	}
 }
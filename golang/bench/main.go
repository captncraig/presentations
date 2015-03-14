package main

import (
	"fmt"
	"testing"
)

// START OMIT
func BenchmarkAddSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddSlow(50000, 50000)
	}
}

func BenchmarkAddFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddFast(50000, 50000)
	}
}

// END OMIT

func AddSlow(a, b int) int {
	x := slowAdder{}
	return x.a(a, b)
}

type slowAdder struct{}

func (slowAdder) a(a, b int) int {
	n := 0
	for i := 0; i < a; i++ {
		n++
	}
	for i := 0; i < b; i++ {
		n++
	}
	return n
}

func AddFast(a, b int) int {
	return a + b
}

func TestFoo(t *testing.T) {
	x := 7 + 7
	if x != 14 {
		t.Errorf("Wrong sum: %v != 14\n", x)
	}
}
func main() {
	r := testing.Benchmark(BenchmarkAddSlow)
	fmt.Println("BenchmarkAddSlow", r.String(), r.MemString())
	r = testing.Benchmark(BenchmarkAddFast)
	fmt.Println("BenchmarkAddFast", r.String(), r.MemString())
}

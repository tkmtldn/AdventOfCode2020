package main

import (
	"path/filepath"
	"testing"
)

func TestDay11(t *testing.T){
	path := filepath.Join(".", "test.txt")
	inp, x, y := ReadData(path)

	res1 := SeatingSystem(inp, x, y, Occupied)
	exp1 := 37
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 11 Part 1: got %v want %v", res1, exp1)
	}

	res2 := SeatingSystem(inp, x, y, OccupiedTwo)
	exp2 := 26
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 11 Part 2: got %v want %v", res2, exp2)
	}
}

func BenchmarkDay11(b *testing.B) {
	path := filepath.Join(".", "input.txt")
	inp, x, y := ReadData(path)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SeatingSystem(inp, x, y, Occupied)
	}
}

/*
	go test -bench . day11_test.go
	go test -bench . -benchmem day11_test.go

	go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1 day11_test.go

	go tool pprof main.test.exe cpu.out
	go tool pprof main.test.exe mem.out

	go-torch main.test.exe cpu.out

*/
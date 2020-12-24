package main

import (
	"path/filepath"
	"testing"
)

func TestDay24(t *testing.T) {
	path := filepath.Join(".", "test.txt")
	inp := ReadData(path)

	res1, bw := LobbyLayout(inp)
	exp1 := 10
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 24 Part 1: got %v want %v", res1, exp1)
	}

	res2 := LobbyLayout100(bw)
	exp2 := 2208
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 24 Part 2: got %v want %v", res2, exp2)
	}
}

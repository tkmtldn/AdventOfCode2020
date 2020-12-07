package main

import (
	"path/filepath"
	"testing"
)

func TestDay07_1(t *testing.T) {
	path := filepath.Join(".", "test1.txt")
	n := ReadData(path)

	res1 := CountColours(n, "shiny_gold")
	exp1 := 4
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 7 Part 1: got %v want %v", res1, exp1)
	}

	res2 := SumBags(n, "shiny_gold")
	exp2 := 32
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 7 Part 2: got %v want %v", res2, exp2)
	}
}

func TestDay07_2(t *testing.T) {
	path := filepath.Join(".", "test2.txt")
	n := ReadData(path)

	res := SumBags(n, "shiny_gold")
	expected := 126
	if res != expected {
		t.Errorf("unexpected result on Day 7 Part 2: got %v want %v", res, expected)
	}
}

package main

import (
	"path/filepath"
	"testing"
)

func TestDay03(t *testing.T) {
	path := filepath.Join(".", "day03_testinput.txt")
	n := ReadData(path)

	result_1 := CountTrees(n, slope)
	expected_1 := 7
	if result_1 != expected_1 {
		t.Errorf("unexpected result on Day 3 Part 1: got %v want %v", result_1, expected_1)
	}

	result_2 := CountTreesAgain(n, slopes)
	expected_2 := 336
	if result_2 != expected_2 {
		t.Errorf("unexpected result on Day 3 Part 2: got %v want %v", result_2, expected_2)
	}
}

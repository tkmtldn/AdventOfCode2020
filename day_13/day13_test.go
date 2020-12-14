package main

import (
	"math/big"
	"path/filepath"
	"testing"
)

func TestDay13(t *testing.T){
	path := filepath.Join(".", "test.txt")
	inp := ReadData(path)
	inp2 := ReadData2(path)

	res1 := ShuttleSearch(inp)
	exp1 := 295
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 13 Part 1: got %v want %v", res1, exp1)
	}

	res2 := ShuttleSearch2(inp2)
	exp2 := big.NewInt(1068781)
	if res2.CmpAbs(exp2) != 0 {
		t.Errorf("unexpected result on Day 13 Part 2: got %v want %v", res2, exp2)
	}
}

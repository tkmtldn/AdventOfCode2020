package main

import (
	"path/filepath"
	"testing"
)

func TestDay16(t *testing.T){
	path := filepath.Join(".", "test.txt")
	easy_rules, easy_nearby := ReadData(path)

	res1 := TicketTranslation(easy_rules, easy_nearby)
	exp1 := 71

	if res1 != exp1 {
		t.Errorf("unexpected result on Day 16 Part 1: got %v want %v", res1, exp1)
	}
}
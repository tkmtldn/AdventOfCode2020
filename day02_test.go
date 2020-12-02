package main

import "testing"

func TestDay02(t *testing.T) {
	n := []PasswordCollection{
		{1, 3, "a", "abcde"},
		{1, 3, "b", "cdefg"},
		{2, 9, "c", "ccccccccc"},
	}
	result_1 := PasswordValidation(n)
	expected_1 := 2
	if result_1 != expected_1 {
		t.Errorf("unexpected result on Day 2 Part 1: got %v want %v", result_1, expected_1)
	}

	result_2 := PasswordSecondValidation(n)
	expected_2 := 1
	if result_2 != expected_2 {
		t.Errorf("unexpected result on Day 2 Part 2: got %v want %v", result_2, expected_2)
	}
}


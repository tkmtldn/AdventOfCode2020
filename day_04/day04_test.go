package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDay04_a(t *testing.T) {
	path := filepath.Join(".", "day04_testinput.txt")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file a")
	}
	defer file.Close()

	n := ReadData(file)
	result := ValidPassportCounter(n)
	expected := 2
	if result != expected {
		t.Errorf("unexpected result on Day 4 Part 1: got %v want %v", result, expected)
	}
}

func TestDay04_b(t *testing.T) {
	path := filepath.Join(".", "day04_testValid.txt")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file b")
	}
	defer file.Close()

	n := ReadData(file)
	result := DeepValidPassportCounter(n)
	expected := 4
	if result != expected {
		t.Errorf("unexpected result on Day 4 Part 2: got %v want %v", result, expected)
	}
}

func TestDay04_c(t *testing.T) {
	path := filepath.Join(".", "day04_testInvalid.txt")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file c")
	}
	defer file.Close()

	n := ReadData(file)
	result := DeepValidPassportCounter(n)
	expected := 0
	if result != expected {
		t.Errorf("unexpected result on Day 4 Part 2: got %v want %v", result, expected)
	}
}

package main

import (
	"path/filepath"
	"testing"
)

func TestDay04_VPC(t *testing.T) {
	path := filepath.Join(".", "day04_testinput.txt")
	n := ReadData(path)

	result := ValidPassportCounter(n)
	expected := 2
	if result != expected {
		t.Errorf("unexpected result on Day 4 Part 1: got %v want %v", result, expected)
	}
}

func TestDay04_DVPC_1(t *testing.T) {
	path := filepath.Join(".", "day04_testValid.txt")
	n := ReadData(path)

	result := DeepValidPassportCounter(n)
	expected := 4
	if result != expected {
		t.Errorf("unexpected result on Day 4 Part 2: got %v want %v", result, expected)
	}
}

func TestDay04_DVPC_2(t *testing.T) {
	path := filepath.Join(".", "day04_testInvalid.txt")
	n := ReadData(path)

	result := DeepValidPassportCounter(n)
	expected := 0
	if result != expected {
		t.Errorf("unexpected result on Day 4 Part 2: got %v want %v", result, expected)
	}
}

func TestDay04_byr(t *testing.T) {
	tests := []struct {
		input  string
		want bool
	}{
		{"1919", false},
		{"1920", true},
		{"2002", true},
		{"2003", false},
	}
	for _, tt := range tests {
		t.Run("byr", func(t *testing.T){
			got := byrValidation(tt.input)
			if got != tt.want{
				t.Errorf("Error when running with value %v, got %v, want %v.", tt.input, got, tt.want)
			}
		})
	}

}

func TestDay04_iyr(t *testing.T) {
	tests := []struct {
		input  string
		want bool
	}{
		{"2009", false},
		{"2010", true},
		{"2020", true},
		{"2021", false},
	}
	for _, tt := range tests {
		t.Run("iyr", func(t *testing.T){
			got := iyrValidation(tt.input)
			if got != tt.want{
				t.Errorf("Error when running with value %v, got %v, want %v.", tt.input, got, tt.want)
			}
		})
	}
}

func TestDay04_eyr(t *testing.T) {
	tests := []struct {
		input  string
		want bool
	}{
		{"2019", false},
		{"2020", true},
		{"2030", true},
		{"2031", false},
	}
	for _, tt := range tests {
		t.Run("eyr", func(t *testing.T){
			got := eyrValidation(tt.input)
			if got != tt.want{
				t.Errorf("Error when running with value %v, got %v, want %v.", tt.input, got, tt.want)
			}
		})
	}
}

func TestDay04_hgt(t *testing.T) {
	tests := []struct {
		input  string
		want bool
	}{
		{"190in", false},
		{"60in", true},
		{"190cm", true},
		{"190", false},
	}
	for _, tt := range tests {
		t.Run("hgt", func(t *testing.T){
			got := hgtValidation(tt.input)
			if got != tt.want{
				t.Errorf("Error when running with value %v, got %v, want %v.", tt.input, got, tt.want)
			}
		})
	}
}

func TestDay04_hcl(t *testing.T) {
	tests := []struct {
		input  string
		want bool
	}{
		{"#123abz", false},
		{"#123abc", true},
		{"#1deabc", true},
		{"123abc", false},
	}
	for _, tt := range tests {
		t.Run("hcl", func(t *testing.T){
			got := hclValidation(tt.input)
			if got != tt.want{
				t.Errorf("Error when running with value %v, got %v, want %v.", tt.input, got, tt.want)
			}
		})
	}
}

func TestDay04_ecl(t *testing.T) {
	tests := []struct {
		input  string
		want bool
	}{
		{"wat", false},
		{"brn", true},
	}
	for _, tt := range tests {
		t.Run("ecl", func(t *testing.T){
			got := eclValidation(tt.input)
			if got != tt.want{
				t.Errorf("Error when running with value %v, got %v, want %v.", tt.input, got, tt.want)
			}
		})
	}
}

func TestDay04_pid(t *testing.T) {
	tests := []struct {
		input  string
		want bool
	}{
		{"0123456789", false},
		{"000000001", true},
		{"896056539", true},
		{"186cm", false},
	}
	for _, tt := range tests {
		t.Run("pid", func(t *testing.T){
			got := pidValidation(tt.input)
			if got != tt.want{
				t.Errorf("Error when running with value %v, got %v, want %v.", tt.input, got, tt.want)
			}
		})
	}
}
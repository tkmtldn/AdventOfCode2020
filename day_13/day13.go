package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// https://adventofcode.com/2020/day/13

func ReadData(path string) (elems []Instruction) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		e := scanner.Text()
		o := string(e[0])
		a, _ := strconv.Atoi(e[1:])
		el := Instruction{
			Operation: o,
			Argument:  a,
		}
		elems = append(elems, el)
	}
	return elems
}

func main() {
	path := filepath.Join(".", "day_12", "input.txt")
	inp := ReadData(path)
}
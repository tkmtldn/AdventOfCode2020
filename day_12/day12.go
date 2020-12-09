package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/12

func ReadData(path string) (elems []Instruction) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		e := scanner.Text()
		f := strings.Split(e, " ")
		arg, _ := strconv.Atoi(f[1])
		el := Instruction{
			operation: f[0],
			argument:  arg,
			executed:  0,
		}
		elems = append(elems, el)
	}
	return elems
}

func main(){
	path := filepath.Join(".", "day_12", "input.txt")
	inp := ReadData(path)
}
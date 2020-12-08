package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/8

type Instruction struct {
	operation string
	argument  int
	executed  int
}

func Accumulator(elems []Instruction) int {
	accum := 0
	i := 0
Loop:
	for {
		curr := &elems[i%len(elems)]
		if curr.executed == 1 {
			break Loop
		}
		curr.executed += 1
		if curr.operation == "acc" {
			accum += curr.argument
			i += 1
		}
		if curr.operation == "jmp" {
			i += curr.argument
		}
		if curr.operation == "nop" {
			i += 1
		}
	}
	return accum
}

func Acc(elems []Instruction) (ok bool, accum int) {

	i := 0
	had := 0
Loop:
	for {
		if i == len(elems) {
			ok = true
			break Loop
		}
		if i >= len(elems) || had > len(elems) {
			break Loop
		}
		curr := &elems[i]
		had += 1

		if curr.operation == "acc" {
			accum += curr.argument
			i += 1
		}
		if curr.operation == "jmp" {
			i += curr.argument
		}
		if curr.operation == "nop" {
			i += 1
		}
	}
	return ok, accum
}

func Accumulator2(elems []Instruction) (acc int) {
Loop:
	for i := range (elems) {
		curr := &elems[i]
		if curr.operation == "jmp" {
			curr.operation = "nop"
			ok, v := Acc(elems)
			if ok {
				acc = v
				break Loop
			}
			curr.operation = "jmp"
		}
		if curr.operation == "nop" {
			curr.operation = "jmp"
			ok, v := Acc(elems)
			if ok {
				acc = v
				break Loop
			}
			curr.operation = "nop"
		}
	}
	return acc
}

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

func main() {
	path := filepath.Join(".", "day_08", "day08input.txt")
	inp := ReadData(path)

	fmt.Println("First answer: ", Accumulator(inp))
	fmt.Println("Second answer: ", Accumulator2(inp))
}

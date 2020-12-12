package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// https://adventofcode.com/2020/day/12

type Instruction struct {
	Operation string
	Argument  int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func CalcDist(elems []Instruction) int {
	x := 0
	y := 0
	pos_x := 1
	pos_y := 0
	directions := []int{0, 1, 0, -1}
	for _, e := range elems {
		switch e.Operation {
		case "R":
			if e.Argument == 90 {
				pos_x += 1
				pos_y += 3
			}
			if e.Argument == 180 {
				pos_x += 2
				pos_y += 2
			}
			if e.Argument == 270 {
				pos_x += 3
				pos_y += 1
			}
		case "L":
			if e.Argument == 90 {
				pos_x += 3
				pos_y += 1
			}
			if e.Argument == 180 {
				pos_x += 2
				pos_y += 2
			}
			if e.Argument == 270 {
				pos_x += 1
				pos_y += 3
			}
		case "F":
			x += directions[pos_x%4] * e.Argument
			y += directions[pos_y%4] * e.Argument
		case "N":
			y += e.Argument
		case "S":
			y -= e.Argument
		case "E":
			x += e.Argument
		case "W":
			x -= e.Argument
		}
	}
	return Abs(x) + Abs(y)
}

func CalcDistToo(elems []Instruction) int {
	wp_x := 10
	wp_y := 1
	x := 0
	y := 0
	for _, e := range elems {
		switch e.Operation {
		case "R":
			if e.Argument == 90 {
				wp_x, wp_y = wp_y, wp_x*-1
			}
			if e.Argument == 180 {
				wp_x, wp_y = wp_x*-1, wp_y*-1
			}
			if e.Argument == 270 {
				wp_x, wp_y = wp_y*-1, wp_x
			}
		case "L":
			if e.Argument == 90 {
				wp_x, wp_y = wp_y*-1, wp_x
			}
			if e.Argument == 180 {
				wp_x, wp_y = wp_x*-1, wp_y*-1
			}
			if e.Argument == 270 {
				wp_x, wp_y = wp_y, wp_x*-1
			}
		case "F":
			x += e.Argument*wp_x
			y += e.Argument*wp_y
		case "N":
			wp_y += e.Argument
		case "S":
			wp_y -= e.Argument
		case "E":
			wp_x += e.Argument
		case "W":
			wp_x -= e.Argument
		}
		//fmt.Println(e, wp_x,wp_y, "==", x, y)
	}
	return Abs(x) + Abs(y)
}

func main() {
	path := filepath.Join(".", "day_12", "input.txt")
	inp := ReadData(path)

	fmt.Println("First answer: ", CalcDist(inp))
	fmt.Println("Second answer: ", CalcDistToo(inp))
}

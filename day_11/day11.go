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

// https://adventofcode.com/2020/day/11

func ReadData(path string) (map[string]string, int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	n := [][]string{}
	dict := map[string]string{}

	for scanner.Scan() {
		a := scanner.Text()
		r := strings.Split(a, "")
		n = append(n, r)
	}

	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n[i]); j++ {
			value := n[i][j]
			key := strconv.Itoa(i) + "_" + strconv.Itoa(j)
			dict[key] = value
		}
	}

	return dict, len(n), len(n[0])
}

func Key(x, y int) string {
	return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

func CountOccupied(dict map[string]string, x int, y int) (sum int) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if val, ok := dict[Key(i, j)]; ok {
				if val == "#" {
					sum++
				}
			}
		}
	}
	if dict[Key(x, y)] == "#" {
		sum--
	}
	return sum
}

func CountOccupiedTwo(dict map[string]string, x int, y int, len_x, len_y int) (sum int) {
	// Count to the right
	for i := y + 1; i < len_y; i++ {
		//fmt.Println("1")
		key := Key(x, i)
		if val, ok := dict[key]; ok {
			if val == "#" {
				sum++
				break
			}
			if val == "L" {
				break
			}
		}
	}
	// Count down
	for i := x + 1; i < len_x; i++ {
		//fmt.Println("2")
		key := Key(i, y)
		if val, ok := dict[key]; ok {
			if val == "#" {
				sum++
				break
			}
			if val == "L" {
				break
			}
		}
	}
	// Count to the left
	for i := y - 1; i >= 0; i-- {
		//fmt.Println("3")
		key := Key(x, i)
		if val, ok := dict[key]; ok {
			if val == "#" {
				sum++
				break
			}
			if val == "L" {
				break
			}
		}
	}
	// Count up
	for i := x - 1; i >= 0; i-- {
		//fmt.Println("4")
		key := Key(i, y)
		if val, ok := dict[key]; ok {
			if val == "#" {
				sum++
				break
			}
			if val == "L" {
				break
			}
		}
	}
	// Count diagonally 2 o'clock
	count_x := x
	for i := y + 1; i < len_y; i++ {

		count_x--
		if count_x < 0 {
			break
		}
		//fmt.Println("5")
		key := Key(count_x, i)
		if val, ok := dict[key]; ok {
			if val == "#" {
				sum++
				break
			}
			if val == "L" {
				break
			}
		}
	}
	// Count diagonally 5 o'clock

	count_y := y
	for i := x + 1; i < len_x; i++ {
		//fmt.Println("6")
		count_y++
		key := Key(i, count_y)
		if val, ok := dict[key]; ok {
			if val == "#" {
				sum++
				break
			}
			if val == "L" {
				break
			}
		}
	}
	// Count diagonally 8 o'clock
	count_x = x
	for i := y - 1; i >= 0; i-- {
		//fmt.Println("7")
		count_x ++
		key := Key(count_x, i)
		if val, ok := dict[key]; ok {
			if val == "#" {
				sum++
				break
			}
			if val == "L" {
				break
			}
		}
	}
	// Count diagonally 11 o'clock
	count_y = y
	for i := x - 1; i >= 0; i-- {
		//fmt.Println("8")
		count_y--
		key := Key(i, count_y)
		if val, ok := dict[key]; ok {
			if val == "#" {
				sum++
				break
			}
			if val == "L" {
				break
			}
		}
	}
	return sum
}

func Occupied(inp map[string]string, len_x, len_y int) (map[string]string, int) {

	outp := map[string]string{}
	for i := 0; i < len_x; i++ {
		for j := 0; j < len_y; j++ {
			key := Key(i, j)
			neighbor := CountOccupied(inp, i, j)
			if (neighbor == 0) && (inp[key] == "L") {
				outp[key] = "#"
			} else if (neighbor >= 4) && (inp[key] == "#") {
				outp[key] = "L"
			} else {
				outp[key] = inp[key]
			}
		}
	}

	count_occ := 0
	for _, v := range outp {
		if v == "#" {
			count_occ++
		}
	}

	return outp, count_occ
}

func OccupiedTwo(inp map[string]string, len_x, len_y int) (map[string]string, int) {

	outp := map[string]string{}
	for i := 0; i < len_x; i++ {
		for j := 0; j < len_y; j++ {
			key := Key(i, j)
			neighbor := CountOccupiedTwo(inp, i, j, len_x, len_y)
			if (neighbor == 0) && (inp[key] == "L") {
				outp[key] = "#"
			} else if (neighbor >= 5) && (inp[key] == "#") {
				outp[key] = "L"
			} else {
				outp[key] = inp[key]
			}
		}
	}

	count_occ := 0
	for _, v := range outp {
		if v == "#" {
			count_occ++
		}
	}

	return outp, count_occ
}

func SeatingSystem(inp map[string]string, len_x, len_y int, f func(map[string]string, int, int) (map[string]string, int)) int {
	num_occupied0 := 0

	for i := 0; i < 1000; i++ {
		input_changed, num_occupied := f(inp, len_x, len_y)
		if num_occupied0 == num_occupied {
			break
		}
		inp = input_changed
		num_occupied0 = num_occupied
	}
	return num_occupied0
}

func main() {
	path := filepath.Join(".", "day_11", "input.txt")
	inp, x, y := ReadData(path)
	fmt.Println("First answer: ", SeatingSystem(inp, x, y, Occupied))
	fmt.Println("Second answer: ", SeatingSystem(inp, x, y, OccupiedTwo))
}
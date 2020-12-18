package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// https://adventofcode.com/2020/day/17

type Cube struct {
	x     int
	y     int
	z     int
	value string
}

func ReadData(path string) (elems []Cube) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	cubs := []Cube{}
	this_y := 0

	for scanner.Scan() {
		e := scanner.Text()
		sl := strings.Split(e,"")
		for k, v:= range sl{
			c :=Cube{
				x: k,
				y: this_y,
				z: 0,
				value: v,
			}
			cubs = append(cubs, c)
		}
		this_y ++
	}
	return cubs
}

func main() {
	path := filepath.Join(".", "day_17", "test.txt")
	inp := ReadData(path)
	for _, v:= range inp{
		fmt.Println(v)
	}
}

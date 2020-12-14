package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// https://adventofcode.com/2020/day/14

func ReadData(path string) (elems []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		el := scanner.Text()
		elems = append(elems, el)
	}
	return elems
}

func main() {
	path := filepath.Join(".", "day_14", "test.txt")
	inp := ReadData(path)
	for _, e := range inp {
		if strings.HasPrefix(e, "mask"){
			f := strings.Replace(e, "mask = ", "",1)
			fmt.Println(f)
		}
		if strings.HasPrefix(e, "mem"){
			e = strings.Replace(e, "mem[", "",1)
			e = strings.Replace(e, "]", "",1)
			f := strings.Split(e, "=")
			fmt.Println(f)
		}
	}

	a := 1473355587699697
	b := 692754432903757
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}

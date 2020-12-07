package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// https://adventofcode.com/2020/day/6

func Count1(s []string) int {
	new_string := ""
	result := ""
	for _, v := range (s) {
		new_string += v
	}
	for _, v := range (new_string) {
		a := fmt.Sprintf("%c", v)
		if !strings.Contains(result, a) {
			result += a
		}
	}
	return len(result)
}

func Count2(s []string) (result int) {
	dict := map[string]int{}

	for _, i := range (s) {
		for _, j := range (i) {
			elem := fmt.Sprintf("%c", j)
			if _, ok := dict[elem]; !ok {
				dict[elem] = 1
			} else {
				dict[elem] += 1
			}
		}
	}

	for _, v := range dict {
		if v == len(s) {
			result++
		}
	}
	return result
}

func ReadData(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()

	s := []string{}
	data := [][]string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		element := scanner.Text()
		if element == "" {
			data = append(data, s)
			s = []string{}
		} else {
			s = append(s, element)
		}
	}
	data = append(data, s)
	return data
}

func CountElements(in [][]string, f (func(i []string) int)) (result int) {
	for _, e := range (in) {
		result += f(e)
	}
	return result
}

func main() {

	path := filepath.Join(".", "day_06", "day06input.txt")
	n := ReadData(path)

	fmt.Println("First answer: ", CountElements(n, Count1))
	fmt.Println("Second answer: ", CountElements(n, Count2))

}

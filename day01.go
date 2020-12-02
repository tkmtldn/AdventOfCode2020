package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const finalSum = 2020

func SearchSum(s []int) int{
	result := 0
	Loop: for k, v1 := range s {
		for _, v2 := range s[k+1:] {
			if v1+v2 == finalSum {
				result = v1 * v2
				break Loop
			}
		}
	}
	return result
}

func SearchTripleSum(s []int) int{
	result := 0
Loop: for k, v1 := range s {
	for _, v2 := range s[k+1:] {
		for _, v3 := range s{
			if v1+v2+v3 == finalSum {
			result = v1 * v2 *v3
			break Loop
		}
		}

	}
}
	return result
}

func main() {

	path := filepath.Join(".", "day01input.txt")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("error")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	n := []int{}

	for scanner.Scan() {
		element, _ := strconv.Atoi(scanner.Text())
		n = append(n, element)
	}

	fmt.Println("First answer: ", SearchSum(n))
	fmt.Println("Second answer: ", SearchTripleSum(n))

}


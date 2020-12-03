package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var slope = []int{1, 3}

func CountTrees(n [][]string, s []int) (result int) {
	right := s[0]
	down := s[1]
	position := 0
	for i := 0; i < len(n); i += right {
		elem := n[i]
		if elem[position%len(elem)] == "#" {
			result ++
		}
		position += down
	}
	return result
}

var slopes = [][]int{{1, 1},
	{1, 3},
	{1, 5},
	{1, 7},
	{2, 1},
}

func CountTreesAgain(n [][]string, s [][]int) int {
	result := 1
	result_ans := make([]int, 0)

	for _, elem := range s {
		ans := CountTrees(n, elem)
		result_ans = append(result_ans, ans)
	}

	for _, i := range (result_ans) {
		result *= i
	}
	return result
}

func main() {
	path := filepath.Join(".", "day_03", "day03input.txt")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	n := [][]string{}

	for scanner.Scan() {
		a := scanner.Text()
		r := strings.Split(a, "")
		n = append(n, r)
	}

	fmt.Println("First answer: ", CountTrees(n, slope))
	fmt.Println("Second answer: ", CountTreesAgain(n, slopes))
}

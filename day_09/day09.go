package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

// https://adventofcode.com/2020/day/9

func PreambleC(elems []int, length int) (ans int) {
	for i := length; i < len(elems); i++ {
		curr_el := elems[i]
		curr_elements := elems[(i - length):i]
		if PreambleCount(curr_el, curr_elements) == false {
			ans = curr_el
			return ans
		}
	}
	return ans
}

func PreambleCount(el int, elem []int) (res bool) {
	for i := 0; i < len(elem); i++ {
		for j := 1; j < len(elem); j++ {
			if elem[i]+elem[j] == el {
				res = true
				return
			}
		}
	}
	return res
}

func Sum(elem []int)(res int){
	for _, v := range elem{
		res += v
	}
	return res
}
func PreambleWeakness(elem []int, in int) (res int) {
	start := 0
	finish := 2
Loop:
	for {
		got := Sum(elem[start:finish])
		if got < in {
			finish++
		}
		if got == in {
			date := elem[start:finish]
			sort.Ints(date)
			res = date[0] + date[len(date)-1]
			break Loop
		}
		if got > in {
			start++
			finish =start+2
		}
	}
	return res
}

func ReadData(path string) (elems []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		e := scanner.Text()
		el, _ := strconv.Atoi(e)
		elems = append(elems, el)
	}

	return elems
}

func main() {
	path := filepath.Join(".", "day_09", "input.txt")
	inp := ReadData(path)

	ans := PreambleC(inp, 25)
	fmt.Println("First answer: ", ans)
	fmt.Println("Second answer: ", PreambleWeakness(inp, ans))
}

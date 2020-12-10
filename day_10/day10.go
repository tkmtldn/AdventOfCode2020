package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

// https://adventofcode.com/2020/day/10

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
	sort.Ints(elems)
	return elems
}

func AdapterDiff(inp []int)int{
	cur := 0
	one := 0
	three := 1
	for i := 0; i < len(inp); i++ {
		if inp[i]-cur == 1 {
			one++
			cur = inp[i]
		} else if inp[i]-cur == 3 {
			three++
			cur = inp[i]
		}
	}
	return one*three
}

func AdapterComb(inp []int) int{
	inp = append(inp, inp[len(inp)-1]+3)
	inp = append(inp, 0)
	sort.Ints(inp)

	res :=1
	currrent :=-1
	dict := map[int]int{}
	for i := 1; i < len(inp); i++ {
		if inp[i]-inp[i-1] == 1 {
			currrent ++
		} else {
			if currrent>0{
				if _, ok := dict[currrent]; ok {
					dict[currrent]++
				} else {
					dict[currrent] = 1
				}
			}
			currrent = -1
		}
	}

	for k, v := range(dict){
		if k>=1 {
			res *= Mult(k,v)
		}
	}
	return res
}

func Mult(x, y int) int{
	var new_x float64
	switch x{
	case 3:
		new_x =7
	case 2:
		new_x =4
	case 1:
		new_x =2
	}
	return int(math.Pow(new_x, float64(y)))

}

func main() {
	path := filepath.Join(".", "day_10", "input.txt")
	inp := ReadData(path)

	fmt.Println("First answer: ", AdapterDiff(inp))
	fmt.Println("Second answer: ", AdapterComb(inp))

}

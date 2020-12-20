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

// https://adventofcode.com/2020/day/20

func ReadData(path string) (map[int][]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	elems := map[int][]string{}
	k := 0
	v := []string{}

	for scanner.Scan() {
		e := scanner.Text()
		if e == "" {
			elems[k] = Ribs(v)
			v = []string{}
			continue
		}
		if strings.HasPrefix(e, "Tile ") {
			e = strings.Replace(e, ":", "", -1)
			el := strings.Split(e, " ")
			k, _ = strconv.Atoi(el[1])
			continue
		}
		v = append(v, e)
	}
	return elems
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func Ribs(s []string) (result []string) {
	top := ""
	left := ""
	right := ""
	bottom := ""

	limit := len(s) - 1
	for i, e := range s {
		dict := strings.Split(e, "")
		if i == 0 {
			top += strings.Join(dict, "")
		}
		if i == limit {
			bottom += strings.Join(dict, "")
		}
		left += dict[0]
		right += dict[limit]
	}
	result = append(result, top)
	result = append(result, right)
	result = append(result, bottom)
	result = append(result, left)
	result = append(result, Reverse(top))
	result = append(result, Reverse(right))
	result = append(result, Reverse(bottom))
	result = append(result, Reverse(left))
	return
}

func JurassicJigsaw(inp map[int][]string) int {
	data := []string{}
	for _, v := range inp {
		data = append(data, v...)
	}

	values := map[int]int{}
	for k, v := range inp {
		count := 0
		for _, vv := range v {
			for _, e := range data {
				if vv == e {
					count++
				}
			}
		}
		values[k] = (count - 8) / 2
	}

	result := 1
	for k, v := range values {
		fmt.Println(k, v)
		if v == 2 {
			result *= k
		}
	}
	return result
}

func main() {
	path := filepath.Join(".", "day_20", "test.txt")
	inp := ReadData(path)
	fmt.Println("First answer:", JurassicJigsaw(inp))
}

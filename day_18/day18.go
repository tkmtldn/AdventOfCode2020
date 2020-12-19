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

// https://adventofcode.com/2020/day/18

func ReadData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	elems := []string{}

	for scanner.Scan() {
		e := scanner.Text()
		elems = append(elems, e)
	}
	return elems
}

func CalculateWithoutPriority(s string) string {
	data := strings.Split(s, " ")
	result, _ := strconv.Atoi(data[0])
	for i := 2; i < len(data); i += 2 {
		operator := data[i-1]
		y, _ := strconv.Atoi(data[i])
		switch operator {
		case "+":
			result += y
		case "*":
			result *= y
		}
	}
	res := strconv.Itoa(result)
	return res
}

func CalculateWithPriority(s string) string {
	data := strings.Split(s, " ")
	datastr := make([]string, 0)

	for i:=0; i<len(data); i++ {
		if data[i]=="+"{
			x, _ := strconv.Atoi(datastr[len(datastr)-1])
			y, _ := strconv.Atoi(data[i+1])
			res := x+y
			z := strconv.Itoa(res)
			datastr = datastr[:len(datastr)-1]
			datastr = append(datastr, z)
			i++
		} else {
			datastr = append(datastr, data[i])
		}
	}

	result := 1
	datastrmul := strings.Join(datastr, "")
	datastrmu := strings.Split(datastrmul, "*")

	for _, e := range datastrmu{
		int_e, _ := strconv.Atoi(e)
		result *= int_e
	}
	res := strconv.Itoa(result)

	return res
}

func WithoutPriority(s string) string {

	levels := strings.Count(s, "(")
	if levels == 0 {
		return CalculateWithoutPriority(s)
	} else {
		start := 0
		finish := 0
		for i, v := range s{
			if v == 40{
			start = i
			}
			if v == 41{
				finish = i
				break
			}
		}
		result := CalculateWithoutPriority(s[start+1 : finish])
		r := s[:start] + result + s[finish+1:]
		return WithoutPriority(r)
	}
	return ""
}

func WithPriority(s string) string {

	levels := strings.Count(s, "(")
	if levels == 0 {
		return CalculateWithPriority(s)
	} else {
		start := 0
		finish := 0
		for i, v := range s{
			if v == 40{
				start = i
			}
			if v == 41{
				finish = i
				break
			}
		}
		result := CalculateWithPriority(s[start+1 : finish])
		r := s[:start] + result + s[finish+1:]
		return WithPriority(r)
	}
	return ""
}

func OperationOrder(inp []string, f func(string)string) (sum int) {
	for _, e := range inp {
		res, _ := strconv.Atoi(f(e))
		sum += res
	}
	return sum
}

func main() {
	path := filepath.Join(".", "day_18", "input.txt")
	inp := ReadData(path)
	fmt.Println("First answer:", OperationOrder(inp, WithoutPriority))
	fmt.Println("Second answer:", OperationOrder(inp, WithPriority))
}

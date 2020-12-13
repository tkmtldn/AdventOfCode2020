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

// https://adventofcode.com/2020/day/13

func ReadData(path string) (elem []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		e := scanner.Text()
		elems := strings.Split(e, ",")
		for _, el := range elems {
			if el != "x" {
				integ, _ := strconv.Atoi(el)
				elem = append(elem, integ)
			}
		}
	}
	return elem
}

func ReadData2(path string) []SingleBus {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	elem := []string{}
	for scanner.Scan() {
		e := scanner.Text()
		elems := strings.Split(e, ",")
		for _, el := range elems {
			elem = append(elem, el)
		}
	}

	list_of_buses := []SingleBus{}
	for i := 1; i < len(elem); i++ {
		result, err := strconv.Atoi(elem[i])
		if err == nil {
			b := SingleBus{
				Frequency: result,
				Interval:  i - 1,
			}
			list_of_buses = append(list_of_buses, b)
		}
	}

	return list_of_buses
}

func Bus(inp []int) int {
	num_bus := 0
	interval := 100000
	for i := 1; i < len(inp); i++ {
		inter := (inp[0]/inp[i]+1)*inp[i] - inp[0]
		if inter < interval {
			interval = inter
			num_bus = inp[i]
		}
	}
	return num_bus * interval
}

type SingleBus struct {
	Frequency int
	Interval  int
}

func Bus2(inp []SingleBus) int {
	delta := inp[0].Frequency
	d := int(150000000000000 / delta * delta)
	//d:=0
	limit := 400000000000000
	for i := 0; i < limit; i += delta {
		if Couting(inp, i) {
			return i
			break
		}
	}
	return d
}

func Couting(inp []SingleBus, i int) bool {
	ans := true
	for _, e := range inp {
		if (i+e.Interval)%e.Frequency != 0 {
			ans = false
			break
		}
	}
	return ans
}

func main() {
	//path := filepath.Join(".", "day_13", "test.txt")
	path := filepath.Join(".", "day_13", "input.txt")
	inp := ReadData(path)
	fmt.Println(Bus(inp))

	inp2 := ReadData2(path)
	fmt.Println(Bus2(inp2))
}

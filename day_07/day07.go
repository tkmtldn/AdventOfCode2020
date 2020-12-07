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

type Container struct {
	color string
	qty   int
}

func ReadData(path string) (map[string][]Container) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	n := [][]string{}
	for scanner.Scan() {
		a := scanner.Text()
		r := strings.Replace(a, ",", "", -1)
		r = strings.Replace(r, ".", "", -1)
		n = append(n, strings.Split(r, " "))
	}

	dict := map[string][]Container{}

	for _, line := range n {
		parent_name := line[0] + "_" + line[1]
		parent_content := []Container{}
		for i := 3; i < len(line); i++ {
			if (line[i] == "bag" || line[i] == "bags") && line[i-2] != "no" {
				child_name := line[i-2] + "_" + line[i-1]
				child_qty, _ := strconv.Atoi(line[i-3])
				e := Container{
					color: child_name,
					qty:   child_qty,
				}
				parent_content = append(parent_content, e)
			}
		}
		dict[parent_name] = parent_content
	}

	return dict
}

func SumBags(c map[string][]Container, s string) (sum int) {
	for k, v := range c {
		if k == s {
			sum += SumLine(c[k])
			for _, e := range v {
				if e.qty != 0 {
					sum += e.qty * SumBags(c, e.color)
				}
			}
		}
	}
	return sum
}

func SumLine(line []Container) (sum int) {
	for _, e := range line {
		sum += e.qty
	}
	return sum
}

func SumColours(sq string, cur string, c map[string][]Container) (sum int) {
Loop:
	for _, e := range c[cur] {
		if e.color == sq {
			sum ++
			break Loop
		} else {
			sum += SumColours(sq, e.color, c)
		}
	}
	return sum
}

func CountColours(inp_data map[string][]Container, inp_str string) (sum int) {
	for k, _ := range inp_data {
		d := SumColours(inp_str, k, inp_data)
		if d != 0 {
			sum++
		}
	}
	return sum
}

func main() {
	path := filepath.Join(".", "day_07", "day07input.txt")
	n := ReadData(path)

	fmt.Println("First answer: ", CountColours(n, "shiny_gold"))
	fmt.Println("Second answer: ", SumBags(n, "shiny_gold"))
}

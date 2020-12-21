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

type Frame struct {
	full_frame       []string
	inner_frame      []string
	all_sides        []string
	sides            map[int]string
	reversed         map[int]string
	sides_neighbours map[int]int
}

func ReadData(path string) (map[int]Frame) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	new_elems := map[int]Frame{}

	k := 0
	v := []string{}
	f := Frame{}

	for scanner.Scan() {
		e := scanner.Text()
		if e == "" {
			f.full_frame = v
			new_elems[k] = f
			f = Frame{}
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
	return new_elems
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func Ribs(s []string) ([]string, map[int]string, map[int]string, []string) {
	side1 := ""
	side2 := ""
	side3 := ""
	side4 := ""

	inner_frame := []string{}

	limit := len(s) - 1
	for i, e := range s {
		dict := strings.Split(e, "")
		if i == 0 {
			side1 += strings.Join(dict, "")
		}
		if i == limit {
			side3 += strings.Join(dict, "")
		}
		side2 += dict[0]
		side4 += dict[limit]
		line := strings.Join(dict[1:limit], "")
		inner_frame = append(inner_frame, line)
	}

	map_sides := map[int]string{
		1: side1,
		2: side2,
		3: side3,
		4: side4,
	}

	map_reversed_sides := map[int]string{
		1: Reverse(side1),
		2: Reverse(side2),
		3: Reverse(side3),
		4: Reverse(side4),
	}

	all_sides := []string{}
	for _, v := range map_sides {
		all_sides = append(all_sides, v)
	}
	for _, v := range map_reversed_sides {
		all_sides = append(all_sides, v)
	}
	return all_sides, map_sides, map_reversed_sides, inner_frame
}

func PrepareData(inp map[int]Frame) map[int]Frame {
	new := map[int]Frame{}
	short := map[int][]string{}

	for k, v := range inp {
		all, _, _, _ := Ribs(v.full_frame)
		short[k] = all
	}

	for k, v := range inp {
		f := Frame{}
		all, sides, reversed, inner := Ribs(v.full_frame)
		f.all_sides = all
		f.sides = sides
		f.reversed = reversed
		f.inner_frame = inner
		f.full_frame = v.full_frame

		dict_side_n := map[int]int{}
		for sides_key, sides_value := range sides {
			for short_key, short_value := range short {
				for _, sv := range short_value {
					if (sv == sides_value) && (short_key != sides_key) && (short_key != k) {
						//fmt.Println("In", k, "at postion", sides_key, "we have", sides_value, "with", short_key)
						dict_side_n[sides_key] = short_key
					}
				}
			}
		}
		f.sides_neighbours = dict_side_n
		new[k] = f
	}

	return new
}

func JurassicJigsaw(input map[int]Frame) int {

	all_frames := PrepareData(input)
	result := 1
	for k, v := range all_frames {
		if len(v.sides_neighbours) == 2 {
			result *= k
		}
	}

	return result
}

func JurassicJigsawSecond(input map[int]Frame) int {

	all_frames := PrepareData(input)
	column := 0
	column_down := 0
	next_column := 0

	for k, v := range all_frames {
		if len(v.sides_neighbours) == 2 {
			this := 0
			if _, ok := v.sides_neighbours[2]; ok {
				this++
			}
			if _, ok := v.sides_neighbours[2]; ok {
				this++
			}
			if this == 2 {
				column = k
				column_down = v.sides_neighbours[3]
				next_column = v.sides_neighbours[2]
			}
		}
	}
	fmt.Println(column, column_down, next_column)

	return 0
}

func main() {
	path := filepath.Join(".", "day_20", "test.txt")
	inp := ReadData(path)
	fmt.Println("First answer:", JurassicJigsaw(inp))
	fmt.Println("Second answer:", JurassicJigsawSecond(inp))
}

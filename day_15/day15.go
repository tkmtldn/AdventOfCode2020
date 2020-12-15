package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/15

type Numbers struct {
	pos      int
	last_pos int
	repeat   int
}

func RambunctiousRecitation (s string, limit int) string {
	inp := strings.Split(s, ",")
	dict := map[string]Numbers{}

	last_spoken := ""
	for i := 0; i < len(inp); i++ {
		dict[inp[i]] = Numbers{
			pos:      i+1,
			last_pos: i+1,
			repeat:   1,
		}
		last_spoken = inp[i]
	}

	for i := len(inp)+1; i <= limit; i++ {
		if val, ok := dict[last_spoken]; ok {
			if val.repeat == 1 {
				last_spoken = "0"
			}
			if val.repeat > 1 {
				last_spoken = strconv.Itoa(val.last_pos - val.pos)
			}
		} else {
			val := Numbers{
				pos : i-1,
				last_pos: i-1,
				repeat: 1,
			}
			dict[last_spoken] = val
			last_spoken = "0"

		}
		if v, ok := dict[last_spoken]; ok {
			v.repeat++
			v.pos, v.last_pos = v.last_pos, i
			dict[last_spoken] = v
		}
	}
	return last_spoken
}


func main() {
	s := "2,0,1,7,4,14,18"
	fmt.Println(RambunctiousRecitation(s, 2020))
	fmt.Println(RambunctiousRecitation(s, 30000000))
}
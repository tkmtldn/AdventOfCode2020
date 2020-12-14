package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/14

var mask string
var dict2 = map[int]int{}

func ReadData(path string) (elems []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		el := scanner.Text()
		elems = append(elems, el)
	}
	return elems
}

func BinString(e string, kv int, kn int) (key int, new string) {
	regular := regexp.MustCompile(`[\d]+`)
	inputs := regular.FindAllString(e, -1)

	key, _ = strconv.Atoi(inputs[kv])
	n, _ := strconv.Atoi(inputs[kn])
	st := strconv.FormatInt(int64(n), 2)
	for i := 0; i < 36-len(st); i++ {
		new += "0"
	}
	new += st
	return key, new
}

func CmpMask(mask string, bin string, priority string) (ans string) {
	mask1 := strings.Split(mask, "")
	bin1 := strings.Split(bin, "")
	for i := 0; i < 36; i++ {
		if mask1[i] == priority {
			ans += bin1[i]
		} else {
			ans += mask1[i]
		}
	}
	return ans
}

func DockingData(inp []string) (sum int) {
	dict := map[int]int{}
	for _, e := range inp {
		if strings.HasPrefix(e, "mask") {
			f := strings.Replace(e, "mask = ", "", 1)
			mask = f
		}
		if strings.HasPrefix(e, "mem") {
			key, s := BinString(e,0, 1)
			res := CmpMask(mask, s, "X")
			output, _ := strconv.ParseInt(res, 2, 64)
			dict[key] = int(output)
		}
	}
	for _, v := range dict {
		sum += v
	}
	return sum
}

func ExplorePerm(r string, val int) {
	count := float64(strings.Count(r, "X"))
	vars := math.Pow(2, count)

	for i := 0; i < int(vars); i++ {
		ans := ""
		integ := strconv.FormatInt(int64(i), 2)
		for i := 0; i < (int(count) - len(integ)); i++ {
			ans += "0"
		}
		ans += integ

		res := ""
		pos := 0

		ans1 := strings.Split(ans, "")
		r1 := strings.Split(r, "")
		for i := 0; i < 36; i++ {
			if r1[i] == "X" {
				res += ans1[pos]
				pos++
			} else {
				res += r1[i]
			}
		}
		output, _ := strconv.ParseInt(res, 2, 64)
		dict2[int(output)] = val
	}
}

func DockingData2(inp []string) (sum int) {
	for _, e := range inp {
		if strings.HasPrefix(e, "mask") {
			f := strings.Replace(e, "mask = ", "", 1)
			mask = f
		}
		if strings.HasPrefix(e, "mem") {
			val, s := BinString(e,1,0)
			res := CmpMask(mask, s, "0")

			ExplorePerm(res, val)
		}
	}
	for _, v := range dict2 {
		sum += v
	}
	return sum
}

func main() {
	path := filepath.Join(".", "day_14", "input.txt")
	inp := ReadData(path)

	fmt.Println("First answer: ", DockingData(inp))
	fmt.Println("Second answer: ", DockingData2(inp))
}

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

// https://adventofcode.com/2020/day/14

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

func BinString(e string) (key int, new string) {
	e = strings.Replace(e, "mem[", "",1)
	e = strings.Replace(e, "]", "",1)
	e = strings.Replace(e, " ", "",-1)
	f := strings.Split(e, "=")
	key, _ = strconv.Atoi(f[0])
	n, _ := strconv.Atoi(f[1])
	st := strconv.FormatInt(int64(n), 2)
	for i:=0; i<36-len(st); i++ {
		new += "0"
	}
	new += st
	return key, new
}

func CmpMask(m string, e string) (r string) {
	m1 := strings.Split(m, "")
	e1 := strings.Split(e, "")
	for i:=0; i<36; i++ {
		if m1[i] == "X" {
			r += e1[i]
		} else {
			r += m1[i]
		}
	}
	return r
}

var mask string

func DockingData(inp []string) (sum int){
	dict := map[int]int{}
	for _, e := range inp {
		if strings.HasPrefix(e, "mask"){
			f := strings.Replace(e, "mask = ", "",1)
			mask = f
		}
		if strings.HasPrefix(e, "mem"){
			key, s := BinString(e)
			res := CmpMask(mask, s)
			output, _ := strconv.ParseInt(res, 2, 64)
			dict[key]=int(output)
		}
	}
	for _, v:= range dict{
		sum +=v
	}
	return sum
}

func main() {
	path := filepath.Join(".", "day_14", "input.txt")
	inp := ReadData(path)

	fmt.Println("First answer: ", DockingData(inp))
}

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

// https://adventofcode.com/2020/day/19

func ReadData(path string) (map[string]string, []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	data := map[string]string{}
	messages := []string{}
	level := 0

	for scanner.Scan() {
		e := scanner.Text()
		if e == "" {
			level++
			continue
		}
		if level == 0 {
			e := strings.Replace(e, "\"", "", -1)
			es := strings.Split(e, ": ")
			key := es[0]
			value := es[1]
			data[key] = value
		}
		if level == 1 {
			messages = append(messages, e)
		}
	}
	return data, messages
}

var replacing = map[string]string{}

func CheckIfNumbers(s string) bool {
	ans := true
	list := strings.Split(s, " ")
	for _, e := range list {
		_, err := strconv.Atoi(e)
		if err == nil {
			ans = false
		}
	}
	return ans
}

func MonsterMessages(d map[string]string, m []string) int {

	up := len(d)

	for k, v := range d {
		if len(v) == 1 && (v == "a" || v == "b") {
			replacing[k] = v
			delete(d, k)
		}
	}
	//fmt.Println(d)

	for {
		if len(replacing) >= up {
			break
		}
		for k, v := range d {
			list := strings.Split(v, " ")
			list2 := []string{}
			for _, e := range list {
				if value, ok := replacing[e]; ok {
					list2 = append(list2, value)
				} else {
					list2 = append(list2, e)
				}
			}
			v = strings.Join(list2, " ")
			//fmt.Println("v=========", v)
			if CheckIfNumbers(v) {
				v = strings.Replace(v, " ", "", -1)
				v = strings.Replace(v, "|", " | ", -1)
				v = strings.Replace(v, "[", " [ ", -1)
				v = strings.Replace(v, "]", " ] ", -1)
				replacing[k] = "[" + v + "]"
				delete(d, k)
			}
		}
	}
	for k, v := range replacing {
		fmt.Println(k, "==", v)
	}
	return 0
}

func main() {
	path := filepath.Join(".", "day_19", "test.txt")
	d, m := ReadData(path)
	fmt.Println(MonsterMessages(d, m))
}

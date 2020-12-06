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

func ValidPassportCounter(n []map[string]string) (count int) {
	for _, elem := range n {
		if len(elem) == 7 {
			count ++
		}
	}
	return count
}

func DeepValidPassportCounter(n []map[string]string) (count int) {
	for _, elem := range n {
		if len(elem) == 7 {
			deep_count := 0
			for key, val := range elem {
				switch key {
				case "byr":
					if byrValidation(val) {
						deep_count++
					}
				case "iyr":
					if iyrValidation(val) {
						deep_count++
					}
				case "eyr":
					if eyrValidation(val) {
						deep_count++
					}
				case "hgt":
					if hgtValidation(val) {
						deep_count++
					}
				case "hcl":
					if hclValidation(val) {
						deep_count++
					}
				case "ecl":
					if eclValidation(val) {
						deep_count++
					}
				case "pid":
					if pidValidation(val) {
						deep_count++
					}
				}

			}
			if deep_count == 7 {
				count ++
			}
		}
	}
	return count
}

func byrValidation(s string) bool {
	v, _ := strconv.Atoi(s)
	return v >= 1920 && v <= 2002
}

func iyrValidation(s string) bool {
	v, _ := strconv.Atoi(s)
	return v >= 2010 && v <= 2020
}

func eyrValidation(s string) bool {
	v, _ := strconv.Atoi(s)
	return v >= 2020 && v <= 2030
}

func hgtValidation(s string) bool {
	if strings.Contains(s, "cm") {
		a := strings.Replace(s, "cm", "", 1)
		v, _ := strconv.Atoi(a)
		return v >= 150 && v <= 193
	}
	if strings.Contains(s, "in") {
		a := strings.Replace(s, "in", "", 1)
		v, _ := strconv.Atoi(a)
		return v >= 59 && v <= 76
	}
	return false
}

func hclValidation(s string) bool {
	first := strings.HasPrefix(s, "#")
	s = strings.Replace(s, "#", "", 1)
	_, err := strconv.ParseUint(s, 16, 64)
	return (err == nil) && first && len(s)==6
}

func eclValidation(s string) bool {
	return (s == "amb") || (s == "blu") || (s == "brn") || (s == "gry") || (s == "grn") || (s == "hzl") || (s == "oth")
}

func pidValidation(s string) bool {
	_, err := strconv.Atoi(s)
	return (err == nil) && len(s) <= 9
}

func ReadData(path string) (n []map[string]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file")
	}
	defer file.Close()

	m := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a := scanner.Text()
		if len(a) != 0 {
			data := strings.Split(a, " ")
			for _, elem := range data {
				e := strings.Split(elem, ":")
				if e[0] != "cid" {
					m[e[0]] = e[1]
				}
			}
		} else {
			n = append(n, m)
			m = make(map[string]string)
		}
	}
	n = append(n, m)
	return n
}

func main() {
	path := filepath.Join(".", "day_04", "day04input.txt")
	n := ReadData(path)

	fmt.Println("First answer: ", ValidPassportCounter(n))
	fmt.Println("Second answer: ", DeepValidPassportCounter(n))

}

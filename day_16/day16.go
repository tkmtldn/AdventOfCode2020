package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/16

type Variety struct {
	Name  string
	Range map[int]int
}

func ReadData(path string) (map[int]int, map[int]int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()

	interval := 0
	regular := regexp.MustCompile(`[\d]+-[\d]+`)
	scanner := bufio.NewScanner(file)

	rules := map[int]int{}
	your := map[int]int{}
	nearby := map[int]int{}

	for scanner.Scan() {
		e := scanner.Text()
		if e == "" {
			interval++
		}
		if interval == 0 {
			inputs := regular.FindAllString(e, -1)
			for _, i := range inputs {
				params := strings.Split(i, "-")
				start, _ := strconv.Atoi(params[0])
				finish, _ := strconv.Atoi(params[1])
				for k := start; k <= finish; k++ {
					rules[k] = 0
				}
			}
		}
		if interval == 1 {
			if e != "your ticket:" && e != "" {
				inputs := strings.Split(e, ",")
				for _, v := range inputs {
					e, _ := strconv.Atoi(v)
					your[e] = 1
				}
			}
		}
		if interval == 2 {
			if e != "nearby tickets:" && e != "" {
				inputs := strings.Split(e, ",")
				for _, v := range inputs {
					e, _ := strconv.Atoi(v)
					if _, ok := nearby[e]; !ok {
						nearby[e] = 1
					} else {
						nearby[e] += 1
					}
				}
			}
		}
	}
	return rules, nearby
}

func ReadData2(path string) ([]Variety, map[int]int, map[int][]int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()

	interval := 0
	regular_name := regexp.MustCompile(`[\w ]+:`)
	regular_digits := regexp.MustCompile(`([\d]+)-([\d])+ or ([\d])+-([\d])+`)
	scanner := bufio.NewScanner(file)

	rules := []Variety{}
	your := map[int]int{}
	nearby := map[int][]int{}

	for scanner.Scan() {
		e := scanner.Text()
		if e == "" {
			interval++
		}
		if interval == 0 {
			input_n := regular_name.FindAllString(e, -1)
			input_d := regular_digits.FindAllString(e, -1)

			name := strings.Replace(input_n[0], ":", "", -1)
			digits_raw := strings.Replace(input_d[0], " or ", "-", -1)
			digits := strings.Split(digits_raw, "-")

			v := map[int]int{}

			s1, _ := strconv.Atoi(digits[0])
			f1, _ := strconv.Atoi(digits[1])
			s2, _ := strconv.Atoi(digits[2])
			f2, _ := strconv.Atoi(digits[3])
			for i := s1; i <= f1; i++ {
				v[i] = 1
			}
			for i := s2; i <= f2; i++ {
				v[i] = 1
			}
			new := Variety{
				Name:  name,
				Range: v,
			}
			rules = append(rules, new)
		}
		if interval == 1 {
			if e != "your ticket:" && e != "" {
				inputs := strings.Split(e, ",")
				for k, v := range inputs {
					e, _ := strconv.Atoi(v)
					your[k] = e
				}
			}
		}

		if interval == 2 {
			if e != "nearby tickets:" && e != "" {
				inputs := strings.Split(e, ",")
				for k, v := range inputs {
					e, _ := strconv.Atoi(v)
					if _, ok := nearby[k]; !ok {
						nearby[k] = []int{e}
					} else {
						all := nearby[k]
						all = append(all, e)
						nearby[k] = all
					}
				}
			}
		}
	}

	return rules, your, nearby
}

func Possibilies(i_num int, r []int, varie []Variety) (int, map[int][]string) {

	coinc_list := map[int][]string{}

	for _, v := range varie {
		count := 0
		for _, vr := range r {
			if _, ok := v.Range[vr]; ok {
				count++
			}
		}
		vname := strings.Replace(v.Name, " ", "_", -1)
		if val, ok := coinc_list[count]; ok {
			val = append(val, vname)
			coinc_list[count] = val
		} else {
			val = []string{vname}
			coinc_list[count] = val
		}
	}
	return i_num, coinc_list
}

func TicketTranslation(varieties map[int]int, nearby map[int]int) (sum int) {
	for key, val := range nearby {
		if _, ok := varieties[key]; !ok {
			sum += key * val
		}
	}
	return sum
}

func TicketTranslation2(varieties []Variety, your map[int]int, nearby map[int][]int) int {
	black_map := map[int][]string{}
	for k, v := range nearby {
		position, dict := Possibilies(k, v, varieties)
		max := 0
		for k, _ := range dict{
			if k>max{
				max = k
			}
		}
		black_map[position] = dict[max]
	}

	white_map := map[int]string{}

	for i:=0; i<20; i++{
		name_str := ""
		for k, v:= range black_map{
			if len(v) == 1{
				name_str = v[0]
				white_map[k] = name_str
			}
		}
		for key, v:= range black_map{
			new_value := []string{}
			for _, vv := range v{
				if vv != name_str{
					new_value = append(new_value, vv)
				}
			}
			black_map[key] = new_value
		}
	}

	result := 1
	for k, v:= range white_map{
		if strings.HasPrefix(v, "departure"){
			result *= your[k]
		}
	}
	return result
}

func main() {
	path := filepath.Join(".", "day_16", "input.txt")

	easy_rules, easy_nearby := ReadData(path)
	fmt.Println("First answer:", TicketTranslation(easy_rules, easy_nearby))

	rules, your, nearby := ReadData2(path)
	fmt.Println("Second answer:", TicketTranslation2(rules, your, nearby))
}

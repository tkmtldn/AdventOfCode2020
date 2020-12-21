package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// https://adventofcode.com/2020/day/21

func ReadData(path string) ([]string, map[string][][]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	regular := regexp.MustCompile(` \([\w ,]+\)`)

	list_of_ingredients := []string{}
	possible_vocab := map[string][][]string{}

	for scanner.Scan() {
		e := scanner.Text()
		raw_allerg := regular.FindAllString(e, -1)
		allerg := raw_allerg[0]
		ingred := strings.Replace(e, allerg, "", -1)
		allerg = strings.Replace(allerg, ",", "", -1)
		allerg = strings.Replace(allerg, ")", "", -1)
		allerg = strings.Replace(allerg, " (contains ", "", -1)
		ingredients := strings.Split(ingred, " ")
		list_of_ingredients = append(list_of_ingredients, ingredients...)
		allergens := strings.Split(allerg, " ")
		for _, v := range allergens {
			if val, ok := possible_vocab[v]; ok {
				val = append(val, ingredients)
				possible_vocab[v] = val
			} else {
				val := [][]string{}
				val = append(val, ingredients)
				possible_vocab[v] = val
			}
		}
	}
	return list_of_ingredients, possible_vocab
}

func NormalizeVocab(pv map[string][][]string) map[string]string {
	normal_vocab := map[string]string{}
	possibles := map[string][]string{}

	for k, v := range pv {
		possibles[k] = Possible(v)
	}

	for i := 0; i < len(pv); i++ {
		for k, v := range possibles {
			if len(v) == 1 {
				normal_vocab[v[0]] = k
			}
			if len(v) > 1 {
				pos := []string{}
				for _, vv := range v {
					if _, ok := normal_vocab[vv]; !ok {
						pos = append(pos, vv)
					}
				}
				if len(pos) == 1 {
					normal_vocab[pos[0]] = k
				}
			}
		}
	}
	return normal_vocab
}

func Possible(values [][]string) []string {
	all_values := []string{}
	for _, v := range values {
		all_values = append(all_values, v...)
	}

	inclusions := map[string]int{}
	for _, v := range all_values {
		if _, ok := inclusions[v]; ok {
			inclusions[v]++
		} else {
			inclusions[v] = 1
		}
	}

	possible := []string{}
	for k, v := range inclusions {
		if v == len(values) {
			possible = append(possible, k)
		}
	}
	return possible
}

func AllergenAssessment(list []string, pv map[string][][]string) (sum int, res string) {
	normalized := NormalizeVocab(pv)
	for _, e := range list {

			if _, ok := normalized[e]; !ok {
				sum++
			}

	}

	canonical := []string{}
	ans := []string{}
	reversed := map[string]string{}

	for k, v := range normalized {
		canonical = append(canonical, v)
		reversed[v] = k
	}
	sort.Strings(canonical)
	for _, e := range canonical {
		ans = append(ans, reversed[e])
	}

	res = strings.Join(ans, ",")

	return sum, res
}

func main() {
	path := filepath.Join(".", "day_21", "input.txt")
	list, pv := ReadData(path)
	ans1, ans2 := AllergenAssessment(list, pv)

	fmt.Println("First answer:", ans1)
	fmt.Println("Second answer:", ans2)
}

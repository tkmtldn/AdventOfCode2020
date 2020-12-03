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

type PasswordCollection struct {
	MinRep      int
	MaxRep      int
	SymbRep     string
	PassExample string
}

func NewPasswordCollection(inp string) *PasswordCollection {
	p := PasswordCollection{}

	s := strings.Split(inp, " ")
	s1 := strings.Split(s[0], "-")

	p.MinRep, _ = strconv.Atoi(s1[0])
	p.MaxRep, _ = strconv.Atoi(s1[1])
	p.SymbRep = strings.Replace(s[1], ":", "", 1)
	p.PassExample = s[2]

	return &p
}

func PasswordValidation(n []PasswordCollection) (result int) {
	for _, elem := range n {
		reps := strings.Count(elem.PassExample, elem.SymbRep)
		if reps >= elem.MinRep && reps <= elem.MaxRep {
			result += 1
		}
	}
	return result
}

func PasswordSecondValidation(n []PasswordCollection) (result int) {
	for _, elem := range n {
		res := 0
		r := strings.Split(elem.PassExample, "")
		if (elem.MinRep <= len(r) && r[elem.MinRep-1] == elem.SymbRep) {
			res += 1
		}
		if (elem.MaxRep <= len(r) && r[elem.MaxRep-1] == elem.SymbRep) {
			res += 1
		}
		if res == 1 {
			result += 1
		}
	}
	return result
}

func main() {

	path := filepath.Join(".", "day_02", "day02input.txt")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	n := []PasswordCollection{}

	for scanner.Scan() {
		a := scanner.Text()
		p := NewPasswordCollection(a)
		n = append(n, *p)
	}

	fmt.Println("First answer: ", PasswordValidation(n))
	fmt.Println("Second answer: ", PasswordSecondValidation(n))
}

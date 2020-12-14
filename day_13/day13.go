package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/13

type Residue struct {
	remainder *big.Int
	modulus   *big.Int
}

func ReadData(path string) (elem []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		e := scanner.Text()
		elems := strings.Split(e, ",")
		for _, el := range elems {
			if el != "x" {
				integ, _ := strconv.Atoi(el)
				elem = append(elem, integ)
			}
		}
	}
	return elem
}

func ReadData2(path string) []*Residue {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	elem := []string{}

	for scanner.Scan() {
		e := scanner.Text()
		elems := strings.Split(e, ",")
		for _, el := range elems {
			elem = append(elem, el)
		}
	}

	residues := []*Residue{}

	for i := 1; i < len(elem); i++ {
		result, err := strconv.Atoi(elem[i])
		if err == nil {
			residues = append(residues, &Residue{
				remainder: big.NewInt(int64(i-1)),
				modulus: big.NewInt(int64(result)),
			})
		}
	}

	return residues
}

func ShuttleSearch(inp []int) int {
	num_bus := 0
	interval := 100000
	for i := 1; i < len(inp); i++ {
		inter := (inp[0]/inp[i]+1)*inp[i] - inp[0]
		if inter < interval {
			interval = inter
			num_bus = inp[i]
		}
	}
	return num_bus * interval
}

func ChineseRemainder(residues []*Residue) *big.Int {
	acc := residues[0]
	for i := 1; i < len(residues); i++ {
		a1, a2 := acc.remainder, residues[i].remainder
		n1, n2 := acc.modulus, residues[i].modulus
		m1, m2 := new(big.Int), new(big.Int)

		solution := new(big.Int).GCD(m1, m2, n1, n2)

		left := new(big.Int).Mul(a1, m2)
		left.Mul(left, n2)

		right := new(big.Int).Mul(a2, m1)
		right.Mul(right, n1)

		solution.Add(left, right)
		combined := new(big.Int).Mul(n1, n2)

		solution.Mod(solution, combined)
		acc = &Residue{remainder: solution, modulus: combined}
	}
	return acc.remainder
}

func ShuttleSearch2(inp []*Residue) *big.Int  {
	chinese := ChineseRemainder(inp)
	eq := big.NewInt(1)
	for _, e := range inp{
		eq.Mul(eq, e.modulus)
	}
	res := big.NewInt(0)
	res.Sub(eq, chinese)
	return res
}

func main() {
	path := filepath.Join(".", "day_13", "input.txt")

	inp := ReadData(path)
	fmt.Println("First answer: ", ShuttleSearch(inp))

	inp2 := ReadData2(path)
	fmt.Println("Second answer: ", ShuttleSearch2(inp2))
}

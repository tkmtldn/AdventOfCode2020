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

func ReadData(path string) ([]int, []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	input_0 := make([]int, 0, 0)
	input_1 := make([]int, 0, 0)

	inp_level := 0

	for scanner.Scan() {
		e := scanner.Text()
		if strings.HasPrefix(e, "Player ") {
			continue
		}
		if e == "" {
			inp_level++
			continue
		}
		if inp_level == 0 {
			num, _ := strconv.Atoi(e)
			input_0 = append(input_0, num)
		}
		if inp_level == 1 {
			num, _ := strconv.Atoi(e)
			input_1 = append(input_1, num)
		}
	}
	return input_0, input_1
}

func CrabCombat(p1 []int, p2 []int) int {
	winner := []int{}
	for {
		if len(p1) == 0{
			winner = p2
			break
		}
		if len(p2) == 0{
			winner = p1
			break
		}
		player1 := p1[0]
		player2 := p2[0]
		if player1 > player2 {
			p1 = p1[1:]
			p2 = p2[1:]
			p1 = append(p1, player1)
			p1 = append(p1, player2)
		} else {
			p1 = p1[1:]
			p2 = p2[1:]
			p2 = append(p2, player2)
			p2 = append(p2, player1)
		}
	}

	res := 0
	mul := len(winner)
	for _, e:= range winner{
		res += e*mul
		mul--
	}

	return res
}

func RecursiveCombat(p1 []int, p2 []int) int {
	return 0
}

func main() {
	path := filepath.Join(".", "day_22", "test.txt")
	inp1, inp2 := ReadData(path)

	fmt.Println("First answer:", CrabCombat(inp1, inp2))
	fmt.Println("Second answer:", RecursiveCombat(inp1, inp2))
}

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
		if len(p1) == 0 {
			winner = p2
			break
		}
		if len(p2) == 0 {
			winner = p1
			break
		}

		var round_winner int

		player1 := p1[0]
		player2 := p2[0]

		if player1 > player2 {
			round_winner = 1
		} else {
			round_winner = 2
		}

		if round_winner == 1 {
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
	return CalcResult(winner)
}

func CalcResult(winner []int) (res int) {
	mul := len(winner)
	for _, e := range winner {
		res += e * mul
		mul--
	}
	return
}

var mem_winner = []int{}

func CrabCombat2(p1 []int, p2 []int) int {

	memory := map[string]int{}
	RecursiveCombat(p1, p2, memory)

	return CalcResult(mem_winner)
}

func RecursiveCombat(p1 []int, p2 []int, memory map[string]int) int {
	var round_winner int
	//var round = 1
	for {
		//fmt.Println(round, p1, p2)
		//round++

		key := strings.Trim(strings.Replace(fmt.Sprint(p1), " ", "_", -1), "[]") + "__" +
			strings.Trim(strings.Replace(fmt.Sprint(p2), " ", "_", -1), "[]")
		if _, ok := memory[key]; ok {
			round_winner = 1
			break
		} else {
			memory[key] = 0
		}

		player1 := p1[0]
		player2 := p2[0]

		if len(p1) >= player1+1 && len(p2) >= player2+1 {
			new_p1 := []int{}
			for _, v := range p1[1 : player1+1] {
				new_p1 = append(new_p1, v)
			}
			new_p2 := []int{}
			for _, v := range p2[1 : player2+1] {
				new_p2 = append(new_p2, v)
			}
			new_memory := map[string]int{}
			round_winner = RecursiveCombat(new_p1, new_p2, new_memory)
		} else {
			if player1 < player2 {
				round_winner = 2
			} else {
				round_winner = 1
			}
		}

		if round_winner == 1 {
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

		if len(p1) == 0 {
			round_winner = 2
			mem_winner = p2
			break
		}
		if len(p2) == 0 {
			round_winner = 1
			mem_winner = p1
			break
		}
	}
	return round_winner
}

func main() {
	path := filepath.Join(".", "day_22", "input.txt")
	inp1, inp2 := ReadData(path)

	fmt.Println("First answer:", CrabCombat(inp1, inp2))
	fmt.Println("Second answer:", CrabCombat2(inp1, inp2))
}

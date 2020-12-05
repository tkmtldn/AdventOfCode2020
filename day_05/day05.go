package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func SeatIDSearch(s string) int {
	row := 0
	rowMax := 128
	seat := 0
	seatMax := 8

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 66:
			row += (rowMax - row) / 2
		case 70:
			rowMax -= (rowMax - row) / 2
		case 76:
			seatMax -= (seatMax - seat) / 2
		case 82:
			seat += (seatMax - seat) / 2
		}
	}
	return 8*row + seat
}

func MaxSeatId(n []int) int {
	sort.Ints(n)
	return n[len(n)-1]
}

func YourSeatId(n []int) (ans int) {
	sort.Ints(n)
	for i := 0; i < len(n)-1; i++ {
		if n[i+1]-n[i] != 1 {
			ans = n[i] + 1
			break
		}
	}
	return ans
}

func main() {
	path := filepath.Join(".", "day_05", "day05input.txt")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	elems := []int{}

	for scanner.Scan() {
		e := SeatIDSearch(scanner.Text())
		elems = append(elems, e)
	}

	fmt.Println("First answer: ", MaxSeatId(elems))
	fmt.Println("Second answer: ", YourSeatId(elems))

}

package main

import (
	"fmt"
)

func Transform(pk int) int {
	loop := 0
	value := 1
	subj_number := 7
	for {
		if value == pk {
			return loop
		}
		loop++
		value *= subj_number
		value %= 20201227
	}
}

func ComboBreaker(card int, door int) int {
	d := Transform(door)
	value := 1
	for i := 1; i <= d; i++ {
		value *= card
		value %= 20201227
	}
	return value
}

func main() {
	cardPK := 8458505
	doorPK := 16050997
	fmt.Println(ComboBreaker(cardPK, doorPK))
}

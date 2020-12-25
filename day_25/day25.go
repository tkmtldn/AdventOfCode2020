package main

import (
	"fmt"
)

func Transform(pk int) int {
	l := 0
	v := 1
	s := 7
	for {
		if v == pk {
			return l
		}
		l++
		v *= s
		v %= 20201227
	}
}

func ComboBreaker(card int, door int) int {
	c := Transform(card)
	d := Transform(door)
	fmt.Println(c, d)
	v := 1
	for i := 1; i <= d; i++ {
		v *= card
		v %= 20201227
	}
	return v
}

func main() {
	cardPKtest := 5764801
	doorPKtest := 17807724
	cardPK := 8458505
	doorPK := 16050997
	fmt.Println(ComboBreaker(cardPKtest, doorPKtest))
	fmt.Println(ComboBreaker(cardPK, doorPK))
}

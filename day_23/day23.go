package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CrabCups(cups []int) {
	for i := 0; i < 10; i++ {

		current := cups[i%9]
		pick_up0 := cups[(i+1)%9]
		pick_up1 := cups[(i+2)%9]
		pick_up2 := cups[(i+3)%9]
		destination := GetDestination(current, pick_up0, pick_up1, pick_up2)

		new_cups := []int{}
		new_cups1 := []int{}
		new_cups2 := []int{}
		position := 0
		for k := 0; k < 9; k++ {
			if cups[k] == destination {
				position = k
			}
		}
		new_cups = append(new_cups, destination)
		new_cups = append(new_cups, pick_up0)
		new_cups = append(new_cups, pick_up1)
		new_cups = append(new_cups, pick_up2)
		for k := position + 1; k < 9; k++ {
			if (cups[k] != destination) || (cups[k] != pick_up0) || (cups[k] != pick_up1) || (cups[k] != pick_up2) {
				new_cups1 = append(new_cups1, cups[k])

			}
		}
		for k := 0; k <= position; k++ {
			if cups[k] == destination {
				continue
			} else if cups[k] == pick_up0 {
				continue
			} else if cups[k] == pick_up1 {
				continue
			} else if cups[k] == pick_up2 {
				continue
			} else {
				new_cups2 = append(new_cups2, cups[k])
			}
		}
		fmt.Println(new_cups2, new_cups, new_cups1)
		final := []int{}
		for _, e := range new_cups2{
			final = append(final,e)
		}
		for _, e := range new_cups{
			final = append(final,e)
		}
		for _, e := range new_cups1{
			final = append(final,e)
		}
		cups = final
		//fmt.Println(i+2, cups, position)
		//fmt.Println(cups, pick_up0, pick_up1, pick_up2, "====", destination)
	}
}

func GetDestination(curr, p0, p1, p2 int) int {
	res := curr - 1
	for {
		if res == 0 {
			res += 9
		}
		if res != p0 && res != p1 && res != p2 {
			break
		}
		res--
	}
	return res
}

func main() {
	input := "389125467"
	//input := "326519478"

	inp_str := strings.Split(input, "")
	cups := []int{}
	for _, e := range inp_str {
		i, _ := strconv.Atoi(e)
		cups = append(cups, i)
	}

	CrabCups(cups)
}

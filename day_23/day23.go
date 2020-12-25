package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetStep(dest int, cups []int) (ans int) {
	for k, v := range cups {
		if v == dest {
			ans = k
			break
		}
	}
	return
}

func CrabCups(cups []int) string {
	for i := 0; i < 100; i++ {
		new_cups := []int{}
		fmt.Println(cups)
		current := cups[i%9]
		pick_up0 := cups[(i+1)%9]
		pick_up1 := cups[(i+2)%9]
		pick_up2 := cups[(i+3)%9]
		destination := GetDestination(current, pick_up0, pick_up1, pick_up2)
		step := GetStep(destination, cups)

		for i, _ := range cups {
			if cups[i] == pick_up0 || cups[i] == pick_up1 || cups[i] == pick_up2 {
				continue
			} else if cups[i] == destination {
				new_cups = append(new_cups, destination)
				new_cups = append(new_cups, pick_up0)
				new_cups = append(new_cups, pick_up1)
				new_cups = append(new_cups, pick_up2)
			} else {
				new_cups = append(new_cups, cups[i])
			}
		}
		if step == 0 {
			new_cups = GetNewCup(new_cups, 3)
		} else if step == 1 {
			new_cups = GetNewCup(new_cups, 2)
		} else if step == 2 {
			new_cups = GetNewCup(new_cups, 1)
		} else if i%9 == 6 {
			new_cups = GetNewCup(new_cups, 2)
		} else if i%9 == 7 {
			new_cups = GetNewCup(new_cups, 1)
		} else if i%9 == 8 {
			new_cups = GetNewCup(new_cups, 0)
		} else {
		}
		cups = new_cups
	}

	key := 0
	ans := []int{}
	for k, v := range cups {
		if v == 1 {
			key = k
		}
	}
	for k, v:= range cups{
		if k>=key{
			ans = append(ans, v)
		}
	}
	for k, v:= range cups{
		if k<key{
			ans = append(ans, v)
		}
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ans)), ""), "[]")
}

func GetNewCup(new_cups []int, num int) []int {
	nc := []int{}
	for k, v := range new_cups {
		if k >= num {
			nc = append(nc, v)
		}
	}
	for k, v := range new_cups {
		if k < num {
			nc = append(nc, v)
		}
	}
	return nc
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

	fmt.Println("First answer:", CrabCups(cups))
}

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

func ReadData(path string) ([]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error. Problem with opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	input := make([]string, 0, 0)

	for scanner.Scan() {
		e := scanner.Text()
		input = append(input, e)
	}
	return input
}

func LobbyLayout(n []string) (int, map[string]int) {

	bw := map[string]int{}

	for _, e := range n {
		x := 0
		y := 0
		z := 0
		for i := 0; i < len(e); i++ {
			if e[i] == 101 {
				//go east
				x += 1
				y -= 1
			} else if e[i] == 119 {
				//go west
				x -= 1
				y += 1
			} else if e[i] == 110 && e[i+1] == 101 {
				//go ne
				x += 1
				z -= 1
				i++
			} else if e[i] == 110 && e[i+1] == 119 {
				//go nw
				z -= 1
				y += 1
				i++
			} else if e[i] == 115 && e[i+1] == 101 {
				//go se
				z += 1
				y -= 1
				i++
			} else if e[i] == 115 && e[i+1] == 119 {
				//go sw
				x -= 1
				z += 1
				i++
			} else {
				continue
			}
		}

		key := strconv.Itoa(x) + "_" + strconv.Itoa(y) + "_" + strconv.Itoa(z)
		if _, ok := bw[key]; ok {
			bw[key] += 1
		} else {
			bw[key] = 0
		}
	}

	for k, v := range bw {
		if v != 0 {
			delete(bw, k)
		}
	}

	return len(bw), bw
}

func GetKeys(x, y, z int) []string {
	keys := []string{}
	keys = append(keys, strconv.Itoa(x+1)+"_"+strconv.Itoa(y-1)+"_"+strconv.Itoa(z))
	keys = append(keys, strconv.Itoa(x+1)+"_"+strconv.Itoa(y)+"_"+strconv.Itoa(z-1))
	keys = append(keys, strconv.Itoa(x)+"_"+strconv.Itoa(y-1)+"_"+strconv.Itoa(z+1))
	keys = append(keys, strconv.Itoa(x)+"_"+strconv.Itoa(y+1)+"_"+strconv.Itoa(z-1))
	keys = append(keys, strconv.Itoa(x-1)+"_"+strconv.Itoa(y)+"_"+strconv.Itoa(z+1))
	keys = append(keys, strconv.Itoa(x-1)+"_"+strconv.Itoa(y+1)+"_"+strconv.Itoa(z))
	return keys
}

func LobbyLayout100(bw map[string]int) int {
	for i := 0; i < 100; i++ {
		newBW := map[string]int{}
		CHECK := map[string]int{}

		for k, _ := range bw {
			CHECK[k] = 0
			xyz := strings.Split(k, "_")
			x, _ := strconv.Atoi(xyz[0])
			y, _ := strconv.Atoi(xyz[1])
			z, _ := strconv.Atoi(xyz[2])
			k := GetKeys(x, y, z)
			for _, e := range k {
				CHECK[e] = 0
			}
		}

		for k, _ := range CHECK {
			xyz := strings.Split(k, "_")
			x, _ := strconv.Atoi(xyz[0])
			y, _ := strconv.Atoi(xyz[1])
			z, _ := strconv.Atoi(xyz[2])
			ans := 0
			keys := GetKeys(x, y, z)
			for _, e := range keys {
				if _, ok := bw[e]; ok {
					ans++
				}
			}
			if _, ok := bw[k]; ok {
				if ans == 1 || ans == 2 {
					newBW[k] = 0
				}
			}
			if _, ok := bw[k]; !ok {
				if ans == 2 {
					newBW[k] = 0
				}
			}
		}
		bw = newBW
	}
	return len(bw)
}

func main() {
	path := filepath.Join(".", "day_24", "input.txt")
	inp := ReadData(path)
	res1, bw := LobbyLayout(inp)
	fmt.Println("First answer: ", res1)
	fmt.Println("Second anser: ", LobbyLayout100(bw))
}

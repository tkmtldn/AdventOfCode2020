package main

import "fmt"

// https://adventofcode.com/2020/day/13

//func ReadData(path string) (elems []Instruction) {
//	file, err := os.Open(path)
//	if err != nil {
//		log.Fatalf("Error. Problem with opening file.")
//	}
//	defer file.Close()
//	scanner := bufio.NewScanner(file)
//
//	for scanner.Scan() {
//		e := scanner.Text()
//		o := string(e[0])
//		a, _ := strconv.Atoi(e[1:])
//		el := Instruction{
//			Operation: o,
//			Argument:  a,
//		}
//		elems = append(elems, el)
//	}
//	return elems
//}

func Reversed(num int) string{
	for i := 1; i <= 36; i++ {
		if (i*2)%num == 0 {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
	return "\n000000000000000000000000000000001011"
}

func main() {
	//path := filepath.Join(".", "day_12", "input.txt")
	//inp := ReadData(path)
	fmt.Println(Reversed(11))
}

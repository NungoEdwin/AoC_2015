package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./santainput2.txt")
	if err != nil {
		fmt.Println("error reading the file")
	}
	datastr := string(data)
	// datastr += ""
	fmt.Println(Floor(datastr))
}

func Floor(n string) (int, int) {
	sum := 0
	before := 0
	for i, l := range n {
		if l == '(' {
			sum += 1
		} else {
			sum -= 1
		}
		if sum == -1 {
			before = i + 1
			break
		}
	}
	return sum, before
}

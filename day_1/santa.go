package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./santainput.txt")
	if err != nil {
		fmt.Println("error reading the file")
	}
	datastr := string(data)
	// datastr += ""
	fmt.Println(Floor(datastr))
}

func Floor(n string) int {
	sum := 0
	for _, l := range n {
		if l == '(' {
			sum += 1
		} else {
			sum -= 1
		}
	}
	return sum
}

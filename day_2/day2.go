package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytedata, err := os.ReadFile("./input1")
	if err != nil {
		fmt.Println("error reading file")
	}
	stringdata := string(bytedata)
	dataslice := strings.Split(stringdata, "\n")
	// dataslice := []string{"2x3x4", "1x1x10"}
	fmt.Println(SurfaceArea(dataslice))
}

func SurfaceArea(n []string) (int, int) {
	components := []string{}
	final := 0
	slicea := []int{}
	sum := 0
	for _, str := range n {
		components = strings.Split(str, "x")
		length, _ := strconv.Atoi(components[0])
		width, _ := strconv.Atoi(components[1])
		height, _ := strconv.Atoi(components[2])

		a := 2 * (length * width)
		b := 2 * (length * height)
		c := 2 * (height * width)

		SA := a + b + c
		intsz := []int{length, width, height}

		sort.Ints(intsz)

		perimeter := intsz[0] + intsz[1] + intsz[0] + intsz[1]
		volume := intsz[0] * intsz[1] * intsz[2]

		ribbon := perimeter + volume
		sum += ribbon
		// fmt.Println(sum)

		// why this part doest do right sorting?
		// slack := a / 2

		// if a < b && a < c {
		// 	slack = a / 2
		// } else if b < a && b < c {
		// 	slack = b / 2
		// } else {
		// 	slack = c / 2
		// }

		// smallest := slack

		smallest := intsz[0] * intsz[1]

		slicea = append(slicea, SA+smallest)

		final += (SA + smallest)
	}
	a := 0
	for _, ab := range slicea {
		a += ab
	}
	return final, a
}

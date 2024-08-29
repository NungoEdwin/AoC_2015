package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	filebyte, _ := os.ReadFile("./input1")
	file := string(filebyte)
	fmt.Println(Direct(file))
}

func Direct(s string) int {
	a, b := 0, 0

	points := make(map[int]string)
	for i, sign := range s {
		if sign == '>' || sign == '<' {
			if sign == '>' {
				b += 1
			} else {
				b -= 1
			}

			b1 := strconv.Itoa(b)
			a1 := strconv.Itoa(a)
			points[i] = string(b1) + string(a1)

		} else if sign == '^' || sign == 'v' {
			if sign == '^' {
				a += 1
			} else {
				a -= 1
			}
			b2 := strconv.Itoa(b)
			a2 := strconv.Itoa(a)
			points[i] = b2 + a2
		}
	}
	// adding the origin which originally wasnt in the map or accounted for
	points[len(points)] = "00"

	// copying a map
	points1 := make(map[int]string)
	for i, v := range points {
		points1[i] = v
	}

	for _, t := range points {
		c := 0
		for i, j := range points1 {
			if t == j {
				if c == 0 {
					c += 1
				} else {
					delete(points1, i)
				}
			}
		}
		if c > 1 {
			b += 1
		}
	}

	return len(points1)
	// alternative way simpler since keys are unique in maps duplicates are handled automatically
	/*points := make(map[[2]int]bool)
	points[[2]int{a, b}] = true

	for _, cha := range s {
		switch cha {
		case '^':
			a++
		case 'v':
			a--
		case '>':
			b++
		case '<':
			b--
		}
		points[[2]int{a, b}] = true
	}
	return len(points)*/
}

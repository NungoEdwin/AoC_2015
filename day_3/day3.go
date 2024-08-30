package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	filebyte, _ := os.ReadFile("./input1")
	filebyte2, _ := os.ReadFile("./input2")
	file := string(filebyte)
	file2 := string(filebyte2)
	fmt.Println(Direct(file))
	fmt.Println(Direct2(file2))
	fmt.Println(Direct2("^>v<"))
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
	points1 = DelDuplicate(points, points1)

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

func Direct2(s string) int {
	b, a := 0, 0
	c, d := 0, 0

	points := make(map[int]string)
	points2 := make(map[int]string)

	for i, sign := range s {
		if sign == '>' || sign == '<' {
			if sign == '>' {
				if i%2 == 0 {
					b += 1
					c = b
					b1 := strconv.Itoa(b)
					a1 := strconv.Itoa(a)
					points[i] = string(b1) + string(a1)

				} else {
					c += 1
					b = c
					c1 := strconv.Itoa(c)
					d1 := strconv.Itoa(d)
					points2[i] = string(c1) + string(d1)
				}
			} else {
				if i%2 == 0 {
					b -= 1
					c = b
					b1 := strconv.Itoa(b)
					a1 := strconv.Itoa(a)
					points[i] = string(b1) + string(a1)

				} else {
					c -= 1
					b = c
					c1 := strconv.Itoa(c)
					d1 := strconv.Itoa(d)
					points2[i] = string(c1) + string(d1)
				}
			}

			// points2[i] = string(c) + string(d)
		} else if sign == '^' || sign == 'v' {
			if sign == '^' {
				if i%2 == 0 {
					a += 1
					d = a
					b1 := strconv.Itoa(b)
					a1 := strconv.Itoa(a)
					points[i] = string(b1) + string(a1)
				} else {
					d += 1
					a = d
					c1 := strconv.Itoa(c)
					d1 := strconv.Itoa(a)
					points2[i] = string(c1) + string(d1)
				}
			} else {
				if i%2 == 0 {
					a -= 1
					d = a
					b1 := strconv.Itoa(b)
					a1 := strconv.Itoa(a)
					points[i] = string(b1) + string(a1)
				} else {
					d -= 1
					a = d
					c1 := strconv.Itoa(c)
					d1 := strconv.Itoa(d)
					points2[i] = string(c1) + string(d1)
				}
			}
			// b2 := strconv.Itoa(b)
			// a2 := strconv.Itoa(a)
			// points[i] = b2 + a2
		}
	}
	// adding the origin which originally wasnt in the map or accounted for
	points[len(points)] = "00"
	// trimming duplicates from acopy of points(Santa)
	Trimpoints := make(map[int]string)
	for i, j := range points {
		Trimpoints[i] = j
	}
	Trimpoints = DelDuplicate(points, Trimpoints)
	// trimming duplicates from points2(Santa-Robot)
	// adding origin first
	points2[len(points2)] = "00"
	Trimpoints2 := make(map[int]string)
	for i, j := range points2 {
		Trimpoints2[i] = j
	}
	Trimpoints2 = DelDuplicate(points2, Trimpoints2)

	/*creating a unified map to use the values as the keys to eliminate duplicates in the unified
	Map consisting of SantaMap(Trimpoints) and SantaRobot(Trimpoints2)*/
	UniquMap := make(map[string]int)
	for i, j := range Trimpoints {
		UniquMap[j] = i
	}
	for i, j := range Trimpoints2 {
		UniquMap[j] = i
	}

	return len(Trimpoints) + len(Trimpoints2)
}

func DelDuplicate(n, n1 map[int]string) map[int]string {
	for _, t := range n {
		c := 0
		for i, j := range n1 {
			if t == j {
				if c == 0 {
					c += 1
				} else {
					delete(n1, i)
				}
			}
		}
		// if c > 1 {
		// 	b += 1
		// }
	}
	return n1
}

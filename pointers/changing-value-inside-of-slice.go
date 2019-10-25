package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var list []*Point
	var address *Point
	for i := 0; i < 10; i++ {
		var t Point
		address = &t
		t.x = i

		list = append(list, &t)
	}
	Iterate(list)
	*address = Point{x: 20, y: 20}
	fmt.Println("------------------------------------")
	Iterate(list)
}

func Iterate(list []*Point) {
	for i, j := range list {
		fmt.Println(i, j)
	}
}

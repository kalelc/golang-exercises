package main

import "fmt"

type Point struct {
	x, y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	spokes int
}

func main() {
	var w = Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v\n", w)
}

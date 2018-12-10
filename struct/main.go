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
	var w Wheel
	w.x = 10     // equivalent to w.Circle.Point.X
	w.y = 12     // equivalent to w.Circle.Point.Y
	w.Radius = 5 // equivalent to w.Circle.Radius
	w.spokes = 8

	fmt.Println(w.Circle.Radius)
}

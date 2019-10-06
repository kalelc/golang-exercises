package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter the acceleration: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Enter the initial velocity: ")
	scanner.Scan()
	i1, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("enter the initial displacement: ")
	scanner.Scan()
	s1, _ := strconv.ParseFloat(scanner.Text(), 64)

	fn := GenDisplaceFn(a, i1, s1)

	fmt.Print("enter the time: ")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Println(fn(t))
}

func GenDisplaceFn(a, i1, s1 float64) func(t float64) float64 {
	f := func(t float64) float64 {
		s := ((0.5 * a) * (math.Pow(t, 2))) + (i1 * t) + s1
		return s
	}
	return f
}

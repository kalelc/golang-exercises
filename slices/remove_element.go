package main

import "fmt"

func main() {
	elements := []int{1, 3, 6, 2, 10, 2, 4, 5}

	removeElement(&elements, 0)
	fmt.Println(elements)

}
func removeElement(elements *[]int, i int) {
	s := *elements
	s = append(s[:i], s[i+1:]...)
	*elements = s
}

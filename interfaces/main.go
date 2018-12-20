package main

import "fmt"

func main() {
	dog := Dog{name: "zeus"}
	var s Animal
	s = &dog
	fmt.Println(s.Bay())
}

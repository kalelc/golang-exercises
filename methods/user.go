package main

import "fmt"

type User struct {
	name string
	age  int
}

// getters example
// Remember: calling a function makes a copy of each argument value
// for that is better pass a pointer

func (u *User) Name() string {
	return u.name
}

func (u *User) Age() int {
	return u.age
}

func (u *User) SetName(value string) {
	u.name = value
}

func (u *User) SetAge(value int) {
	u.age = value
}

func main() {
	user := User{"foo", 22}
	fmt.Println(user)
	user.SetName("bar")
	user.SetAge(25)
	fmt.Println(user)
}

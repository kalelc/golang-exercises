package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var animals map[string]Animal
var elements map[string]Animal

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (a Cow) Eat() string {
	return a.food
}

func (a Cow) Move() string {
	return a.locomotion
}

func (a Cow) Speak() string {
	return a.noise
}

func (a Bird) Eat() string {
	return a.food
}

func (a Bird) Move() string {
	return a.locomotion
}

func (a Bird) Speak() string {
	return a.noise
}

func (a Snake) Eat() string {
	return a.food
}

func (a Snake) Move() string {
	return a.locomotion
}

func (a Snake) Speak() string {
	return a.noise
}
func main() {
	PopulateAnimals()
	elements = make(map[string]Animal)

	scanner := bufio.NewScanner(os.Stdin)
	var err error

	for {
		fmt.Println("Enter the command that you want to execute: 'newanimal', 'query'")
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		err = ValidCommand(input)
		if err == nil {
			if input == "query" {
				fmt.Print("1. The string 'query'. 2. the type of the animal. 3. third is the action('eat', 'move', or 'speak'): ")
				scanner.Scan()
				input := scanner.Text()
				data := strings.Split(input, " ")
				err := GetAnimalInfo(data)
				if err != nil {
					fmt.Println(err)
					continue
				}
			} else {
				fmt.Print("1. The string 'newanimal'. 2. the name of the new animal. 3. the type of the new animal ('cow', 'bird', or 'snake'): ")
				scanner.Scan()
				input := scanner.Text()
				data := strings.Split(input, " ")
				CreateAnimals(data)
			}
		} else {
			fmt.Println(err)
			continue
		}
	}
}

func CreateAnimals(data []string) (err error) {
	command := string(data[0])

	err = ValidCommand(command)

	if err == nil {
		name := string(data[1])
		animalType := string(data[2])
		var animal Animal

		switch animalType {
		case "cow":
			animal = Cow{"grass", "walk", "moo"}
		case "bird":
			animal = Bird{"worms", "fly", "peep"}
		case "snake":
			animal = Snake{"mice", "slither", "hsss"}
		}
		elements[name] = animal
	}
	fmt.Println("Created it!")
	return
}

func PopulateAnimals() {
	animals = map[string]Animal{
		"cow":   Cow{"grass", "walk", "moo"},
		"bird":  Bird{"worms", "fly", "peep"},
		"snake": Snake{"mice", "slither", "hsss"},
	}
}

func GetAnimalInfo(data []string) (err error) {
	command := string(data[0])
	err = ValidCommand(command)
	if err == nil {
		name := string(data[1])
		action := string(data[2])

		animal := animals[name]

		switch action {
		case "eat":
			fmt.Println(animal.Eat())
		case "move":
			fmt.Println(animal.Move())
		case "speak":
			fmt.Println(animal.Speak())
		}
	}
	return
}

func ValidCommand(command string) (err error) {
	if command == "newanimal" || command == "query" {
		err = nil
	} else {
		err = errors.New("invalid command")
	}
	return
}

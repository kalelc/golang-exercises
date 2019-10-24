package main

import (
	"fmt"
	"time"
)

const Split = 4

func main() {
	numbers := [17]int{2, 3, 5, 7, 11, 13, 2, 4, 5, 6, 7, 8, 2, 4, 5, 7, 9}
	parts := len(numbers) / Split
	mod := len(numbers) % Split
	counter := 0

	if mod != 0 {
		parts++
	}

	elements := make(chan []int, parts)

	for i := 0; i < parts; i++ {
		if i > 0 {
			counter += Split
		}

		if mod != 0 && (i+1) == parts {
			go ConcurrentSort(numbers[counter:counter+mod], elements)
		} else {
			go ConcurrentSort(numbers[counter:(counter+Split)], elements)
		}
	}

	time.Sleep(4 * time.Second)

	close(elements)

	var result []int

	for element := range elements {
		result = append(result, element...)
	}

	Sort(&result)
	fmt.Println(result)
}

func ConcurrentSort(numbers []int, elements chan []int) {
	Sort(&numbers)
	elements <- numbers
}

func Sort(numbers *[]int) {
	list := *numbers
	for i := 1; i < len(list); i++ {
		for j := 0; j < len(list)-i; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}
}

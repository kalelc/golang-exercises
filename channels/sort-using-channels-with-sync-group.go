package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const Split = 4

var wg sync.WaitGroup

func main() {
	var numbers []int

	fmt.Print("Enter a series of integers, each integer must be separated by space, Example: 5 2 1 3 :> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	elements := strings.Split(s, " ")

	for _, value := range elements {

		num, _ := strconv.Atoi(strings.Trim(value, " "))
		numbers = append(numbers, num)
	}

	parts := len(numbers) / Split
	mod := len(numbers) % Split
	counter := 0

	if mod != 0 {
		parts++
	}

	ch := make(chan []int, parts)
	wg.Add(parts)

	for i := 0; i < parts; i++ {
		if i > 0 {
			counter += Split
		}

		if mod != 0 && (i+1) == parts {
			go ConcurrentSort(numbers[counter:counter+mod], ch)
		} else {
			go ConcurrentSort(numbers[counter:(counter+Split)], ch)
		}
	}

	wg.Wait()

	close(ch)

	var result []int

	for value := range ch {
		fmt.Print("Sorted subarray -> ")
		fmt.Println(value)
		result = append(result, value...)
	}

	Sort(&result)
	fmt.Print("Sorted Array -> ")
	fmt.Println(result)
}

func ConcurrentSort(numbers []int, elements chan []int) {
	Sort(&numbers)
	elements <- numbers
	wg.Done()
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

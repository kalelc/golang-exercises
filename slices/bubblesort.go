package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers []int

	fmt.Print("Enter 10 integers separated by space to sort in order from least to greatest: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	elements := strings.Split(s, " ")

	for _, value := range elements {

		num, _ := strconv.Atoi(strings.Trim(value, " "))
		numbers = append(numbers, num)
	}

	BubbleSort(&numbers)
	fmt.Println(numbers)
}

func BubbleSort(numbers *[]int) {
	list := *numbers
	for i := 1; i < len(list); i++ {
		Swap(&list, i)
	}
}

func Swap(numbers *[]int, i int) {
	list := *numbers
	for j := 0; j < len(list)-i; j++ {
		if list[j] > list[j+1] {
			temp := list[j]
			list[j] = list[j+1]
			list[j+1] = temp
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter an array of numbers separated by spaces: ")
	numbers, err := ReadArrayOfIntegers(reader)

	if err != nil {
		panic(err)
	}

	BubbleSort(numbers)
	fmt.Println(numbers)
}

func ReadArrayOfIntegers(reader *bufio.Reader) (numbers []int, err error) {
	input := ""

	// read the string from console
	input, err = reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// format the string
	input = strings.TrimSuffix(input, "\n")
	input = strings.Trim(input, " ")

	// convert the array of strings to array of integers
	arrInt := []int{}

	for _, str := range regexp.MustCompile("\\s+").Split(input, -1) {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		arrInt = append(arrInt, num)
	}

	return arrInt, nil
}

func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				temp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
			}
		}
	}
}

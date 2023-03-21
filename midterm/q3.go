package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter an array of numbers separated by spaces: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Could not read string from console: ", err)
	}

	// format the string
	input = strings.TrimSuffix(input, "\n")
	input = strings.Trim(input, " ")

	// convert the array of strings to array of integers
	arrInt := []int{}

	for _, str := range regexp.MustCompile("\\s+").Split(input, -1) {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal("Could not convert a string to an integer: ", err)
		}
		arrInt = append(arrInt, num)
	}

	arrBool := []bool{}
	for i := 0; i < len(arrInt); i++ {
		arrBool = append(arrBool, isPrime(arrInt[i]))
	}

	fmt.Println(arrBool)
}

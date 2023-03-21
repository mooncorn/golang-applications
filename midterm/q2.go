package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func printSequence(num int) {
	if num < 1 {
		return
	}

	for i := 0; i < num; i++ {
		fmt.Print(math.Pow(2, float64(i)))
		if i < num-1 {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter an integer: ")
	// read the string from console
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// format the string
	input = strings.TrimSuffix(input, "\n")
	input = strings.Trim(input, " ")

	// convert to int
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Could not convert to string: ", input)
	}

	printSequence(num)
}

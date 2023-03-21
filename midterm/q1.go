package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func reverseString(str *string) {
	runes := []rune(*str)
	for i, j := 0, len(*str)-1; i < len(*str)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*str = string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string to be reversed: ")
	// read the string from console
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	reverseString(&input)
	fmt.Println("Reversed: ", input)
}

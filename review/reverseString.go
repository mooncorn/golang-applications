package main

import "fmt"

func reverseString(str *string) {
	runes := []rune(*str)
	for i, j := 0, len(*str)-1; i < len(*str)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*str = string(runes)
}

func main() {
	str := "hello world"
	reverseString(&str)
	fmt.Println(str)
}

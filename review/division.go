package main

import (
	"fmt"
)

func main() {
	res, err := Divide(1, 0)
	if err != nil {
		fmt.Println("Exception happened")
	}
	fmt.Println("res: ", res)
}

func Divide(num1 float64, num2 float64) (float64, error) {
	if num2 == 0 {
		panic("Division by 0")
		// log.Fatal("Division by 0")
	}
	return num1 / num2, nil
}

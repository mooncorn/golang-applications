package main

import (
	"fmt"
)

func main() {
	arr := []float64{1, 2, 5, 654, 3, 64, 4, 63, -2}
	fmt.Println(Max(arr))
	fmt.Println(Min(arr))
}

func Max(numbers []float64) (max float64) {
	max = numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func Min(numbers []float64) (min float64) {
	min = numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

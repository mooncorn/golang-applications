package main

import (
	src "calculator/src"
)

func main() {
	calculator := src.NewCalculator()
	ui := src.NewUI(calculator)
	ui.Run()
}

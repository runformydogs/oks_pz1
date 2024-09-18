package main

import "fmt"

// Версия 0.0.1.1 от 18.09.2024

type CCalculator struct {
	x, y int
}

func NewCCalculator(x, y int) *CCalculator {
	return &CCalculator{x: x, y: y}
}

func (c *CCalculator) Sum() int {
	return c.x + c.y
}

func (c *CCalculator) Multiple() int {
	return c.x * c.y
}

func main() {
	myCalc := NewCCalculator(3, 2)

	fmt.Println("Сумма:", myCalc.Sum())
	fmt.Println("Произведение:", myCalc.Multiple())
}

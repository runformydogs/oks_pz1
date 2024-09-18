package main

import "fmt"

// Версия 0.0.1. от 18.09.2024

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

func (c *CCalculator) Dec() int {
	return c.x - c.y
}

func main() {
	x := 3
	y := 2 // добавил инициализацию переменных и вывод, тк фиксить нечего

	myCalc := NewCCalculator(x, y)

	fmt.Println("Сумма:", myCalc.Sum())
	fmt.Println("Произведение:", myCalc.Multiple())
	fmt.Println("Разность", myCalc.Dec())
}

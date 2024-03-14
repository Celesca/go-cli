package main

import "fmt"

type Calculator struct {
	x float64
	y float64
}

func (c *Calculator) add() float64 {
	return c.x + c.y
}

func (c *Calculator) sub() float64 {
	return c.x - c.y
}

func main() {
	var x, y float64
	fmt.Println("Enter the first number")
	fmt.Scan(&x)
	fmt.Println("Enter the second number")
	fmt.Scan(&y)
	calculator := Calculator{x: x, y: y}
	fmt.Println("Addition: ", calculator.add())
	fmt.Println("Subtraction: ", calculator.sub())
}

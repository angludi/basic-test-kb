package main

import (
	"fmt"

	"command-pattern/command"
	"command-pattern/receiver"
)

func main() {
	// input := &receiver.Calc{
	// 	X: 1,
	// 	Y: 2,
	// 	N: 4,
	// }

	input := GetInput()

	sumCommand := &command.SumCommand{Calc: input}
	sumCommand.Calc.Sum()
	fmt.Println()

	multiplyCommand := &command.MultiplyCommand{Calc: input}
	multiplyCommand.Calc.Multiply()
	fmt.Println()

	primeCommand := &command.PrimeCommand{Calc: input}
	primeCommand.Calc.Prime()
	fmt.Println()

	fibonacciCommand := &command.FibonacciCommand{Calc: input}
	fibonacciCommand.Calc.Fibonacci()
}

func GetInput() *receiver.Calc {
	var x, y, n int

	fmt.Println("To calculate following statements;")
	fmt.Println("1. Sum X & Y")
	fmt.Println("2. Multiply X & Y")
	fmt.Println("3. Print first N prime number")
	fmt.Println("4. Print first N fibonacci sequence")
	fmt.Println()
	fmt.Println("Please input variable below:")
	fmt.Println()

	fmt.Print("X value: ")
	fmt.Scanln(&x)

	fmt.Print("Y value: ")
	fmt.Scanln(&y)

	fmt.Print("N value: ")
	fmt.Scanln(&n)

	fmt.Println()

	input := &receiver.Calc{
		X: x,
		Y: y,
		N: n,
	}

	return input
}

package main

import (
	"fmt"

	"github.com/angludi/tech-test-kb/command"
	"github.com/angludi/tech-test-kb/receiver"
)

func main() {
	input := &receiver.Calc{
		X: 1,
		Y: 2,
		N: 4,
	}

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

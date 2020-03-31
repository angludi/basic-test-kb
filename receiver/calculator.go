package receiver

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Calc struct {
	X int
	Y int
	N int
}

func (c *Calc) Sum() {
	total := c.X + c.Y
	fmt.Println("1. Sum X & Y")
	fmt.Printf("Input: %d, %d\n", c.X, c.Y)
	fmt.Println("Output:", total)
}

func (c *Calc) Multiply() {
	total := c.X * c.Y
	fmt.Println("2. Multiply X & Y")
	fmt.Printf("Input: %d, %d\n", c.X, c.Y)
	fmt.Println("Output:", total)
}

func (c *Calc) Prime() {
	var prime []int

	number := 1
	count := 1

	for count <= c.N {
		if IsPrime(number) {
			prime = append(prime, number)
			count++
		}

		number += 1
	}

	fmt.Println("First N Prime Number")
	fmt.Println("Input:", c.N)
	fmt.Println("Output:", JoinSlice(prime, ","))
}

func (c *Calc) Fibonacci() {
	var fibonacci []int
	for i := 0; i < c.N; i++ {
		fibonacci = append(fibonacci, FibonacciLoop(i))
	}

	fmt.Println("First N Fibonacci Sequence")
	fmt.Println("Input:", c.N)
	fmt.Println("Output:", JoinSlice(fibonacci, ","))
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func FibonacciLoop(n int) int {
	if n <= 1 {
		return n
	}

	var n1, n2 int = 0, 1

	for i := 2; i < n; i++ {
		n1, n2 = n2, n1+n2
	}

	return n1 + n2
}

func JoinSlice(items []int, sep string) (stringSlice string) {
	if len(items) == 1 {
		return strconv.Itoa(items[0])
	}

	var res []string
	for _, item := range items {
		res = append(res, strconv.Itoa(item))
	}
	return strings.Join(res, ",")
}

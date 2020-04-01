package command

import (
	"command-pattern/receiver"
)

type PrimeCommand struct {
	Calc receiver.ICalc
}

func (c *PrimeCommand) Execute() {
	c.Calc.Prime()
}

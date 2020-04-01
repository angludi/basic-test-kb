package command

import (
	"github.com/angludi/command-pattern/receiver"
)

type PrimeCommand struct {
	Calc receiver.ICalc
}

func (c *PrimeCommand) Execute() {
	c.Calc.Prime()
}

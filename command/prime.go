package command

import (
	"github.com/angludi/tech-test-kb/receiver"
)

type PrimeCommand struct {
	Calc receiver.ICalc
}

func (c *PrimeCommand) Execute() {
	c.Calc.Prime()
}

package invoker

import (
	"github.com/angludi/command-pattern/command"
)

type Process struct {
	Command command.ICommand
}

func (p *Process) process() {
	p.Command.Execute()
}

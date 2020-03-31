package invoker

import (
	"github.com/angludi/tech-test-kb/command"
)

type Process struct {
	Command command.ICommand
}

func (p *Process) process() {
	p.Command.Execute()
}

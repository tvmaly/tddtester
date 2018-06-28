package runner

type Command struct {
	command string
}

func NewCommand(cmd string) *Command {

	return &Command{command: cmd}

}

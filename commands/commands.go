package commands

import (
	"os"
	"os/exec"
	"strings"
)

const Docker = "docker"
const Placeholder = "$"

type Subcommand struct {
}

// Run picks subcommand and run it with args
func Run(sub string, args []string) error {
	cmd := new(Subcommand)
	var err error
	switch sub {
	case "ps":
		err = cmd.PS(args)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *Subcommand) CommandWithStd(name string, arg ...string) *exec.Cmd {
	cmd := exec.Command(Docker, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd
}

func (s *Subcommand) templifyArgs(template string, value string) []string {
	res := strings.Replace(template, Placeholder, value, 1)
	return strings.Split(res, " ")
}

func (s Subcommand) SplitCheckEmpty(in string) []string {
	lines := strings.Split(in, "\n")
	for k, v := range lines {
		lines[k] = strings.TrimSpace(v)
	}
	if len(lines) == 1 && lines[0] == "" {
		return nil
	}
	return lines
}

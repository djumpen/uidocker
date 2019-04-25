package commands

import (
	"fmt"
	"github.com/djumpen/console-probe/surv"
	"github.com/spf13/viper"
	survey "gopkg.in/AlecAivazis/survey.v1"
	"os/exec"
	"regexp"
	"strings"
)

func (s *Subcommand) PS(args []string) error {
	format := viper.GetString("defaults.ps.format")
	cmd := exec.Command(Docker, "ps", "--format", format) // TODO: add support args
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	outStr := strings.TrimSpace(string(out))

	containers := s.SplitCheckEmpty(outStr)
	if containers == nil {
		fmt.Println("No containers")
		return nil
	}

	qs, err := surv.NewQuestion("Choose a container:", containers).Ask()
	if err != nil {
		return err
	}
	containerID := s.parseContainerID(qs.Answer())

	defaultAction := "[Type template manually]"
	configActions := viper.GetStringSlice("templates.ps.single")
	actions := append([]string{defaultAction}, configActions...)

	qs2, err := surv.NewQuestion("Choose action:", actions).Ask()
	if err != nil {
		return err
	}

	action := qs2.Answer()
	if action == defaultAction {
		survey.AskOne(&survey.Input{}, &action, nil)
	}

	actionArgs := s.templifyArgs(action, containerID)

	cmd = s.CommandWithStd(Docker, actionArgs...)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("cmd.Run() failed with %s\n", err)
	}

	return nil
}

func (s *Subcommand) parseContainerID(input string) string {
	var re = regexp.MustCompile(`[0-9a-z]{12,12}`)
	return re.FindString(input)
}

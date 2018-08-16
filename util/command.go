package util

import (
	"errors"
	"os/exec"
	"strconv"
)

type Cmd struct {
	Name string
	Args []string
}

func RunCmds(cmds []Cmd) (logs []string, err error) {
	for idx, command := range cmds {
		cmd := exec.Command(command.Name, command.Args...)
		out, err := cmd.Output()
		if err != nil {
			err = errors.New(strconv.Itoa(idx) + "行command错误")
			return nil, err
		}
		logs = append(logs, string(out))
	}
	return logs, nil
}

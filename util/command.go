package util

import "os/exec"

type Cmd struct {
	Name string
	Args []string
}

func RunCmds(cmds []Cmd) (logs []string, err error) {
	for _, command := range cmds {
		cmd := exec.Command(command.Name, command.Args...)
		out, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		logs = append(logs, string(out))
	}
	return logs, nil
}

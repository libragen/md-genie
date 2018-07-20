package main

import (
	"os/exec"
	_ "strconv"
	"time"

	"github.com/mojocn/movie-board/util"
)

func main() {
	for {
		util.FetchMaoyanApi()
		util.ParseMarkdown()
		runGitCmds()
		time.Sleep(6 * time.Hour)
	}
}

func runGitCmds() {
	commitMsg := time.Now().Format(time.RFC3339)
	cmds := [][]string{
		[]string{"stash"},
		[]string{"pull", "origin", "master"},
		[]string{"stash", "apply"},
		[]string{"add", "."},
		[]string{"commit", "-am", commitMsg},
		[]string{"push", "origin", "master"},
	}
	for _, arguments := range cmds {
		gitCommand(arguments...)
	}
}

func gitCommand(args ...string) {
	app := "git"
	cmd := exec.Command(app, args...)
	out, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
	util.DingLog("Git日志", string(out))
}

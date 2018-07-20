package main

import (
	"github.com/mojocn/movie-board/util"
	"os/exec"
	"time"
)

func main() {
	for {
		util.SpiderHackNews()
		util.ParseMarkdownHacknews()

		util.FetchMaoyanApi()
		util.ParseMarkdown()
		runGitCmds()
		time.Sleep(3 * time.Hour)
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
	util.DingLog(string(out), "Git日志")
}

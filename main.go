package main

import (
	"github.com/dejavuzhou/md-genie/util"
	"os/exec"
	"time"
)

func main() {
	for {
		util.SpiderHackNews()
		util.ParseMarkdownHacknews()

		util.FetchMaoyanApi()
		util.ParseMaoyanMarkdown()

		util.ParseReadmeMarkdown()
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
	var outLog string

	for _, arguments := range cmds {
		out := gitCommand(arguments...)
		outLog += out
	}

	//util.DingLog(string(outLog), "Git日志")
	subject := "Git日志:" + commitMsg
	util.SendMsgToEmail(subject,string(outLog),"erikchau@me.com")

}

func gitCommand(args ...string)string {
	app := "git"
	cmd := exec.Command(app, args...)
	out, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(out)
}

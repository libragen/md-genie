package main

import (
	"fmt"
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
		mailTitle, mailBody := runGitCmds()

		go util.SendMsgToEmail(mailTitle, mailBody, "erikchau@me.com")
		time.Sleep(6 * time.Hour)

	}
}

func runGitCmds() (string, string) {
	commitMsg := time.Now().Format(time.RFC3339)
	cmds := [][]string{
		{"stash"},
		{"pull", "origin", "master"},
		{"stash", "apply"},
		{"add", "."},
		//{"merge", "--strategy-option","ours"},
		{"commit", "-am", commitMsg},
		{"push", "origin", "master"},
	}
	var mailBody string

	for _, arguments := range cmds {
		out := gitCommand(arguments...)
		mailBody += fmt.Sprintf("<p>%s</p>", out)
	}
	//util.DingLog(string(mailBody), "Git日志")
	mailTitle := "Git日志:" + commitMsg
	return mailTitle, mailBody
}

func gitCommand(args ...string) string {
	app := "git"
	cmd := exec.Command(app, args...)
	out, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(out)
}

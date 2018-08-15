package main

import (
	"fmt"
	"github.com/dejavuzhou/md-genie/util"
	"os/exec"
	"time"
)

func main() {
	for {
		if err := util.SpiderHackNews(); err != nil {
			fmt.Println(err)
		}
		if err := util.ParseMarkdownHacknews(); err != nil {
			fmt.Println(err)
		}

		if err := util.FetchMaoyanApi(); err != nil {
			fmt.Println(err)
		}
		if err := util.ParseMaoyanMarkdown(); err != nil {
			fmt.Println(err)
		}

		util.ParseReadmeMarkdown()
		mailTitle, gitlogs := runGitCmds()
		if err, mailBody := util.ParseEmailContent(gitlogs); err == nil {
			util.SendMsgToEmail(mailTitle, mailBody)
		} else {
			fmt.Println(err)
		}
		time.Sleep(6 * time.Hour)
	}
}

func runGitCmds() (string, []string) {
	commitMsg := time.Now().Format("2006年01月02日15点04分")
	cmds := [][]string{
		{"stash"},
		{"pull", "origin", "master"},
		{"stash", "apply"},
		{"add", "."},
		//{"merge", "--strategy-option","ours"},
		{"commit", "-am", commitMsg},
		{"push", "origin", "master"},
	}
	var gitlogs []string

	for _, arguments := range cmds {
		out := gitCommand(arguments...)
		gitlogs = append(gitlogs, out)
	}
	//util.DingLog(string(gitlogs), "Git日志")
	mailTitle := "每日新闻" + commitMsg
	return mailTitle, gitlogs
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

package main

import (
	"github.com/dejavuzhou/md-genie/util"
	"os/exec"
	"time"
)

func main() {
	for {
		if err := util.SpiderHackNews(); err != nil {
			util.SendMsgToEmail("spider hack news error", err.Error(), "erikchau@me.com")
		}
		if err := util.ParseMarkdownHacknews(); err != nil {
			util.SendMsgToEmail("pasrse hack news markdown error", err.Error(), "erikchau@me.com")
		}

		if err := util.FetchMaoyanApi(); err != nil {
			util.SendMsgToEmail("fetch maoyan api error", err.Error(), "erikchau@me.com")
		}
		if err := util.ParseMaoyanMarkdown(); err != nil {
			util.SendMsgToEmail("parse maoyan movie markdown error", err.Error(), "erikchau@me.com")
		}

		util.ParseReadmeMarkdown()
		mailTitle, gitlogs := runGitCmds()
		if err, mailBody := util.ParseEmailContent(gitlogs); err == nil {
			util.SendMsgToEmail(mailTitle, mailBody, "erikchau@me.com")
		} else {
			util.SendMsgToEmail("parse email content hmtl error", err.Error(), "erikchau@me.com")
		}
		time.Sleep(6 * time.Hour)

	}
}

func runGitCmds() (string, []string) {
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
	var gitlogs []string

	for _, arguments := range cmds {
		out := gitCommand(arguments...)
		gitlogs = append(gitlogs, out)
	}
	//util.DingLog(string(gitlogs), "Git日志")
	mailTitle := "Git日志:" + commitMsg
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

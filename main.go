package main

import (
	"github.com/dejavuzhou/md-genie/util"
	"log"
	"time"
)

var cmds = []util.Cmd{
	{"git", []string{"stash"}},
	{"git", []string{"pull", "origin", "master"}},
	{"git", []string{"stash", "apply"}},
	{"git", []string{"add", "."}},
	{"git", []string{"stash"}},
	{"git", []string{"commit", "-am", time.Now().Format(time.RFC3339)}},
	{"ps", []string{"ps", "-ef", "|", "grep", "md-genie"}},
	{"netstat", []string{"-lntp"}},
	{"free", []string{"-m"}},
	{"ps", []string{"aux"}},
}

func main() {
	for {
		if err := util.SpiderHackNews(); err != nil {
			log.Fatal(err)
		}
		if err := util.ParseMarkdownHacknews(); err != nil {
			log.Fatal(err)
		}

		if err := util.FetchMaoyanApi(); err != nil {
			log.Fatal(err)
		}
		if err := util.ParseMaoyanMarkdown(); err != nil {
			log.Fatal(err)
		}
		util.ParseReadmeMarkdown()

		gitlogs, err := util.RunCmds(cmds)
		if err != nil {
			log.Fatal(err)
		}
		if err, mailBody := util.ParseEmailContent(gitlogs); err == nil {
			mailTitle := "md-genie+hacknews日志:" + time.Now().Format(time.RFC3339)
			util.SendMsgToEmail(mailTitle, mailBody)
		} else {
			log.Fatal(err)
		}
		time.Sleep(6 * time.Hour)
	}
}

package main
import (

	"github.com/dejavuzhou/md-genie/util"
		"testing"
)

/*
	有道智云
	文档页面  http://ai.youdao.com/docs/doc-trans-api.s#p01
*/
func Test_main(t *testing.T) {

	if err := util.SpiderHackNews();err != nil {
		util.SendMsgToEmail("spider hack news error", err.Error(), "erikchau@me.com")
	}
	if err := util.ParseMarkdownHacknews();err != nil {
		util.SendMsgToEmail("pasrse hack news markdown error", err.Error(), "erikchau@me.com")
	}

	if err := util.FetchMaoyanApi();err != nil {
		util.SendMsgToEmail("fetch maoyan api error", err.Error(), "erikchau@me.com")
	}
	if err := util.ParseMaoyanMarkdown();err != nil {
		util.SendMsgToEmail("parse maoyan movie markdown error", err.Error(), "erikchau@me.com")
	}

	util.ParseReadmeMarkdown()
	mailTitle, gitlog := runGitCmds()

	if err,mailBody := util.ParseEmailContent(gitlog);err == nil {
		util.SendMsgToEmail(mailTitle, mailBody, "erikchau@me.com")
	}else {
		util.SendMsgToEmail("parse email content hmtl error", err.Error(), "erikchau@me.com")
	}
}


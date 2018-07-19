package main

import (
	"os/exec"
	_ "strconv"
	"time"

	"github.com/mojocn/movie-board/util"
	"fmt"
)





func main() {
	//loop
	//for {
		util.FetchMaoyanApi()
		util.ParseMarkdown()

		runGitCmds()
	//	time.Sleep(3 * time.Hour)
	//}
}


func runGitCmds(){
	commitMsg := time.Now().Format(time.RFC3339)

	cmds := [][]string{
		[]string{"stash"},
		[]string{"pull","origin","master"},
		[]string{"stash","apply"},
		[]string{"add","archives/*"},
		[]string{"commit","-am",commitMsg},
		[]string{"push","origin","master"},
	}
	for _,arguments := range cmds {
		gitCommand(arguments...)
	}
}




func gitCommand(args ...string) {
	app := "git"
	cmd := exec.Command(app,args ...)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Print(string(out))
}

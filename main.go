package main

import (
	"github.com/acierto/go-jira-client"
	"os"
	"fmt"
	"flag"
	"log"
	"runtime"
	"os/exec"
	"time"
)

var list_of_dirs = []string{"conf/", "ext/", "repository/"}

var jira *gojira.Jira

func init() {
	jira = GetJira()
}

func main() {

	issueParam := flag.String("issue", "", "Specify your Jira issue, i.e. -issue=DEPL-6501")
	exportParam := flag.Bool("export", true, "By default you are exporting")
	forceParam := flag.Bool("force", false, "Export XLA package and replace the previous uploaded package")
	importParam := flag.Bool("import", false, "Imports the data for specified issue")
	restartParam := flag.Bool("restart", false, "Restart the server after importing the XLA")

	flag.Parse()

	if *issueParam == "" {
		log.Fatal("Please provide the issue number. I.e. xl-aide -issue=DEPL-6501")
		os.Exit(1)
	}

	if *importParam {
		importXlaArchive(*issueParam)
		if *restartParam {
			restartXlDeploy()
		}
	} else if *forceParam {
		exportXlaArchive(*issueParam, true)
	} else if *exportParam {
		exportXlaArchive(*issueParam, false)
	}
}

func getXldLocation() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd + string(os.PathSeparator)
}

func restartXlDeploy() {
	println("Sent request to restart XLD server")
	stopXlDeploy()
	time.Sleep(time.Second * 5)
	startXlDeploy()
}

func stopXlDeploy() {
	url := fmt.Sprintf("http://localhost:%s/deployit/server/shutdown", ReadXldAppKey("http.port"))
	FireAndForget("POST", url)
}

func startXlDeploy() {
	if runtime.GOOS == "windows" {
		exec.Command("cmd", "/C", "bin\\server.cmd").Start()
	} else {
		fmt.Println("bin/server.sh")
		exec.Command("sh","-c", "bin/server.sh").Start()
	}
}

func logErrorCleanAndExit(attachmentPath string, err error) {
	os.RemoveAll(attachmentPath)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}

package main

import (
	"github.com/twinj/uuid"
	"github.com/acierto/archivex"
	"github.com/acierto/unzipit"
	"github.com/acierto/go-jira-client"
	"os"
	"fmt"
	"flag"
	"log"
	"runtime"
	"os/exec"
	"time"
)

const (
	archive_name = "xla-snapshot"
	full_archive_name = archive_name + ".zip"
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

func createArchive() string {

	xld_location := getXldLocation()

	arc := archivex.ZipFile{}

	arc.Create(archive_name)

	//	arc.AddAll(xld_location+"plugins/", true) TODO: should read the plugins to metadata of archive
	for _,dir := range list_of_dirs {
		arc.AddAll(xld_location+dir, true)
	}

	arc.Close()

	return xld_location + full_archive_name
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

func exportXlaArchive(issueId string, replace bool) {
	attachmentPath := createArchive()

	if replace {
		err := jira.UpdateAttachment(issueId, attachmentPath)
		logErrorCleanAndExit(attachmentPath, err)
	} else {
		err := jira.AddAttachment(issueId, attachmentPath)
		logErrorCleanAndExit(attachmentPath, err)
	}

	log.Printf("XLA attachment [%s] has been successfully uploaded.", attachmentPath)
}

func logErrorCleanAndExit(attachmentPath string, err error) {
	os.RemoveAll(attachmentPath)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}

func importXlaArchive(issueKey string) {
	attachmentPath,_ := jira.DownloadAttachment(issueKey, full_archive_name)

	if attachmentPath == "" {
		log.Printf("Nothing to import. XLA attachment for issue [%s] has not been found.", issueKey)
		os.Exit(1)
	}

	attachment, _ := os.Open(attachmentPath)
	defer attachment.Close()

	unzipit.Unpack(attachment, archive_name)
	os.RemoveAll(attachmentPath)

	for _,dir := range list_of_dirs {
		os.RemoveAll(dir)
		CopyDir(archive_name + string(os.PathSeparator) + dir, dir)
	}

	os.RemoveAll(archive_name)

	log.Print("XLA attachment has been successfully imported.")
}

func generateId() string {
	u := uuid.NewV4()
	return uuid.Formatter(u, uuid.CleanHyphen)
}

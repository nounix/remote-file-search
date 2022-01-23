package main

import (
	"log"
	"os/exec"
	"os/user"
	"strings"
)

func logErr(err error) {
	if err != nil {
		exec.Command("zenity", "--error", "--text", err.Error()).Run()
		log.Fatal(err)
	}
}

func getCmdOutput(name string, arg ...string) string {
	out, err := exec.Command(name, arg...).Output()
	logErr(err)
	if len(out) > 0 {
		return string(out[:len(out)-1])
	}
	return ""
}

func searchList(str []string, substr string) string {
	for _, v := range str {
		if strings.Contains(v, substr) {
			return v
		}
	}
	return ""
}

func main() {
	usr, _ := user.Current()
	iSearchDir := searchList(strings.Split(getCmdOutput("ls", "/run/user/"+usr.Uid+"/gvfs/"), "\n"), "linux")
	iSearchFrontend := "/run/user/" + usr.Uid + "/gvfs/" + iSearchDir + "/iSearch/frontend"
	iSearch := "/run/user/" + usr.Uid + "/gvfs/" + iSearchDir + "/iSearch/iSearch.sh"
	getCmdOutput("cp", iSearch, usr.HomeDir)
	getCmdOutput("chmod", "755", usr.HomeDir+"/iSearch.sh")
	getCmdOutput("cp", iSearchFrontend, "/tmp/iSearch-frontend")
	getCmdOutput("chmod", "755", "/tmp/iSearch-frontend")
	getCmdOutput("gksudo", "cp", "/tmp/iSearch-frontend", "/usr/local/sbin/iSearch-frontend")
	getCmdOutput("gksudo", "apt install -y zenity")
	exec.Command("zenity", "--info", "--text", "Successfully installed!").Run()
}

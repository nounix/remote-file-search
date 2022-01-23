package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func logErr(err error) {
	if err != nil {
		log.Println(err)
		exec.Command("zenity", "--error", "--text", err.Error()).Run()
	}
}

func httpPost(url string, data map[string]string) []string {
	jData, err := json.Marshal(data)
	logErr(err)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jData))
	logErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	logErr(err)
	str := strings.Split(string(body), "\n")
	return str
}

func getCmdOutput(name string, arg ...string) string {
	out, err := exec.Command(name, arg...).Output()
	logErr(err)
	if len(out) > 0 {
		return string(out[:len(out)-1])
	}
	return ""
}

func strContainsMul(str string, substr ...string) bool {
	for _, v := range substr {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

func rmDirs(dirs []string, substr ...string) []string {
	var clearDir []string
	for _, v := range dirs {
		if !strContainsMul(v, substr...) {
			// dirs = append(dirs[:i], dirs[i+1:]...)
			clearDir = append(clearDir, v)
		}
	}
	return clearDir
}

func main() {
	ip := os.Args[1]
	localMountDir := os.Args[2]
	startSearchDir := os.Args[3]
	hideDirs := strings.Split(os.Args[4], ";")
	fileManager := os.Args[5]

	path := startSearchDir
	resp := "init"

	for resp != "" {
		dirs := httpPost("http://"+ip+":8989/lsDir", map[string]string{"Dir": path})
		dirs = append([]string{"..UP"}, dirs...)
		resp = getCmdOutput("zenity", append([]string{"--height=400", "--width=600", "--list", "--title=Directory", "--column=Directory"}, rmDirs(dirs, hideDirs...)...)...)
		if resp == "..UP" {
			path = path[:strings.LastIndex(path, "/")]
		}
		if resp != "" && resp != "..UP" {
			path = resp
		}
	}

	sk := getCmdOutput("zenity", "--entry", "--text=Search for:")
	files := httpPost("http://"+ip+":8989/searchFiles", map[string]string{"Dir": path, "SearchKey": sk})
	selDir := getCmdOutput("zenity", append([]string{"--height=400", "--width=600", "--list", "--title=Searching", "--column=Found"}, files...)...)
	getCmdOutput(fileManager, strings.Replace(selDir, startSearchDir, localMountDir, 1))
}

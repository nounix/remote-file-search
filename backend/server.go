package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func readJSON(req *http.Request, v interface{}) {
	decoder := json.NewDecoder(req.Body)
	logErr(decoder.Decode(&v))
	defer req.Body.Close()
}

func getCmdOutput(name string, arg ...string) string {
	out, err := exec.Command(name, arg...).Output()
	logErr(err)
	if len(out) > 0 {
		return string(out[:len(out)-1])
	}
	return ""
}

func lsDir(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Dir string
	}{}
	readJSON(r, &req)
	out := getCmdOutput("find", req.Dir, "-maxdepth", "1", "-type", "d")
	fmt.Fprintf(w, out)
}

func searchFiles(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Dir       string
		SearchKey string
	}{}
	readJSON(r, &req)
	out := getCmdOutput("find", req.Dir, "-name", "*"+req.SearchKey+"*")
	fmt.Fprintf(w, out)
}

func main() {
	http.HandleFunc("/searchFiles", searchFiles)
	http.HandleFunc("/lsDir", lsDir)
	log.Fatal(http.ListenAndServe(":8989", nil))
}

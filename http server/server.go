package main

import (
	"fmt"
	//"fmt"
	"bufio"
	"encoding/json"
	//"io/ioutil"
	"net/http"
	"os"
	//"strings"
)

type result struct {
	Msg  string
	Code int
}

var res result = result{"success", 200}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func indexOf(arrString []string, word string) (int, bool) {
	for i, e := range arrString {
		if e == word {
			return i, true
		}
	}
	return -1, false
}

func SetBlackUid(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "setblackuid!\n")
	fileName := "./blackuid.txt"
	uid := r.FormValue("uid")
	lines, err := readLines(fileName)
	if err != nil {
		res.Code = 500
		res.Msg = "internal error!"
		jsonRes, _ := json.Marshal(res)
		w.Write([]byte(jsonRes))
		return
	}
	fmt.Println(lines)

	_, ok := indexOf(lines, uid)
	if ok == false {
		lines = append(lines, uid)
	}
	err = writeLines(lines, fileName)
	if err != nil {
		res.Code = 500
		res.Msg = "internal error!"
		jsonRes, _ := json.Marshal(res)
		w.Write([]byte(jsonRes))
		return
	}
	resOut, _ := json.Marshal(res)
	w.Write([]byte(resOut))
}

func main() {
	http.HandleFunc("/setblackuid", SetBlackUid)
	http.ListenAndServe(":8000", nil)
}

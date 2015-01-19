package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var noticeLogFileName string = "./go.log"
var errorLogFileName string = "./error.log"
var noticeLogger *log.Logger
var errorLogger *log.Logger

type retMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func main() {
	noticeLogfile, err := os.OpenFile(noticeLogFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer noticeLogfile.Close()

	errorLogfile, err := os.OpenFile(errorLogFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer errorLogfile.Close()

	noticeLogger = log.New(noticeLogfile, "NOTICE: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(errorLogfile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	fileName := "./error.log"
	re := regexp.MustCompile("FAIL:(.*):(.*)")
	var mutex = &sync.Mutex{}
	mutex.Lock()
	lines, err := readLines(fileName)
	if err != nil {

	}

	//writeLines([]string{}, fileName)
	mutex.Unlock()

	for _, line := range lines {
		fmt.Println(line)

		r2 := re.FindAllStringSubmatch(line, -1)[0]
		fmt.Println(r2)
	}

}

func flushCbaseCache(server, mid string) {
	action := "delete"
	key := "xxxxxxxxx"
	tm := strconv.FormatInt(time.Now().Unix(), 10)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(tm + action + key))
	sigSum := md5Ctx.Sum(nil)
	sig := hex.EncodeToString(sigSum)
	url := "http://" + server + "/tool?tm=" + tm + "&action=" + action + "&sig=" + sig + "&key=MI" + mid
	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		errorLogger.Println("FAIL:" + server + ":" + mid)
		return
	}
	var res retMsg
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	err = json.Unmarshal(body, &res)
	if err != nil {

	}
	if res.Code != 200 {
		//logToFile("FAIL:" + server + ":" + mid)
		errorLogger.Println("FAIL:" + server + ":" + mid)
		return
	}
	//logToFile("OK:" + server + ":" + mid)
	noticeLogger.Println("OK:" + server + ":" + mid)
	fmt.Println(res)

	return
}

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

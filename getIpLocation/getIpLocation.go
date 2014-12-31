package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	//"time"
)

func main() {
	fileName := "./tmp.txt"
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	urls := strings.Split(string(fileData), "\n")
	done := make(chan string)

	for _, url := range urls {
		fmt.Println(url)
		url = strings.Trim(url, "\n\r")
		go getIpLocation(done, url)
	}
	lenUrl := len(urls)
	count := 0
	for {
		<-done
		count++
		if count >= lenUrl {
			break
		}

	}

}

func getIpLocation(done chan string, ip string) {
	res, err := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip)
	if err != nil {
		log.Fatal(err)
	}

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)
	defer func() { done <- "done" }()
}

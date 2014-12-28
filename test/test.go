package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileName := "./tmp.txt"
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(fileData))
	urls := strings.Split(string(fileData), "\n")
	for _, url := range urls {
		fmt.Println(url)
		go getIpLocation(url)
	}
}

func getIpLocation(ip string) {

}

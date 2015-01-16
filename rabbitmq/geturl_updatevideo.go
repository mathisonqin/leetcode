package main

import (
	"crypto/md5"
	"encoding/json"
	//"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	//"net/http"
	"encoding/hex"
	"io/ioutil"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"time"
)

var noticeLogFileName string = "/xxx/logs/geturl_go.log"
var errorLogFileName string = "/xxx/logs/geturl_go_error.log"
var noticeLogger *log.Logger
var errorLogger *log.Logger

func flushRedisCache(mid string) error {
	c, err := redis.Dial("tcp", "xxx.xxx.xxx.xxx:7379")
	if err != nil {
		// handle error
		fmt.Println(err)

	}
	key := "V2:MI:" + mid

	_, err = c.Do("DEL", key)
	if err != nil { //del fail
		//logToFile("FAIL_REDIS:" + mid)
		errorLogger.Println("FAIL_REDIS:" + mid)
	}
	defer c.Close()
	return nil
}

type retMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func logToFile(text string) {
	date := time.Now().Format("2006-01-02 15:04:05")
	text = "[" + date + "]" + text + "\n"
	f, err := os.OpenFile(noticeLogFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func flushCbaseCache(mid string) error {
	servers := []string{
		"10.10.10.10",
	}

	action := "delete"
	key := "xxxxxxxxx"
	for _, server := range servers {
		go func(server string) {
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
		}(server)

	}
	return nil
}

func main() {
	//init log file

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

	// Connects opens an AMQP connection from the credentials in the URL.
	//var amqpConfig amqp.URI
	amqpConfig := amqp.URI{
		Host:     "xxx.xxx.xxx.xxx",
		Port:     5673,
		Username: "test",
		Password: "test", //need urlencode
		Vhost:    "test",
		Scheme:   "amqp",
	}

	conn, err := amqp.Dial(amqpConfig.String())
	if err != nil {
		log.Fatalf("connection.open: %s", err)
	}
	defer conn.Close()

	c, err := conn.Channel()
	if err != nil {
		log.Fatalf("channel.open: %s", err)
	}

	// We declare our topology on both the publisher and consumer to ensure they
	// are the same.  This is part of AMQP being a programmable messaging model.
	//
	// See the Channel.Publish example for the complimentary declare.
	err = c.ExchangeDeclare("update_video_ex", "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("exchange.declare: %s", err)
	}

	// Establish our queue topologies that we are responsible for
	type bind struct {
		queue string
		key   string
	}
	bindInfo := bind{"update_video_queue", "update_video_rt"}
	_, err = c.QueueDeclare(bindInfo.queue, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("queue.declare: %v", err)
	}

	err = c.QueueBind(bindInfo.queue, bindInfo.key, "update_video_test_ex", false, nil)
	if err != nil {
		log.Fatalf("queue.bind: %v", err)
	}

	// Set our quality of service.  Since we're sharing 3 consumers on the same
	// channel, we want at least 3 messages in flight.
	err = c.Qos(50, 0, false)
	if err != nil {
		log.Fatalf("basic.qos: %v", err)
	}

	// Establish our consumers that have different responsibilities.  Our first
	// two queues do not ack the messages on the server, so require to be acked
	// on the client.

	// This consumer requests that every message is acknowledged as soon as it's
	// delivered.

	firehose, err := c.Consume(bindInfo.queue, "go_consumer_mms", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("basic.consume: %v", err)
	}

	// To show how to process the items in parallel, we'll use a work pool.
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(work <-chan amqp.Delivery) {
			for msg := range work {
				// ... this consumer pulls from the firehose and doesn't need to acknowledge
				fmt.Printf("%s\n", msg.Body)

				mid := string(msg.Body)
				//logToFile("message:" + mid)
				noticeLogger.Println("message:" + mid)
				go flushRedisCache(mid)
				go flushCbaseCache(mid)
			}
		}(firehose)
	}

	// Wait until you're ready to finish, could be a signal handler here.
	//time.Sleep(1000 * time.Second)
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, os.Kill)

	s := <-exitChan
	fmt.Println("Got signal:", s)
	// Cancelling a consumer by name will finish the range and gracefully end the
	// goroutine
	err = c.Cancel("go_consumer_mms", false)
	if err != nil {
		log.Fatalf("basic.cancel: %v", err)
	}

	// deferred closing the Connection will also finish the consumer's ranges of
	// their delivery chans.  If you need every delivery to be processed, make
	// sure to wait for all consumers goroutines to finish before exiting your
	// process.
}

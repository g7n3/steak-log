package main

import (
	"log/syslog"
	"time"
	"sync"
	"strconv"
	"fmt"
)

var counter = 0

func writeLog(id string, group *sync.WaitGroup) {
	tagName := `mee_test`
	sysLog, err := syslog.Dial("tcp", "127.0.0.1:4321",
		syslog.LOG_INFO|syslog.LOG_ALERT, tagName)
	if err != nil {
		fmt.Println("Test failed, " + id)
		panic(err)
	}
	sysLog.Info("Silence is Golden!" + id)
	sysLog.Close()
	counter++
	group.Done()
}

func main() {
	start := time.Now()
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(100)

	for i:=0; i<100; i++ {
		go writeLog(strconv.Itoa(i), &waitGroup)
	}

	waitGroup.Wait()

	fmt.Println("Test finished, time cost" + string(time.Since(start)))
}

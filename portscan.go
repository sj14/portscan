package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

const max = 65535

var (
	wg             sync.WaitGroup
	timeoutFlag    *string
	hostFlag       *string
	listClosedFlag *bool
)

func main() {
	// Command Line Flags
	hostFlag = flag.String("host", "localhost", "the host to scan")
	timeoutFlag = flag.String("timeout", "1000ms", "timeout (e.g. 50ms or 1s)")
	startPortFlag := flag.Int("start", 80, "the lower end to scan")
	endPortFlag := flag.Int("end", -1, "the upper end to scan")
	pauseFlag := flag.String("pause", "1ms", "pause after each port scan (e.g. 5ms)")
	listClosedFlag = flag.Bool("closed", false, "list closed ports (true/false)")

	flag.Parse()

	// End port not set
	if *endPortFlag == -1 {
		endPortFlag = startPortFlag
	}

	if *startPortFlag < 0 || *startPortFlag > max {
		log.Fatalf("start port out of range (should be between 0 and %d)\n", max)
	}
	if *endPortFlag < 0 || *endPortFlag > max {
		log.Fatalf("end port out of range (should be between 0 and %d)\n", max)
	}
	if *endPortFlag < *startPortFlag {
		log.Fatal("end port should be higher than start port")
	}

	startTime := time.Now()

	for port := *startPortFlag; port <= *endPortFlag; port++ {
		wg.Add(1)
		go scan(port)
		pauseDuration, err := time.ParseDuration(*pauseFlag)
		if err != nil {
			log.Print("wrong time input format?\n", err)
		}
		time.Sleep(pauseDuration)
	}
	wg.Wait()
	scanDuration := time.Since(startTime)
	fmt.Printf("scan finished in %v", scanDuration)
}

func scan(port int) {
	portString := strconv.Itoa(port)
	timeoutDuration, err := time.ParseDuration(*timeoutFlag)

	if err != nil {
		log.Print("wrong time input format?\n", err)
	}

	conn, err := net.DialTimeout("tcp", *hostFlag+":"+portString, timeoutDuration)
	if err != nil {
		if *listClosedFlag == true {
			log.Println(err)
		}
	} else {
		conn.Close()
		fmt.Printf("open %d\n", port)
	}
	wg.Done()
}

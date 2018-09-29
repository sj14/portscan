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

var (
	wg             sync.WaitGroup
	timeoutFlag    *string
	hostFlag       *string
	startPortFlag  *int
	endPortFlag    *int
	pauseFlag      *string
	listClosedFlag *bool
)

func init() {
	const max = 65535

	// Command Line Flags
	hostFlag = flag.String("host", "localhost", "the host to scan")
	timeoutFlag = flag.String("timeout", "1000ms", "timeout (e.g. 50ms or 1s)")
	startPortFlag = flag.Int("start", 80, "the lower end to scan")
	endPortFlag = flag.Int("end", -1, "the upper end to scan")
	pauseFlag = flag.String("pause", "1ms", "pause after each scanned port (e.g. 5ms)")
	listClosedFlag = flag.Bool("closed", false, "list closed ports (true/false)")

	flag.Parse()

	// End port not set
	if *endPortFlag == -1 {
		endPortFlag = startPortFlag
	}

	if *startPortFlag < 1 || *startPortFlag > max {
		log.Fatalf("starting port out of range (should be between 1 and %d)\n", max)
	}
	if *endPortFlag < 1 || *endPortFlag > max {
		log.Fatalf("ending port out of range (should be between 1 and %d)\n", max)
	}
	if *endPortFlag < *startPortFlag {
		log.Fatalln("ending port must be greater than starting port")
	}
}

func main() {
	startTime := time.Now()

	pauseDuration, err := time.ParseDuration(*pauseFlag)
	if err != nil {
		log.Println(err)
	}

	for port := *startPortFlag; port <= *endPortFlag; port++ {
		wg.Add(1)
		go scan(*hostFlag, port, *timeoutFlag, *listClosedFlag)
		time.Sleep(pauseDuration)
	}

	wg.Wait()
	scanDuration := time.Since(startTime)
	fmt.Printf("scan finished in %v\n", scanDuration)
}

func scan(host string, port int, timeout string, listClosed bool) {
	defer wg.Done()

	portStr := strconv.Itoa(port)

	timeoutDuration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Println(err)
		return
	}

	conn, err := net.DialTimeout("tcp", host+":"+portStr, timeoutDuration)
	if err != nil {
		if listClosed {
			log.Println(err)
		}
		return
	}

	conn.Close()
	log.Printf("open %d\n", port)
}

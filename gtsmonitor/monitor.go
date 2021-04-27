package gtsmonitor

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// Monitor if your websites are running and runs forever
func MonitorWebsiteUptime(interval int, urlString ...string) {
	if urlString == nil {
		panic("please specify one or more URLs to query")
	}
	if interval == 0 || interval < 1 {
		panic("please specify an interval greater or equal to 1")
	}

	for _, url := range urlString {
		go checkStatus(url)
	}
	for range time.Tick(time.Duration(interval) * time.Second) {
		for _, url := range urlString {
			go checkStatus(url)
		}
	}
}

// Coordination MonitorWebsiteUptime()
// checkStatus performs a get request to the specified url and logs the status code or error
func checkStatus(url string) {
	url = strings.TrimSpace(url)
	resp, err := http.Get(url)
	if err != nil {
		if !(resp != nil && resp.StatusCode > 0) {
			log.Printf("%s - %v\n", url, err)
			return
		}
	}
	// check the status code
	log.Printf(" %s - %d\n", url, resp.StatusCode)
}

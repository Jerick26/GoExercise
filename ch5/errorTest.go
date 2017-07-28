package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

// NOTE: Five Error Handling Strategies

func main() {
	value, ok := cache.Lookup(key)
	if !ok {
		// ...cache[key] does not exist...
	}
	fmt.Println(err) // or fmt.Printf("%v", err)
	// NOTE: First, to propagate the error
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	// NOTE: Third, the caller prints the error and stop the program gracefully
	// this course of action should generally be reservered for the main package
	// of a program.
	// (In function main.)
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
		// a more convenient way
		log.Fatalf("Site is down: %v\n", err)
		// output: 2006/01/02 15:04:05 Site is down: no such domain: bad.gopl.io
		// NOTE: Fourth, log the error and then continue
		if err := Ping(); err != nil {
			log.Printf("ping failed: %v; networking disabled", err)
			// fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
		}
		// NOTE: Fifth, safely ignore an error entirely
		dir, err := ioutil.tempDir("", "scratch")
		if err != nil {
			return fmt.Errorf("failed to create temp dir: %v", err)
		}
		// ...use temp dir...
		os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically
	}
}

// NOTE: Second, retry the failed operation with a limit on the
// number of attempts or the time spent trying before giving up
// entirely.
// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForserver(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

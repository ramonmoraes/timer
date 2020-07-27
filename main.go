package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ramonmoraes/timer/cmd"
)

//go:generate go run ./audio/generator.go

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 2 {
		fmt.Println(`Timer should receive a time arguments, may it be 2s or 2 s,
		e.g:
		3s or 3 s-> 3 seconds
		2m or 2 m-> 2 minutes `)
		return
	}

	waitTime := cmd.ParseTime(strings.Join(args, ""))
	shouldPlay := handleInterruptions(waitTime)

	if shouldPlay {
		cmd.PlayBeep(2 * time.Second)
	}
}

func handleInterruptions(waitTime time.Duration) bool {
	started := time.Now()

	done := make(chan bool, 1)
	// Wait default time expected time
	go func() {
		time.Sleep(waitTime)
		done <- true
	}()

	// Custom handler for signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP)
	go func() {
		sig := <-sigs

		executionTime := time.Now().Sub(started)
		minutes := executionTime.Minutes()
		seconds := executionTime.Seconds()
		fmt.Printf("\n[%s] Timer runned for: %.0fm %.0fs\n", sig, minutes, seconds)
		done <- false
	}()

	return <-done
}

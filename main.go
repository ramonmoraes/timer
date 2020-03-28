package main

import (
	"fmt"
	"os"
	"strings"
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
	time.Sleep(waitTime)
	cmd.PlayBeep(2 * time.Second)
}

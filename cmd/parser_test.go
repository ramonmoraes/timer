package cmd

import (
	"fmt"
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	input := `2s`
	parsed := ParseTime(input)
	var err error

	if parsed != 2*time.Second {
		err = fmt.Errorf("2s should be parsed to two seconds")
	}

	input = `2 m  `
	parsed = ParseTime(input)

	if parsed != 2*time.Minute {
		err = fmt.Errorf("'2 m  ' should be parsed to two minutes, but got parsed in %v", parsed)
	}

	if err != nil {
		t.Error(err)
	}
}

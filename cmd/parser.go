package cmd

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var UnitMap = map[string]time.Duration{
	"s": time.Second,
	"m": time.Minute,
}

func ParseTime(input string) time.Duration {
	input = strings.TrimSpace(input)
	reg := regexp.MustCompile(`(?P<amount>\d+) ?(?P<unit>\w{1})`)

	results := map[string]string{}
	for i, name := range reg.FindStringSubmatch(input) {
		results[reg.SubexpNames()[i]] = name
	}

	var err error
	for _, key := range []string{"amount", "unit"} {
		if _, ok := results[key]; !ok {
			err = fmt.Errorf("input must have an %s", key)
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	amount, err := strconv.Atoi(results["amount"])
	if err != nil {
		log.Fatal(err)
	}

	unit := results["unit"]

	if _, ok := UnitMap[results["unit"]]; !ok {
		log.Fatal(fmt.Sprintf("Unit map does not have %s", unit))
	}

	unitConverted := UnitMap[unit]

	return time.Duration(amount) * unitConverted
}

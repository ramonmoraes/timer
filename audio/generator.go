// +build

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var template string = `package audio

	
var AudioBytes = %#v
`

func main() {
	bytes := getFileBytes()
	generated := fmt.Sprintf(template, bytes)
	err := ioutil.WriteFile("audio/audio.go", []byte(generated), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getFileBytes() []byte {
	file, err := os.OpenFile("audio/audio.mp3", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

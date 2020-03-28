// +build

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	bytes := getFileBytes()
	template := fmt.Sprintf(`package audio

	
var AudioBytes := []bytes{%v}
`, bytes)

	err := ioutil.WriteFile("audio.go", []byte(template), 0644)
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

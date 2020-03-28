package audio

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func LoadRaw() io.ReadCloser {
	f, err := os.Open("./audio/audio.mp3")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func LoadSerialized() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(AudioBytes))
}

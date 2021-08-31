package audio

import (
	"bytes"
	_ "embed"
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

//go:embed audio.mp3
var AudioBytes []byte

func LoadSerialized() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(AudioBytes))
}

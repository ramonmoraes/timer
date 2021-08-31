package audio

import (
	"bytes"
	_ "embed"
	"io"
	"io/ioutil"
)

//go:embed audio.mp3
var AudioBytes []byte

func LoadSerialized() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(AudioBytes))
}

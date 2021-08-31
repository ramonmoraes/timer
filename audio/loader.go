package audio

import (
	"bytes"
	_ "embed"
	"io"
)

//go:embed audio.mp3
var AudioBytes []byte

func LoadSerialized() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(AudioBytes))
}

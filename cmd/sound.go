package cmd

import (
	"io"
	"log"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/ramonmoraes/timer/audio"
)

func PlayBeep(soundDuration time.Duration) {
	f := audio.LoadSerialized()
	playBeep(soundDuration, f)
}

func playBeep(soundDuration time.Duration, rc io.ReadCloser) {
	streamer, format, err := mp3.Decode(rc)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	time.Sleep(soundDuration)
}

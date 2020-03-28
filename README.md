## Timer cli

## Installing

```
sudo apt-get install libasound2-dev
go get -u https://github.com/ramonmoraes/timer
```

## Using

Just run timer and the expected time

```
$ timer 4s
```

Seconds `s` and minutes are suported `m`

## Development

To embed the audio file in the builded executable, the audio was serialized in the `audio/audio.go` file

In order to generate it, run:
```
go generate
```

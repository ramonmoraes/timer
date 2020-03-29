## Timer cli

## Installing

```
go get -u github.com/ramonmoraes/timer
```

### Dependencies:

Linux:
```
sudo apt install libasound2-dev libsdl2-dev
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

package main

import (
	"flag"
	"github.com/harunyasar/commandify/control"
	"os"
	"strings"
	"fmt"
)

func execute(flag string) bool {
	commands := map[string]interface{}{
		"next":     control.Next,
		"previous": control.Previous,
		"pause":    control.Pause,
		"play":     control.Play,
		"current":  control.Current,
		"status":   control.Status,
		"open":     control.Open,
		"close":    control.Close,
	}

	if command := commands[flag]; command == nil {
		return false
	}

	commands[flag].(func())()

	return true
}

var help string = `Usage: commandify [options...]
Options:
  -next		Skip to the next track.
  -previous	Skip to the previous track.
  -pause	Pause playback.
  -play		Resume playback.
  -current	Get the name of the track with that album artist of the track.
  -status	Is Spotify stopped or running?
  -open		Activate Spotify application.
  -close	Deactivate Spotify application.
`
var sNextTrack, sPreviousTrack, sPause, sPlay, sCurrentTrack, sStatus, sOpen, sClose bool

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, help)
	}

	control.CheckExistence()

	flag.BoolVar(&sNextTrack, "next", false, "Skip to the next track.")
	flag.BoolVar(&sPreviousTrack, "previous", false, "Skip to the previous track.")
	flag.BoolVar(&sPause, "pause", false, "Pause playback.")
	flag.BoolVar(&sPlay, "play", false, "Resume playback.")
	flag.BoolVar(&sCurrentTrack, "current", false, "Get the name of the track with that album artist of the track.")
	flag.BoolVar(&sStatus, "status", false, "Is Spotify stopped or running?")
	flag.BoolVar(&sOpen, "open", false, "Activate Spotify application.")
	flag.BoolVar(&sClose, "close", false, "Deactivate Spotify application.")

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	flagValue := strings.Replace(os.Args[1], "-", "", -1)

	if res := execute(flagValue); !res {
		flag.Usage()
		os.Exit(1)
	}
}

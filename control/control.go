package control

import (
	"fmt"
	"github.com/everdev/mack"
	"math"
	"os"
	"strconv"
)

// Constants of the tell method which will run via AppleScript
const (
	APP           = "Spotify"
	SYSTEM_EVENTS = "Finder"
	EXISTENCE     = `exists application file ((path to applications folder as string) & "` + APP + `")`
	NEXT          = `next track`
	PREVIOUS      = `previous track`
	PAUSE         = `pause`
	PLAY          = `play`
	ARTIST        = `artist of current track as string`
	TITLE         = `name of current track as string`
	ALBUM         = `album of current track as string`
	DURATION      = `duration of current track as string`
	STATUS        = `it is running`
	OPEN          = `activate`
	CLOSE         = `quit`
	PLAYER_STATE  = `player state`
)

// Mapping the result of player state
var playerState = map[string]int{
	"stopped": -1,
	"playing": 1,
	"paused":  0,
}

// Check the Spotify app is exists or not
func CheckExistence() {
	res, err := mack.Tell(SYSTEM_EVENTS, EXISTENCE)

	if err != nil {
		exit(err)
	}

	state, err := strconv.ParseBool(res)

	if err != nil {
		exit(err)
	} else {
		if !state {
			fmt.Printf("%s was not found on your computer!\n", APP)
			os.Exit(1)
		}
	}
}

// Exit with existing error message
func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

// Wrapper of AppleScript tell method
func tell(command string) string {
	res, err := mack.Tell(APP, command)
	if err != nil {
		exit(err)
	}
	return res
}

// Skip to the next track
func Next() {
	tell(NEXT)
	Current()
}

// Skip to the previous track
func Previous() {
	tell(PREVIOUS)
}

// Pause playback
func Pause() {
	tell(PAUSE)
}

// Resume playback
func Play() {
	tell(PLAY)
}

// Open Spotify if it is not already running
func Open() {
	if state := checkStatus(); state {
		fmt.Printf("%s is already running...\n", APP)
	} else {
		tell(OPEN)
	}
}

// Close Spotify if it is running
func Close() {
	if state := checkStatus(); state {
		fmt.Printf("%s is closing...\n", APP)
		tell(CLOSE)
	} else {
		fmt.Printf("%s is not running...\n", APP)
	}
}

// Check the state of Spotify application is running or not
func checkStatus() bool {
	res := tell(STATUS)
	state, err := strconv.ParseBool(res)
	if err != nil {
		exit(err)
	}
	return state
}

// Exported Spotify's status checker function
func Status() {
	if state := checkStatus(); state {
		fmt.Printf("%s is running...\n", APP)
	} else {
		fmt.Printf("%s is stopped...\n", APP)
	}
}

// Get the name of the track with that artist and duration of the track
func Current() {
	if playerState := getPlayerState(); playerState != -1 {
		artist, song, duration := getArtistOfCurrentTrack(), getNameOfCurrentTrack(), getDurationOfCurrentTrack()
		length := durationFormat(duration)
		fmt.Printf("♫ Now playing: %s - %s (%s) ♫\n", artist, song, length)
	} else {
		fmt.Printf("%s is stopped...\n", APP)
	}
}

// The artist of the track
func getArtistOfCurrentTrack() string {
	return tell(ARTIST)
}

// The name of the track
func getNameOfCurrentTrack() string {
	return tell(TITLE)
}

// The album name of the track
func getAlbumOfCurrentTrack() string {
	return tell(ALBUM)
}

//  The length of the track in seconds
func getDurationOfCurrentTrack() string {
	return tell(DURATION)
}

// Check the Spotify player state
func getPlayerState() int {
	state := tell(PLAYER_STATE)
	res := playerState[state]
	return res
}

//  Format the length of the track
func durationFormat(length string) string {
	duration, _ := strconv.ParseFloat(length, 64)
	minutes := math.Floor(duration / 1000 / 60)
	seconds := math.Mod(duration/1000, 60)
	return fmt.Sprintf("%1.0f:%1.0f", minutes, seconds)
}

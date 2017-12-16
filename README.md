![Commandify](http://i.imgur.com/vHhEMLD.png "Commandify")

# commandify
Control the Spotify app on CLI

## Objective
The main objective of this repository is my first command line application with Golang.
This application controls the Spotify with AppleScript scripts.

## Features
* **Next():** Skip to the next track
* **Previous():** Skip to the previous track
* **Pause():** Pause playback
* **Play():** Resume playback
* **Current():** Get the name of the track with that artist and duration of the track
* **Status():** Check Spotify is running or not
* **Open():** Open Spotify if it is not already running
* **Close():** Close Spotify if it is running

## Usages
### Skip to the next track
```
commandify -next
```
### Skip to the previous track
```
commandify -previous
```
### Pause playback
```
commandify -pause
```
### Resume playback
```
commandify -play
```
### Get the name of the track with that artist and duration of the track
```
commandify -current
```
### Check Spotify is running or not
```
commandify -status
```
### Open Spotify if it is not already running
```
commandify -open
```
### Close Spotify if it is running
```
commandify -close
```
### For getting help
```
commandify -help
```

## TODOs
* Write tests

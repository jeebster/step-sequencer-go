# step-sequencer-go
Coding Challenge:

A step sequencer program, written in Go. This program was authored extensibly to handle various i/o.

## Packages

### Sequencer
The `sequencer` package includes functionality for rendering `project` data via playback controls:
* `play()`: triggers track input via a step `scheduler`
* `stop()`: stops the sequencer

### Scheduler
The `scheduler` package includes functionality for scheduling step notifications relative to tempo-specific interval. It depends on a reference to project to allow for real-time changes (like changing project tempo)

### Project
The `project` package includes functionality for instantiating the elements of a project:
* Name
* Tempo
* Steps (number of steps for the sequencer to render)
* Triggers (index of tracks relative to step)

### Track
The `track` package includes functionality for instantiating the elements of a track:
* Name
* Input (any type that implements `io.Reader`, e.g. string, portaudio stream, etc.)

## Testing
NOTE: invoke tests with all .go files relative to package, e.g.

```
go test track/*.go
go test project/*.go
go test scheduler/*.go
go test sequencer/*.go
```

## Demo Application
Build and run the binary
```
go build
./step-sequencer-go
```

or invoke the main function
```
go run main.go
```

## Notes
* The architecture allows for flexibility commonly found in step sequencers (rename titles, change input/output, tempo, etc.). Future iterations could include support for multiple time signatures, effects processing, etc.
* Sequencer output is buggy. I added log statements to assist in my attempt to fix the issue. The log output mimics the desired functionality of the sequencer output - so this serves as a temporary solution to the problem. I am new to Golang, but from what I can gather, the issue appears to be related to [io.Copy](https://golang.org/pkg/io/#Copy): copying aggregate track inputs [io.MultiReader](https://golang.org/pkg/io/#MultiReader) to output [os.Stdout](https://golang.org/pkg/os/#pkg-variables)
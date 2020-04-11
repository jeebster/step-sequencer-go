package project

import (
	"github.com/jeebster/step-sequencer-go/track"
)

// Project contains project-specific metadata, a map of tracks to be triggered at specific steps, and a reference to tracks
// e.g. project name ('song' title), tempo, etc.
type Project struct {
	Name     string
	Tempo    int
	Steps    int
	Tracks   []*track.Track
	Triggers map[int][]*track.Track // map a step (int) to a track reference
}

// NewProject instantiates a new Project and returns a reference
func NewProject(name string, tempo int, steps int, tracks []*track.Track, triggers map[int][]*track.Track) *Project {
	p := &Project{
		Name:     name,
		Tempo:    tempo,
		Steps:    steps,
		Tracks:   tracks,
		Triggers: triggers,
	}

	return p
}

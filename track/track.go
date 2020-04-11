package track

import (
	"io"
)

// Track contains track-specific metadata and output within a Project context
// e.g. track name, audio file reference
// Track should support multiple i/o (audio, text, etc.)
type Track struct {
	Name  string
	Input io.Reader
}

// NewTrack instantiates a new Track and returns a reference
func NewTrack(name string, input io.Reader) *Track {
	t := &Track{
		Name:  name,
		Input: input,
	}

	return t
}

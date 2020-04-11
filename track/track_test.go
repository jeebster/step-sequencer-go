package track

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTrack(t *testing.T) {
	trackName := "Kick Drum"
	trackInput := strings.NewReader("Kick")
	newTrack := NewTrack(trackName, trackInput)

	assert.Equal(t, trackName, newTrack.Name, "track name should be equal")
	assert.Equal(t, trackInput, newTrack.Input, "track input should be equal")
}

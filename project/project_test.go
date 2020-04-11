package project

import (
	"strings"
	"testing"

	"github.com/jeebster/step-sequencer-go/track"
	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	track1Name := "Kick Drum"
	track1Input := strings.NewReader("Kick")

	track2Name := "Hi-Hat"
	track2Input := strings.NewReader("Hi-Hat")

	track1 := track.NewTrack(track1Name, track1Input)
	track2 := track.NewTrack(track2Name, track2Input)

	projectName := "Beat 1"
	projectTempo := 120
	projectSteps := 4
	projectTracks := []*track.Track{track1, track2}
	projectTriggers := map[int][]*track.Track{
		1: []*track.Track{track1},
		2: []*track.Track{track2},
		3: []*track.Track{track1},
		4: []*track.Track{track2},
	}
	newProject := NewProject(projectName, projectTempo, projectSteps, projectTracks, projectTriggers)

	assert.Equal(t, projectName, newProject.Name, "project name should be equal")
	assert.Equal(t, projectTempo, newProject.Tempo, "project tempo should be equal")
	assert.Equal(t, projectSteps, newProject.Steps, "project steps should be equal")
	assert.Equal(t, projectTracks, newProject.Tracks, "project tracks should be equal")
	assert.Equal(t, projectTriggers, newProject.Triggers, "project triggers should be equal")
}

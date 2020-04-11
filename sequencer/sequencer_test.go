package sequencer

import (
	"os"
	"strings"
	"testing"

	"github.com/jeebster/step-sequencer-go/project"
	"github.com/jeebster/step-sequencer-go/sequencer"
	"github.com/jeebster/step-sequencer-go/track"
	"github.com/stretchr/testify/assert"
)

// TODO: Move scaffolding export to package for leveraging in other tests
func scaffoldProject() *project.Project {
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
	return project.NewProject(projectName, projectTempo, projectSteps, projectTracks, projectTriggers)
}

func TestNewSequencer(t *testing.T) {
	project := scaffoldProject()
	output := os.Stdout
	newSequencer := sequencer.NewSequencer(project, output)

	assert.Equal(t, project, newSequencer.Project, "Sequencer project should be equal")
	assert.Equal(t, output, newSequencer.Output, "Sequencer project should be equal")
}

// TODO: Start(), Stop() test: research async/parallel testing with channels

package scheduler

import (
	"strings"
	"testing"
	"time"

	"github.com/jeebster/step-sequencer-go/project"
	"github.com/jeebster/step-sequencer-go/scheduler"
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

func TestNewScheduler(t *testing.T) {
	project := scaffoldProject()
	newScheduler := scheduler.NewScheduler(project)

	assert.Equal(t, project, newScheduler.Project, "Scheduler project should be equal")
}

func TestDurationLengthPerStep(t *testing.T) {
	project := scaffoldProject()
	s := scheduler.NewScheduler(project)
	expectation := time.Duration(60000000/s.Project.Tempo) * time.Microsecond
	assert.Equal(t, expectation, s.DurationLengthPerStep())
}

// TODO: Start() test: research how to test channels

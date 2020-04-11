package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jeebster/step-sequencer-go/project"
	"github.com/jeebster/step-sequencer-go/sequencer"
	"github.com/jeebster/step-sequencer-go/track"
)

func main() {
	// track initialization
	kickTrackInput := strings.NewReader("Kick")
	kickTrack := track.NewTrack("Kick", kickTrackInput)

	snareTrackInput := strings.NewReader("Snare")
	snareTrack := track.NewTrack("Snare", snareTrackInput)

	hiHatTrackInput := strings.NewReader("Hi-Hat")
	hiHatTrack := track.NewTrack("Hi-Hat", hiHatTrackInput)

	// project initialization
	projectName := "Four On The Floor"
	projectTempo := 120
	projectSteps := 16
	projectTracks := []*track.Track{kickTrack, snareTrack, hiHatTrack}
	projectTriggers := map[int][]*track.Track{
		1:  []*track.Track{kickTrack},
		2:  []*track.Track{},
		3:  []*track.Track{hiHatTrack},
		4:  []*track.Track{},
		5:  []*track.Track{kickTrack, snareTrack},
		6:  []*track.Track{},
		7:  []*track.Track{hiHatTrack},
		8:  []*track.Track{},
		9:  []*track.Track{kickTrack},
		10: []*track.Track{},
		11: []*track.Track{hiHatTrack},
		12: []*track.Track{},
		13: []*track.Track{kickTrack, snareTrack},
		14: []*track.Track{},
		15: []*track.Track{hiHatTrack},
		16: []*track.Track{},
	}
	project := project.NewProject(projectName, projectTempo, projectSteps, projectTracks, projectTriggers)

	// sequencer initialization
	sequencerOutput := os.Stdout
	sequencer := sequencer.NewSequencer(project, sequencerOutput)

	// Rawk \m/
	playingMessage := fmt.Sprintf("Playing: %s at %d BPM", project.Name, project.Tempo)
	fmt.Println(playingMessage)

	err := sequencer.Play()
	if err != nil {
		log.Fatal("Fatal error raised by sequencer.Play()")
	}

	// loop
	for {
		time.Sleep(time.Second)
	}
}

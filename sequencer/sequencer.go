package sequencer

import (
	"errors"
	"io"
	"log"

	"github.com/jeebster/step-sequencer-go/project"
	"github.com/jeebster/step-sequencer-go/scheduler"
)

// Sequencer contains a reference to Project (this allows a user to be able to switch projects)
// a channel to communicate with the Scheduler (receive playback information),
// and an output to render playback information
type Sequencer struct {
	Project   *project.Project
	Scheduler *scheduler.Scheduler
	IsPlaying bool
	Output    io.Writer
}

// NewSequencer instantiates a new Sequencer and returns a reference
func NewSequencer(project *project.Project, output io.Writer) *Sequencer {
	sch := scheduler.NewScheduler(project)
	seq := &Sequencer{
		Project:   project,
		IsPlaying: false,
		Scheduler: sch,
		Output:    output,
	}

	return seq
}

// Play(): playback control to start the sequencer:
// start and subscribe to scheduler, trigger tracks per step, and sum and output track inputs
func (s *Sequencer) Play() error {
	if s.IsPlaying {
		return errors.New("Sequencer is already playing")
	}

	go func() {
		for {
			stepToTrigger := <-s.Scheduler.Notifications
			log.Println("[DEBUG] step to trigger: ", stepToTrigger)
			tracksToSum := s.Project.Triggers[stepToTrigger]
			trackInputsToSum := make([]io.Reader, len(tracksToSum))
			for i := range tracksToSum {
				trackInputsToSum[i] = tracksToSum[i].Input
			}
			log.Println("[DEBUG] track inputs to sum: ", trackInputsToSum)

			summedTrackInputs := io.MultiReader(trackInputsToSum...)
			// NOTE: consider leveraging io.CopyBuffer if track input is audio
			if _, err := io.Copy(s.Output, summedTrackInputs); err != nil {
				log.Fatal("Fatal error attempting to copy aggregate track input to sequencer output")
			}
		}
	}()
	go s.Scheduler.Start()

	return nil
}

// Stop(): playback control to stop the sequencer
func (s *Sequencer) Stop() {
	// Stop/reset the schedule timer
	log.Println("Stopping the sequencer")
	close(s.Scheduler.Notifications)
	s.IsPlaying = false
}

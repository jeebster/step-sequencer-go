package scheduler

import (
	"time"

	"github.com/jeebster/step-sequencer-go/project"
)

// Scheduler contains a reference to Project (this allows a user to be able to switch project tempo without having to recall Start())
// and channel to communicate with the Sequencer (send playback information),
type Scheduler struct {
	Notifications chan int
	Project       *project.Project
}

// NewSequencer instantiates a new Sequencer and returns a reference
func NewScheduler(project *project.Project) *Scheduler {
	s := &Scheduler{
		Notifications: make(chan int),
		Project:       project,
	}
	return s
}

// Stop(): begin sending messages to receiver
// NOTE: The sequencer domain context is synchronous 1-1 (one receiver, one sender)
// this use case does not to conform to the golang channel closing principle
// so it is the receiver singleton's responsibility to close the channel
func (s *Scheduler) Start() {
	currentStep := 1
	for {
		if currentStep == s.Project.Steps {
			// we've reached the last step; start over from the beginning
			currentStep = 1
		}
		s.Notifications <- currentStep
		currentStep++
		sleepDuration := s.DurationLengthPerStep()
		time.Sleep(sleepDuration)
	}
}

// DurationLengthPerStep(): calculate distance between steps relative to tempo
func (s *Scheduler) DurationLengthPerStep() time.Duration {
	// not sure about precision of human ear/hearing so going with microseconds
	microSecondsPerMinute := 60000000
	microSecondsPerStep := time.Duration(microSecondsPerMinute/s.Project.Tempo) * time.Microsecond
	return time.Duration(microSecondsPerStep)
}

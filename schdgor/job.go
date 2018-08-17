package schdgor

import "time"

// ticker should be out of the job, another struct?
// or may be job should be an interface?
type Job struct {
	Name string
	// At      time.Time
	// chan as status?
	// ticker  time.Ticker
	status  int8
	Delay   time.Duration
	Period  time.Duration
	Handler func()
	// work chan Duration?
	// pause chan
	stop chan struct{}
	// Handler func(context.Context) error
}

// func (j *Job) SetDelay(dl time.Duration)  {
// 	delay = dl
// }

func (j *Job) Status() string {
	switch j.status {
	case 0:
		return "Ready"
	case 1:
		return "Running"
	default:
		return "Stopped"
	}
}

// schedule

// func (ctx context.Context) error {
//
// }

type JobWrap struct {
	Delay  time.Duration
	Period time.Duration

	Job Job
}

// https://stackoverflow.com/questions/3073948/job-task-and-process-whats-the-difference#3073961

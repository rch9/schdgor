package schdgor

import "time"

const (
	StatReady   = "Ready"
	StatStopped = "Stopped"
	StatRunning = "Running"
)

// ticker should be out of the job, another struct?
// or may be job should be an interface?
type Job struct {
	name string
	// At      time.Time
	// chan as status?
	// ticker  time.Ticker
	status  string
	delay   time.Duration
	period  time.Duration
	handler func()
	// work chan Duration?
	// pause chan
	stop chan struct{}
	// Handler func(context.Context) error
}

func (j *Job) Name() string {
	return j.name
}

func NewJob(name string, handler func(), delay, period time.Duration) *Job {
	j := Job{
		name:    name,
		handler: handler,
		delay:   delay,
		period:  period,
		stop:    make(chan struct{}, 1),
	}

	return &j
}

// func (j *Job) SetDelay(dl time.Duration)  {
// 	delay = dl
// }

// func (j *Job) Status() string {
// 	switch j.status {
// 	case 0:
// 		return "Ready"
// 	case 1:
// 		return "Running"
// 	default:
// 		return "Stopped"
// 	}
// }

// schedule

// func (ctx context.Context) error {
//
// }

// type JobWrap struct {
// 	Delay  time.Duration
// 	Period time.Duration
//
// 	Job Job
// }

// https://stackoverflow.com/questions/3073948/job-task-and-process-whats-the-difference#3073961

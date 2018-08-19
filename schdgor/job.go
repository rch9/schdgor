package schdgor

import (
	"context"
	"log"
	"time"
)

// available states of any job
const (
	StatReady    = "Ready"
	StatStopped  = "Stopped"
	StatRunning  = "Running"
	StatModified = "Modified"
)

// jobConf represents timeconfig of a job
type jobConf struct {
	Delay  time.Duration
	Period time.Duration
	// TODO: Timeout  time.Duration
	// TODO: WorkTime time.Duration
}

// job represents parameters of running gorutine
type job struct {
	name    string
	status  string
	conf    jobConf
	handler func(context.Context) error
	cancel  context.CancelFunc
	// done    chan struct{}
}

// Name returns job name
func (j *job) Name() string {
	return j.name
}

// Conf returns time config of the job
func (j *job) Conf() jobConf {
	return j.conf
}

// SetConf sets timeconfig to the job
func (j *job) SetConf(delay, period time.Duration) {
	j.conf.Delay = delay
	j.conf.Period = period
}

// NewJob creates new job with parameters
func NewJob(name string, handler func(context.Context) error, delay, period time.Duration) *job {
	// TODO: check why one of this works with t <= 0
	if delay < 0 {
		log.Fatal("delay < 0")
	}
	if period <= 0 {
		// TODO: check period == 0
		log.Fatal("period < 0")
	}

	j := job{
		name:    string(name),
		handler: handler,
		conf:    jobConf{delay, period},
	}

	return &j
}

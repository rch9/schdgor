package schdgor

import (
	"context"
	"log"
	"time"
)

// available states of any job
const (
	StatReady   = "Ready"
	StatStopped = "Stopped"
	StatRunning = "Running"
)

// jobConf represents timeconfig of a job
type jobConf struct {
	Delay  time.Duration
	Period time.Duration
	// TODO: Timeout  time.Duration
	// TODO: WorkTime time.Duration
}

type JobNameKey string

// job represents parameters of running gorutine
type job struct {
	name   JobNameKey
	status string
	// handler func()
	stop    chan struct{}
	conf    jobConf
	handler func(context.Context) error
}

// Name returns job name
func (j *job) Name() JobNameKey {
	return j.name
}

func (jk JobNameKey) String() string {
	return string(jk)
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
	if delay < 0 {
		log.Fatal("delay < 0")
	}
	if period <= 0 {
		// TODO: check period == 0
		log.Fatal("period < 0")
	}

	j := job{
		name:    JobNameKey(name),
		handler: handler,
		conf:    jobConf{delay, period},
		stop:    make(chan struct{}, 1),
	}

	return &j
}

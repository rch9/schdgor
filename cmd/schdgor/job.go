package main

import "time"

// ticker should be out of the job, another struct?
// or may be job should be an interface?
type Job struct {
	Name string
	// At      time.Time
	// chan as status?
	ticker  time.Ticker
	Delay   time.Duration
	Period  time.Duration
	Handler func()
	done    chan int
	// Handler func(context.Context) error
}

// https://stackoverflow.com/questions/3073948/job-task-and-process-whats-the-difference#3073961

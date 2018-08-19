package schdgor

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// // Test_Scheduler_Add tests adding jobs into scheduler
func Test_Scheduler_AddJobs(t *testing.T) {
	// job names
	name1, name2 := "Job-1", "Job-2"

	// creating job 1, 2
	j1 := NewJob(name1, func(context.Context) error {
		fmt.Println("I am job-1")
		return nil
	}, 0, time.Millisecond)

	j2 := NewJob(name2, func(context.Context) error {
		fmt.Println("I am job-2")
		return nil
	}, 0, time.Millisecond)

	// creating scheduler
	sc := New()

	// adding jobs to
	sc.AddJobs(j1, j2)

	//getting jobs
	jobs := sc.JobsPool()

	// testing jobs len
	assert(t, "len", len(jobs), 2)

	// testing jobs names
	assert(t, "names", jobs[name1].name, name1)
	assert(t, "names", jobs[name2].name, name2)

	// testing jobs statuses
	assert(t, "statuses", jobs[name1].status, StatReady)
	assert(t, "statuses", jobs[name2].status, StatReady)
}

// Test_Scheduler_WithStatus tests method which filters jobsPool by job statuses
func Test_Scheduler_WithStatus(t *testing.T) {
	// job names
	name1, name2 := "Job-1", "Job-2"

	// creating job 1, 2
	j1 := NewJob(name1, func(context.Context) error {
		fmt.Println("I am job-1")
		return nil
	}, 0, time.Millisecond)

	j2 := NewJob(name2, func(context.Context) error {
		fmt.Println("I am job-2")
		return nil
	}, 0, time.Millisecond)

	// creating scheduler
	sc := New()

	// adding jobs to
	sc.AddJobs(j1, j2)

	readyjobs := sc.JobsPool().WithStatus(StatReady)
	assert(t, "ready jobs len", len(readyjobs), 2)

	// creating new context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// start job 1
	sc.StartJob(ctx, cancel, j1.name)

	// getting running jobs
	runningjobs := sc.JobsPool().WithStatus(StatRunning)
	assert(t, "running jobs len after starting 0", len(runningjobs), 1)
	assert(t, "running jobs names after starting 0", runningjobs[j1.name].status, StatRunning)

	// getting jobs
	jobs := sc.JobsPool()

	// testing jobs len
	assert(t, "len", len(jobs), 2)

	// stopping job-1
	sc.StopJob(name1)

	// getting stopped jobs
	stoppedjobs := sc.JobsPool().WithStatus(StatStopped)
	fmt.Println(sc.JobsPool()["Job-1"])
	fmt.Println(sc.JobsPool()["Job-2"])
	assert(t, "stopped jobs len after stopping 2", len(stoppedjobs), 1)
	assert(t, "stopped jobs names after stopping 2", stoppedjobs[j1.name].status, StatStopped)
}

// Test_Scheduler_Start_Stop tests StartJobs and StopJob scheduler methods
func Test_Scheduler_StartJobs_StopJob(t *testing.T) {
	// job names
	name1, name2 := "Job-1", "Job-2"

	// creating job 1, 2
	j1 := NewJob(name1, func(context.Context) error {
		fmt.Println("I am job-1")
		return nil
	}, time.Hour, time.Millisecond)

	j2 := NewJob(name2, func(context.Context) error {
		fmt.Println("I am job-2")
		return nil
	}, time.Nanosecond, time.Nanosecond)

	// creating scheduler
	sc := New()

	// adding jobs to
	sc.AddJobs(j1, j2)

	// start jobs, cancel func will be created automatically
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	err := sc.StartJobs(ctx, cancel, j1.Name(), j2.Name())
	if err != nil {
		t.Errorf("error with start jobs: %v", err)
	}

	time.Sleep(time.Millisecond * 10)

	// getting running jobs
	runningjobs := sc.JobsPool().WithStatus(StatRunning)
	assert(t, "running jobs len", len(runningjobs), 2)

	// stopping j1, j2 will be stopped automatically because j1 and j2 has
	// same CancelFuncs
	sc.StopJob(j1.Name())

	// time.Sleep(time.Second)

	// getting stopped jobs
	stoppedjobs := sc.JobsPool().WithStatus(StatStopped)
	assert(t, "stopped jobs len", len(stoppedjobs), 2)
}

func assert(t *testing.T, what string, got, exp interface{}) {
	if got != exp {
		t.Errorf("Error in asserting %s, got: %v, exp: %v", what, got, exp)
	}
}

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

	// adding same jobs to pool cause error, Second j2 does not added
	err := sc.AddJobs(j1, j1)
	assertNotNil(t, "add same jobs", err)

	// add j2
	err = sc.AddJobs(j2)
	assert(t, "add jobs", err, nil)

	readyjobs := sc.JobsPool().WithStatus(StatReady)
	assert(t, "ready jobs len", len(readyjobs), 2)

	// creating new context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// start job 1
	err = sc.StartJob(ctx, cancel, j1.name)
	assert(t, "starting job-1", err, nil)

	// getting running jobs
	runningjobs := sc.JobsPool().WithStatus(StatRunning)
	assert(t, "running jobs len after starting 0", len(runningjobs), 1)
	assert(t, "running jobs names after starting 0", runningjobs[j1.name].status, StatRunning)

	// getting jobs
	jobs := sc.JobsPool()

	// testing jobs len
	assert(t, "len", len(jobs), 2)

	// stopping job-1
	err = sc.StopJob(name1)
	assert(t, "stopping job-1", err, nil)

	// getting stopped jobs
	stoppedjobs := sc.JobsPool().WithStatus(StatStopped)
	fmt.Println(sc.JobsPool()["Job-1"])
	fmt.Println(sc.JobsPool()["Job-2"])
	assert(t, "stopped jobs len after stopping 2", len(stoppedjobs), 1)
	assert(t, "stopped jobs names after stopping 2", stoppedjobs[j1.name].status, StatStopped)
}

// Test_Scheduler_Start_Stop tests StartJobs and StopJob scheduler methods
func Test_Scheduler_StartJobs_StopJob_RemoveAllJobs(t *testing.T) {
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
	err := sc.AddJobs(j1, j2)
	assert(t, "add jobs", err, nil)

	// start jobs, common cancel func will be created automatically
	ctx := context.Background()
	err = sc.StartJobs(ctx, nil, j1.Name(), j2.Name())
	if err != nil {
		t.Errorf("error with start jobs: %v", err)
	}

	time.Sleep(time.Millisecond * 10)

	// getting running jobs
	runningjobs := sc.JobsPool().WithStatus(StatRunning)
	assert(t, "running jobs len", len(runningjobs), 2)

	// stopping j1, j2 will be stopped automatically because j1 and j2 has
	// same CancelFuncs
	err = sc.StopJob(j1.Name())
	assert(t, "stop job1", err, nil)

	// getting stopped jobs
	stoppedjobs := sc.JobsPool().WithStatus(StatStopped)
	assert(t, "stopped jobs len", len(stoppedjobs), 2)

	// removing all jobs from pool
	err = sc.RemoveAllJobs()
	assert(t, "remove all jobs", err, nil)

	// checking empty pool of jobs
	emptypool := sc.JobsPool()
	assert(t, "empty jobspool len", len(emptypool), 0)
}

func Test_Scheduler_StartJob_ModifyJobConf_RemoveJob(t *testing.T) {
	// job names
	name1 := "Job-1"
	olddelay, oldperiod := time.Nanosecond, time.Nanosecond
	newdelay, newperiod := time.Millisecond, time.Millisecond

	// creating job 1
	j1 := NewJob(name1, func(context.Context) error {
		fmt.Println("I am job-1")
		return nil
	}, olddelay, oldperiod)

	// creating scheduler
	sc := New()

	// adding jobs to scheduler
	err := sc.AddJobs(j1)
	assert(t, "add jobs", err, nil)

	// start jobs, cancel func will be created automatically
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	err = sc.StartJob(ctx, cancel, j1.name)
	assert(t, "StartJob", err, nil)

	// modifying running job must cause error
	err = sc.ModifyJobConf(j1.name, newdelay, newperiod)
	assertNotNil(t, "modify job config", err)

	// wait 10 Milliseconds
	time.Sleep(time.Millisecond * 10)

	// stopping j1
	err = sc.StopJob(j1.Name())
	assert(t, "stop job 1", err, nil)

	// modifying stopped job must not cause error
	err = sc.ModifyJobConf(j1.name, newdelay, newperiod)
	assert(t, "modify job config", err, nil)

	// getting j1 config
	modifjobs := sc.JobsPool().WithStatus(StatModified)
	assert(t, "modifjobs jobs len", len(modifjobs), 1)

	// getting job-1 from pool and checking its status
	j1 = modifjobs[j1.name]
	assert(t, "modified data delay", j1.conf.Delay, newdelay)
	assert(t, "modified data period", j1.conf.Period, newperiod)

	// start j1, wait it and stop
	ctx = context.Background()
	err = sc.StartJob(ctx, nil, j1.name)
	assert(t, "start job 1", err, nil)
	time.Sleep(time.Millisecond * 10)
	err = sc.StopAllJobs()
	assert(t, "stop all jobs", err, nil)

	// getting stopped jobs from pool and checking amount
	stopped := sc.JobsPool()
	assert(t, "empty stopped len", len(stopped), 1)

	// removing all jobs from pool
	err = sc.RemoveJob(j1.name)
	assert(t, "remove job", err, nil)

	// checking empty pool of jobs
	emptypool := sc.JobsPool()
	assert(t, "empty jobspool len", len(emptypool), 0)
}

func assert(t *testing.T, what string, got, exp interface{}) {
	if got != exp {
		t.Errorf("Error in asserting %s, got: %v, exp: %v", what, got, exp)
	}
}

func assertNotNil(t *testing.T, what string, got interface{}) {
	if got == nil {
		t.Errorf("Error: expected not nil, got nil")
	}
}

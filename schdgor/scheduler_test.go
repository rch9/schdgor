package schdgor

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// // Test_Scheduler_Add tests adding jobs into scheduler
// func Test_Scheduler_Add(t *testing.T) {
// 	// job names
// 	name1, name2 := "Job-1", "Job-2"
//
// 	// creating job 1, 2
// 	j1 := NewJob(name1, func() {
// 		fmt.Println("I am job-1")
// 	}, 0, time.Millisecond)
//
// 	j2 := NewJob(name2, func() {
// 		fmt.Println("I am job-2")
// 	}, 0, time.Millisecond)
//
// 	// creating scheduler
// 	sc := New()
//
// 	// adding jobs to
// 	sc.AddJobs(j1, j2)
//
// 	//getting jobs
// 	jobs := sc.JobsPool()
//
// 	// testing jobs len
// 	assert(t, "len", len(jobs), 2)
//
// 	// testing jobs names
// 	assert(t, "names", jobs[name1].name, name1)
// 	assert(t, "names", jobs[name2].name, name2)
//
// 	// testing jobs statuses
// 	assert(t, "statuses", jobs[name1].status, StatReady)
// 	assert(t, "statuses", jobs[name2].status, StatReady)
// }

// // Test_Scheduler_WithStatus tests method which filters jobsPool by job statuses
// func Test_Scheduler_WithStatus(t *testing.T) {
// 	// job names
// 	name1, name2 := "Job-1", "Job-2"
//
// 	// creating job 1, 2
// 	j1 := NewJob(name1, func() {
// 		fmt.Println("I am job-1")
// 	}, 0, time.Millisecond)
//
// 	j2 := NewJob(name2, func() {
// 		fmt.Println("I am job-2")
// 	}, 0, time.Millisecond)
//
// 	// creating scheduler
// 	sc := New()
//
// 	// adding jobs to
// 	sc.AddJobs(j1, j2)
//
// 	readyjobs := sc.JobsPool().WithStatus(StatReady)
// 	assert(t, "ready jobs len", len(readyjobs), 2)
//
// 	// start job 1
// 	sc.StartJob(j1.name)
//
// 	// getting running jobs
// 	runningjobs := sc.JobsPool().WithStatus(StatRunning)
// 	assert(t, "running jobs len after starting 1", len(runningjobs), 1)
// 	assert(t, "running jobs names after starting 1", runningjobs[j1.name].status, StatRunning)
//
// 	// getting jobs
// 	jobs := sc.JobsPool()
//
// 	// testing jobs len
// 	assert(t, "len", len(jobs), 2)
//
// 	// stopping job-1
// 	sc.StopJob(name1)
//
// 	// getting stopped jobs
// 	stoppedjobs := sc.JobsPool().WithStatus(StatStopped)
// 	assert(t, "stopped jobs len after stopping 1", len(stoppedjobs), 1)
// 	assert(t, "stopped jobs names after stopping 1", stoppedjobs[j1.name].status, StatStopped)
// }

func Test_Scheduler_Start(t *testing.T) {
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

	ctx := context.Background()

	// start job 1
	sc.StartJob(ctx, j1.Name().String())

	// getting running jobs
	runningjobs := sc.JobsPool().WithStatus(StatRunning)
	assert(t, "running jobs len after starting 1", len(runningjobs), 1)
	assert(t, "running jobs names after starting 1", runningjobs[j1.name].status, StatRunning)

	// getting jobs
	jobs := sc.JobsPool()

	// testing jobs len
	assert(t, "len", len(jobs), 2)

	// stopping job-1
	sc.StopJob(ctx, name1)

	// getting stopped jobs
	stoppedjobs := sc.JobsPool().WithStatus(StatStopped)
	assert(t, "stopped jobs len after stopping 1", len(stoppedjobs), 1)
	assert(t, "stopped jobs names after stopping 1", stoppedjobs[j1.name].status, StatStopped)
}

// func Test_Scheduler_StartAll(t *testing.T) {
//
// }

// func Test_Scheduler_Stop(t *testing.T) {
// 	j1 := NewJob("Job-1", func() {
// 		fmt.Println("I am job-1")
// 	}, 0, time.Second)
//
// 	j2 := NewJob("Job-2", func() {
// 		fmt.Println("I am job-2")
// 	}, 0, time.Second)
//
// 	// j3 := new(Job)
// 	// j3.Name = "J3"
// 	fmt.Printf("%p\n", &j1)
//
// 	sc := New()
// 	sc.AddJobs(j1, j2)
//
// 	sc.StartJob(j1.Name())
// 	time.Sleep(time.Second * 1)
// 	fmt.Println("hi")
// 	time.Sleep(time.Second * 1)
// 	sc.StopJob("Job-1")
// 	time.Sleep(time.Second * 2)
// 	fmt.Println(sc.JobsPool())
// }
//
// func Test_Scheduler_StopAll(t *testing.T) {
//
// }
//
// func Test_Scheduler_Delete(t *testing.T) {
//
// }
//
// func Test_Scheduler_DeleteAll(t *testing.T) {
//
// }
//
// func Test_Scheduler_Modify(t *testing.T) {
//
// }
//
func assert(t *testing.T, what string, got, exp interface{}) {
	if got != exp {
		t.Errorf("Error in asserting %s, got: %v, exp: %v", what, got, exp)
	}
}

package schdgor

import (
	"fmt"
	"testing"
	"time"
)

// Test_Scheduler_Add tests adding jobs into scheduler
func Test_Scheduler_Add(t *testing.T) {
	// job names
	name1, name2 := "Job-1", "Job-2"

	// creating job 1, 2
	j1 := NewJob(name1, func() {
		fmt.Println("I am job-1")
	}, 0, time.Millisecond)

	j2 := NewJob(name2, func() {
		fmt.Println("I am job-2")
	}, 0, time.Millisecond)

	// creating scheduler
	sc := New()

	// adding jobs to
	sc.Add(j1, j2)

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

func Test_Scheduler_WithStatus(t *testing.T) {
	// job names
	name1, name2 := "Job-1", "Job-2"

	// creating job 1, 2
	j1 := NewJob(name1, func() {
		fmt.Println("I am job-1")
	}, 0, time.Millisecond)

	j2 := NewJob(name2, func() {
		fmt.Println("I am job-2")
	}, 0, time.Millisecond)

	// creating scheduler
	sc := New()

	// adding jobs to
	sc.Add(j1, j2)

	// start job 1
	sc.Start(j1.name)

	// getting jobs
	jobs := sc.JobsPool()

	// testing jobs len
	assert(t, "len", len(jobs), 2)

	// testing jobs names
	assert(t, "names", jobs[name1].name, name1)
	assert(t, "names", jobs[name2].name, name2)

	// testing jobs statuses
	assert(t, "statuses", jobs[name1].status, StatRunning)
	assert(t, "statuses", jobs[name2].status, StatReady)

	// stopping
	sc.Stop(name1)

	// testing jobs statuses
	assert(t, "statuses", jobs[name1].status, StatStopped)
	assert(t, "statuses", jobs[name2].status, StatReady)
}

func Test_Scheduler_Start(t *testing.T) {

}

func Test_Scheduler_StartAll(t *testing.T) {

}

func Test_Scheduler_Stop(t *testing.T) {
	j1 := NewJob("Job-1", func() {
		fmt.Println("I am job-1")
	}, 0, time.Second)

	j2 := NewJob("Job-2", func() {
		fmt.Println("I am job-2")
	}, 0, time.Second)

	// j3 := new(Job)
	// j3.Name = "J3"
	fmt.Printf("%p\n", &j1)

	sc := New()
	sc.Add(j1, j2)

	sc.Start(j1.Name())
	time.Sleep(time.Second * 1)
	fmt.Println("hi")
	time.Sleep(time.Second * 1)
	sc.Stop("Job-1")
	time.Sleep(time.Second * 2)
	fmt.Println(sc.JobsPool())
}

func Test_Scheduler_StopAll(t *testing.T) {

}

func Test_Scheduler_Delete(t *testing.T) {

}

func Test_Scheduler_DeleteAll(t *testing.T) {

}

func Test_Scheduler_Modify(t *testing.T) {

}

func assert(t *testing.T, what string, got, exp interface{}) {
	if got != exp {
		t.Errorf("Error in asserting %s, got: %v, exp: %v", what, got, exp)
	}
}

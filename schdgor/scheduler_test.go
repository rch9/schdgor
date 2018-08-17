package schdgor

import (
	"fmt"
	"testing"
	"time"
)

func Test_Scheduler_Add(t *testing.T) {
	// creating job 1, 2
	j1 := NewJob("Job-1", func() {
		fmt.Println("I am job-1")
	}, 0, time.Millisecond)

	j2 := NewJob("Job-2", func() {
		fmt.Println("I am job-2")
	}, 0, time.Millisecond)

	// creating scheduler
	sc := New()

	// adding jobs to
	sc.Add(j1, j2)

	//getting jobs
	jobs := sc.JobsPool()

	if len(jobs) != 2 {
		t.Errorf("Scheduler_Add Error")
	}
}

func Test_Scheduler_WithStatus(t *testing.T) {

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

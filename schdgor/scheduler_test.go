package schdgor

import (
	"fmt"
	"testing"
	"time"
)

func Test_Scheduler_Stop(t *testing.T) {
	j1 := Job{Name: "Job-1", Period: time.Second * 1, Handler: func() {
		fmt.Println("I am job-1")
	}}

	j2 := Job{Name: "Job-2", Handler: func() {
		fmt.Println("I am job-2")
	}}

	sc := New()
	sc.Add(j1, j2)

	sc.Start("Job-1")
	time.Sleep(time.Second * 4)
	fmt.Println(sc.JobsPool)
	sc.Stop("Job-1")
	fmt.Println(sc.JobsPool)

	time.Sleep(time.Second * 4)
}

func Test_Scheduler_StopAll(t *testing.T) {
	j1 := Job{Name: "Job-1", Period: time.Second * 1, Handler: func() {
		fmt.Println("I am job-1")
	}}

	j2 := Job{Name: "Job-2", Handler: func() {
		fmt.Println("I am job-2")
	}}

	sc := New()
	sc.Add(j1, j2)

	sc.StopAll()
}

func Test_Scheduler_Delete(t *testing.T) {
	j1 := Job{Name: "Job-1", Period: time.Second * 1, Handler: func() {
		fmt.Println("I am job-1")
	}}

	j2 := Job{Name: "Job-2", Handler: func() {
		fmt.Println("I am job-2")
	}}

	sc := New()
	sc.Add(j1, j2)

	sc.Remove(j1.Name)

	sc.StopAll()
}

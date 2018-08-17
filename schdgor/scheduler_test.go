package schdgor

import (
	"fmt"
	"testing"
	"time"
)

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
	time.Sleep(time.Second * 2)
	sc.Stop("Job-1")
	time.Sleep(time.Second * 2)
	fmt.Println(sc.JobsPool())
}

// func Test_Scheduler_StopAll(t *testing.T) {
// 	j1 := Job{Name: "Job-1", Period: time.Second * 1, Handler: func() {
// 		fmt.Println("I am job-1")
// 	}}
//
// 	j2 := Job{Name: "Job-2", Handler: func() {
// 		fmt.Println("I am job-2")
// 	}}
//
// 	sc := New()
// 	sc.Add(j1, j2)
//
// 	sc.StopAll()
// }
//
// func Test_Scheduler_Delete(t *testing.T) {
// 	j1 := Job{Name: "Job-1", Period: time.Second * 1, Handler: func() {
// 		fmt.Println("I am job-1")
// 	}}
//
// 	j2 := Job{Name: "Job-2", Handler: func() {
// 		fmt.Println("I am job-2")
// 	}}
//
// 	sc := New()
// 	sc.Add(j1, j2)
//
// 	sc.Remove(j1.Name)
//
// 	sc.StopAll()
// }

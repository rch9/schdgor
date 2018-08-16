package main

import (
	"fmt"
	"time"
)

func main() {
	j1 := Job{Name: "Job-1", Period: time.Second * 1, Handler: func() {
		fmt.Println("I am job-1")
	}}

	j2 := Job{Name: "Job-2", Handler: func() {
		fmt.Println("I am job-2")
	}}

	sc := New()
	sc.Add(j1, j2)

	fmt.Println(sc.JobsPool)

	sc.Start("Job-1")
	time.Sleep(time.Second * 3)
	sc.Stop("Job-1")

	time.Sleep(time.Second * 5)
}

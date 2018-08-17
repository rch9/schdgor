package main

import (
	"fmt"
	"time"

	"github.com/rch9/schdgor/schdgor"
)

func main() {
	j1 := schdgor.Job{Name: "Job-1", Period: time.Second * 1, Handler: func() {
		fmt.Println("I am job-1")
	}}

	j2 := schdgor.Job{Name: "Job-2", Handler: func() {
		fmt.Println("I am job-2")
	}}

	sc := schdgor.New()
	sc.Add(j1, j2)

	sc.Start("Job-1")
	time.Sleep(time.Second * 4)
	fmt.Println(sc.JobsPool)
	sc.Stop("Job-1")
	fmt.Println(sc.JobsPool)

	time.Sleep(time.Second * 4)
}

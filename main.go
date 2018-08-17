package main

import (
	"fmt"
	"time"

	"github.com/rch9/schdgor/schdgor"
)

func main() {

	j1 := schdgor.NewJob("Job-1", func() {
		fmt.Println("I am job-1")
	}, 0, time.Second*1)

	j2 := schdgor.NewJob("Job-2", func() {
		fmt.Println("I am job-2")
	}, 0, time.Second*2)

	sc := schdgor.New()
	sc.Add(j1, j2)
	//
	sc.Start("Job-1")
	time.Sleep(time.Second * 4)
	// fmt.Println(sc.JobsPool)
	// sc.Stop("Job-1")
	// fmt.Println(sc.JobsPool)
	//
	// time.Sleep(time.Second * 4)
}

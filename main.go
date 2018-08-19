package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rch9/schdgor/schdgor"
)

func main() {
	// creating context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// creating jobs
	j1 := schdgor.NewJob("Job-1", func(context.Context) error {
		fmt.Println("I am job-1")
		return nil
	}, 0, time.Second*1)

	j2 := schdgor.NewJob("Job-2", func(context.Context) error {
		fmt.Println("I am job-2")
		return nil
	}, 0, time.Second*2)

	// creting scheduler
	sc := schdgor.New()

	// adding jobs to scheduler
	sc.AddJobs(j1, j2)

	// starting job-1
	sc.StartJob(ctx, cancel, j1.Name())
	time.Sleep(time.Second * 3)

	// stopping job-1
	sc.StopJob(j1.Name())
	time.Sleep(time.Second * 3)

	// creating new context because last have already used
	ctx, cancel = context.WithCancel(context.Background())

	// starting again job-1
	err := sc.StartJob(ctx, cancel, j1.Name())
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 3)

	// stopping job-1
	err = sc.StopJob(j1.Name())
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 3)

	// remove job-1
	sc.RemoveJob(j1.Name())

	// remove job-2
	sc.RemoveJob(j2.Name())

	// waiting all jobs
	sc.WaitJobs()
}

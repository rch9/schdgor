package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rch9/schdgor/schdgor"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	j1 := schdgor.NewJob("Job-1", func(context.Context) error {
		fmt.Println("I am job-1")
		return nil
	}, 0, time.Second*1)

	j2 := schdgor.NewJob("Job-2", func(context.Context) error {
		fmt.Println("I am job-2")
		return nil
	}, 0, time.Second*2)

	sc := schdgor.New()
	sc.AddJobs(j1, j2)

	sc.StartJob(ctx, j1.Name().String())
	sc.StartJob(ctx, j2.Name().String())

	time.Sleep(time.Second * 5)
	cancel()
	// sc.StopJob(ctx, j1.Name().String())
	// sc.StopJob(ctx, j2.Name().String())
	time.Sleep(time.Second * 5)

	// sc.WaitJobs()

	// // fmt.Println(sc.JobsPool)
	// // sc.Stop("Job-1")
	// // fmt.Println(sc.JobsPool)
	// //
	// // time.Sleep(time.Second * 4)
}

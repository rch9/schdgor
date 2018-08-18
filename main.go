package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rch9/schdgor/schdgor"
)

func main() {

	// for {
	// 	select {
	// 	case <-stop:
	//
	// 	default:
	// 		time.Sleep(time.Millisecond * 100)
	// 	}
	// }

	ctx := context.Background()

	j1 := schdgor.NewJob("Job-1", func(context.Context) error {
		fmt.Println("I am job-1")
		return nil
	}, 0, time.Second*1)

	// j2 := schdgor.NewJob("Job-2", func() {
	// 	fmt.Println("I am job-2")
	// }, 0, time.Second*2)
	//
	// sc := schdgor.New()
	// sc.AddJobs(j1, j2)
	// //
	// sc.StartJob("Job-1")
	// time.Sleep(time.Second * 4)
	// // fmt.Println(sc.JobsPool)
	// // sc.Stop("Job-1")
	// // fmt.Println(sc.JobsPool)
	// //
	// // time.Sleep(time.Second * 4)
}

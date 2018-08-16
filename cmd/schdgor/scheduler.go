package main

import (
	"fmt"
	"time"
)

// sc.Start().WithDelay().Periodically()

type Scheduler struct {
	JobsPool map[string]Job
}

func New() *Scheduler {
	sc := new(Scheduler)
	sc.JobsPool = map[string]Job{}
	// fmt.Println(sc)
	return sc
}

func (sc *Scheduler) Add(jobs ...Job) {
	for _, j := range jobs {
		sc.JobsPool[j.Name] = j
	}
}

// check errors
func (sc *Scheduler) Start(jn string) {
	j := sc.JobsPool[jn]
	j.ticker = *time.NewTicker(j.Period) // time.Tick()
	go func() {
		// for t := range j.ticker.C {
		// 	fmt.Println(t)
		// 	j.Handler()
		// }

		for {
			select {
			case res := <-j.ticker.C:
				j.Handler()
				fmt.Println(res)
				// case <-done:
				// 	return
			}
		}

	}()
}

func (sc *Scheduler) StartAll() {

}

func (sc *Scheduler) Stop(jn string) {
	j := sc.JobsPool[jn]
	j.ticker.Stop()
}

func (sc *Scheduler) StopAll() {

}

func (sc *Scheduler) Pause(jn string) {

}

func (sc *Scheduler) PauseAll() {

}

func (sc *Scheduler) Kill(jn string) {

}

func (sc *Scheduler) KillAll() {

}

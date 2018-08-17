package schdgor

import (
	"fmt"
	"time"
)

// sc.Start().WithDelay().Periodically()

type Scheduler struct {
	JobsPool map[string]*Job
}

func New() *Scheduler {
	sc := new(Scheduler)
	sc.JobsPool = map[string]*Job{}
	// fmt.Println(sc)
	return sc
}

func (sc *Scheduler) Add(jobs ...Job) {
	for _, j := range jobs {
		sc.JobsPool[j.Name] = &j
	}
}

// TODO: check empty struct
// func start() chan bool {
//
// }

// check errors
func (sc *Scheduler) Start(jn string) {
	j := sc.JobsPool[jn]
	// j.ticker = *time.NewTicker(j.Period) // time.Tick()
	// FIXME: check memory
	tick := time.Tick(j.Period)
	j.stop = make(chan struct{}, 1)
	fmt.Println(j.stop)
	go func() {
		for {
			select {
			// TODO: why a can not close it?
			case res := <-tick:
				j.Handler()
				fmt.Println(res)
			case <-j.stop:
				fmt.Println("done")
				return
			}
		}

	}()
}

func (sc *Scheduler) StartAll() {

}

func (sc *Scheduler) Stop(jn string) {
	fmt.Println("called Stop")
	j := sc.JobsPool[jn]
	fmt.Println(j.stop)
	j.stop <- struct{}{}
	close(j.stop)
	// close(j.done)
	// j.ticker.Stop()
	// fmt.Println("called Stop")
}

func (sc *Scheduler) StopAll() {

}

func (sc *Scheduler) Pause(jn string) {

}

func (sc *Scheduler) PauseAll() {

}

func (sc *Scheduler) Remove(jn string) {

}

func (sc *Scheduler) RemoveAll() {

}

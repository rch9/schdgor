package schdgor

// TODO: fix job in runtime

import (
	"fmt"
	"log"
	"time"
)

// sc.Start().WithDelay().Periodically()

type JobsPool map[string]*Job

//
type Scheduler struct {
	JobsPool JobsPool
}

// New creates new Scheduler
func New() *Scheduler {
	sc := new(Scheduler)
	sc.JobsPool = JobsPool{}

	return sc
}

func (p JobsPool) WithStatus(s string) {
	for _, j := range p {
		if j.status != s {
			delete(p, j.Name)
		}
	}
}

func (sc *Scheduler) Add(jobs ...Job) {
	for _, j := range jobs {
		sc.JobsPool[j.Name] = &j
		(&j).status = StatReady
	}
}

// check errors
// Should i return error?
func (sc *Scheduler) Start(jn string) {
	j, ok := sc.JobsPool[jn]
	if !ok {
		log.Printf("Can not find job %s", jn)
		return
	}

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

// Should i return error?
func (sc *Scheduler) Stop(jn string) error {
	// fmt.Println("called Stop")
	j, ok := sc.JobsPool[jn]
	if !ok {
		return fmt.Errorf("error....")
	}

	fmt.Println(j.stop)
	j.stop <- struct{}{}
	close(j.stop)

	return nil
	// close(j.done)
	// j.ticker.Stop()
	// fmt.Println("called Stop")
}

func (sc *Scheduler) StopAll() {
	for _, j := range sc.JobsPool {
		fmt.Println(j.Name)
	}
}

func (sc *Scheduler) Remove(jn string) {
	// TODO: firstly stop
	delete(sc.JobsPool, jn)
}

func (sc *Scheduler) RemoveAll() {
	// for _, j := range sc.JobsPool {
	//
	// }
}

// func (sc *Scheduler) Pause(jn string) {
//
// }
//
// func (sc *Scheduler) PauseAll() {
//
// }

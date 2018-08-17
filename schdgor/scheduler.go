package schdgor

import (
	"fmt"
	"time"
)

// jobsPool is an alias for map[string]*Job
type jobsPool map[string]*Job

// Scheduler manages jobs which run in gorutines
// it stores jobs into jobsPool
type Scheduler struct {
	jobsPool jobsPool
}

// New creates new Scheduler
func New() *Scheduler {
	sc := new(Scheduler)
	sc.jobsPool = jobsPool{}

	return sc
}

// JobsPool returns copy of jobsPool
func (sc *Scheduler) JobsPool() jobsPool {
	return sc.jobsPool
}

// WithStatus filters jobs by their status (Ready/Running/Stopped)
func (p jobsPool) WithStatus(s string) jobsPool {
	var res []*Job
	for _, j := range p {
		if j.status == s {
			res = append(res, j)
		}
	}
	return p
}

// Add adds pointers of jobs into scheduler jobsPool
func (sc *Scheduler) Add(jobs ...*Job) {
	for _, j := range jobs {
		j.status = StatReady
		sc.jobsPool[j.name] = j
	}
}

// TODO: : Should i return error?
// Start runs specific job by its name
func (sc *Scheduler) Start(jn string) error {
	j, ok := sc.jobsPool[jn]
	if !ok {
		return fmt.Errorf("can not find job %s", jn)
	}

	j.status = StatRunning
	ticker := time.NewTicker(j.period)
	go func() {
		for {
			select {
			case <-ticker.C:
				j.handler()
			case <-j.stop:
				ticker.Stop()
				return
			}
		}
	}()

	return nil
}

// StartAll starts all jobs in jobsPool
func (sc *Scheduler) StartAll() {
	for _, j := range sc.jobsPool {
		// TODO: what if gorutine have already Running
		sc.Start(j.name)
	}
}

// TODO: what if gorutine have already Stopped
// Stop stops specific job by its name
func (sc *Scheduler) Stop(jn string) error {
	j, ok := sc.jobsPool[jn]
	if !ok {
		return fmt.Errorf("can not find job %s", jn)
	}

	j.stop <- struct{}{}
	j.status = StatStopped
	close(j.stop)
	return nil
}

// StartAll starts all jobs in jobsPool
func (sc *Scheduler) StopAll() {
	for _, j := range sc.jobsPool {
		fmt.Println(j.name)
	}
}

// //
// func (sc *Scheduler) Modify(jn) error {
// 	_, ok := sc.jobsPool[jn]
// 	if !ok {
// 		return fmt.Errorf("can not find job %s", jn)
// 	}
// }

// Remove removes specific job by its name
func (sc *Scheduler) Remove(jn string) error {
	_, ok := sc.jobsPool[jn]
	if !ok {
		return fmt.Errorf("can not find job %s", jn)
	}
	sc.Stop(jn)
	delete(sc.jobsPool, jn)
	return nil
}

// RemoveAll removes all jobs in jobsPool
func (sc *Scheduler) RemoveAll() {
	for _, j := range sc.jobsPool {
		sc.Remove(j.name)
	}
}

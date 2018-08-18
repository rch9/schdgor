package schdgor

// TODO: think about stopping scheduler with chan

import (
	"fmt"
	"log"
	"time"
)

// jobsPool is an alias for map[string]*Job
type jobsPool map[string]*job

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
func (p jobsPool) WithStatus(s string) jobsPool { //// TODO: check out tipe
	var res []*job
	for _, j := range p {
		if j.status == s {
			res = append(res, j)
		}
	}
	return p
}

// Add adds pointers of jobs into scheduler jobsPool
func (sc *Scheduler) Add(jobs ...*job) {
	for _, j := range jobs {
		j.status = StatReady
		sc.jobsPool[j.name] = j
	}
}

// TODO: check infinity working
// Start runs specific job by its name
func (sc *Scheduler) Start(jn string) error {
	j, ok := sc.jobsPool[jn]
	if !ok {
		return fmt.Errorf("can not find job %s", jn)
	}
	if j.status == StatRunning {
		return fmt.Errorf("job %s has already running", j.name)
	}

	j.status = StatRunning
	go func() {
		if j.conf.Delay > 0 {
			timer := time.NewTimer(j.Conf().Delay)
			select {
			case <-timer.C:
				j.handler()
				log.Println("timer")
			}
		}

		ticker := time.NewTicker(j.conf.Period)
		for {
			select {
			case <-ticker.C:
				j.handler()
				log.Println("ticker")
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
		if j.status != StatRunning {
			sc.Start(j.name)
		}
	}
}

// Stop stops specific job by its name
func (sc *Scheduler) Stop(jn string) error {
	j, ok := sc.jobsPool[jn]
	if !ok {
		return fmt.Errorf("can not find job %s", jn)
	}
	if j.status == StatStopped {
		return fmt.Errorf("job %s has already stopped", j.name)
	}

	j.stop <- struct{}{}
	j.status = StatStopped
	close(j.stop)
	return nil
}

// StartAll starts all jobs in jobsPool
func (sc *Scheduler) StopAll() {
	for _, j := range sc.jobsPool {
		if j.status == StatRunning {
			sc.Stop(j.name)
		}
	}
}

// ModifyJobConf modifies job time configuration
func (sc *Scheduler) ModifyJobConf(jn string, delay, period time.Duration) error {
	j, ok := sc.jobsPool[jn]
	if !ok {
		return fmt.Errorf("can not find job %s", jn)
	}
	if j.status == StatRunning {
		return fmt.Errorf("job %s has already running", j.name)
	}

	// TODO: check tests with stopped
	j.SetConf(delay, period)

	return nil
}

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

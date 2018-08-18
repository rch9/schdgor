package schdgor

// TODO: think about stopping scheduler with chan

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// jobsPool is an alias for map[string]*Job
type jobsPool map[jobNameKey]*job

// Scheduler manages jobs which run in gorutines
type Scheduler struct {
	jobsPool jobsPool
	wg       sync.WaitGroup
}

// New creates new Scheduler
func New() *Scheduler {
	sc := new(Scheduler)
	sc.jobsPool = jobsPool{}

	return sc
}

// WaitJobs force scheduler wait gorutines
func (sc *Scheduler) WaitJobs() {
	sc.wg.Wait()
}

// JobsPool returns copy of jobsPool
func (sc *Scheduler) JobsPool() jobsPool {
	return sc.jobsPool
}

// WithStatus filters jobs by their status (Ready/Running/...)
func (p jobsPool) WithStatus(s string) jobsPool {
	res := make(jobsPool)
	for _, j := range p {
		if j.status == s {
			res[j.name] = j
		}
	}
	return res
}

// addJob add one job to scheduler
func (sc *Scheduler) addJob(j *job) {
	j.status = StatReady
	sc.jobsPool[j.name] = j
	sc.wg.Add(1)
}

// AddJobs adds pointers of jobs into scheduler jobsPool
func (sc *Scheduler) AddJobs(jobs ...*job) {
	for _, j := range jobs {
		sc.addJob(j)
	}
}

// StartJob runs specific job by its name
func (sc *Scheduler) StartJob(ctx context.Context,
	cancel context.CancelFunc, jn string) error {

	if ctx == nil {
		return fmt.Errorf("context of job %s is nil", jn)
	}
	j, ok := sc.jobsPool[jobNameKey(jn)]
	if !ok {
		return fmt.Errorf("can not find job %s", jn)
	}
	if j.status == StatRunning {
		return fmt.Errorf("job %s has already running", j.name)
	}
	j.status = StatRunning

	// if user does not place CancelFunc
	// then new context creates with CancelFunc for closing job in future
	if cancel == nil {
		ctx, cancel = context.WithCancel(ctx)
	}
	j.cancel = cancel

	// stating job in new gorutine with context
	go startJob(ctx, j)

	return nil
}

// TODO: check timeout in context
// startJob starts new job and handles signals from channal
func startJob(ctx context.Context, j *job) {

	// waiting if delay > 0 or cancel
	if j.conf.Delay > 0 {
		timer := time.NewTimer(j.Conf().Delay)
		select {
		case <-timer.C:
			j.handler(ctx)
			log.Println("timer")
		case <-ctx.Done():
			log.Println("stopped")
			return
		}
	}

	// run periodic ticker
	ticker := time.NewTicker(j.conf.Period)
	for {
		select {
		case <-ticker.C:
			j.handler(ctx)
			log.Println("ticker")
		case <-ctx.Done():
			log.Println("stopped")
			ticker.Stop()
			return
		}
	}
}

// Stop stops specific job by its name
func (sc *Scheduler) StopJob(jn string) error {
	j, ok := sc.jobsPool[jobNameKey(jn)]
	if !ok {
		return fmt.Errorf("can not find job %s", jn)
	}
	if j.status == StatStopped {
		return fmt.Errorf("job %s has already stopped", j.name)
	}

	j.cancel()
	j.status = StatStopped
	return nil
}

// Remove removes specific job by its name
// job must not be running
func (sc *Scheduler) RemoveJob(jn string) error {
	j, ok := sc.jobsPool[jobNameKey(jn)]
	if !ok {
		return fmt.Errorf("can not find job %s", jn)
	}
	if j.status == StatRunning {
		return fmt.Errorf("job %s has already running", j.name)
	}
	defer sc.wg.Done()
	delete(sc.jobsPool, j.name)
	return nil
}

// TODO: does not work yet
// ModifyJobConf modifies job time configuration
// func (sc *Scheduler) ModifyJobConf(jn string, delay, period time.Duration) error {
// 	j, ok := sc.jobsPool[JobNameKey(jn)]
// 	if !ok {
// 		return fmt.Errorf("can not find job %s", jn)
// 	}
// 	if j.status == StatRunning {
// 		return fmt.Errorf("job %s has already running", j.name)
// 	}
//
// 	// TODO: check tests with stopped
// 	j.SetConf(delay, period)
//
// 	return nil
// }

// // RemoveAll removes all jobs in jobsPool
// func (sc *Scheduler) RemoveAllJobs() {
// 	for _, j := range sc.jobsPool {
// 		sc.RemoveJob(j.name)
// 	}
// }

// TODO: does not work yet
// // StartAll starts all jobs in jobsPool
// func (sc *Scheduler) StopAllJobs() {
// 	for _, j := range sc.jobsPool {
// 		if j.status == StatRunning {
// 			sc.StopJob(j.name)
// 		}
// 	}
// }

// // StartAll starts all jobs in jobsPool
// func (sc *Scheduler) StartAllJobs() {
// 	for _, j := range sc.jobsPool {
// 		if j.status != StatRunning {
// 			sc.StartJob(j.name)
// 		}
// 	}
// }

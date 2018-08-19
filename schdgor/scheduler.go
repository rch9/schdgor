package schdgor

// TODO: think about stopping scheduler with chan

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// jobsPool is an alias for map[string]*Job
type jobsPool map[string]*job

// Scheduler manages jobs which run in gorutines
type Scheduler struct {
	jobsPool jobsPool
	wg       sync.WaitGroup
	wgjobs   sync.WaitGroup
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
func (sc *Scheduler) addJob(j *job) error {
	_, ok := sc.jobsPool[j.name]
	if ok {
		return fmt.Errorf("job %s have already exists", j.name)
	}

	j.status = StatReady
	sc.jobsPool[j.name] = j
	sc.wg.Add(1)

	return nil
}

// AddJobs adds pointers of jobs into scheduler jobsPool
func (sc *Scheduler) AddJobs(jobs ...*job) {
	for _, j := range jobs {
		sc.addJob(j)
	}
}

// StartJobs runs specific jobs by their names with one context and cancel func
func (sc *Scheduler) StartJobs(ctx context.Context,
	cancel context.CancelFunc, jns ...string) error {

	if ctx == nil {
		return fmt.Errorf("context of jobs %v is nil", jns)
	}
	if cancel == nil {
		ctx, cancel = context.WithCancel(ctx)
	}

	for _, jn := range jns {
		err := sc.StartJob(ctx, cancel, jn)
		if err != nil {
			return fmt.Errorf("can not start job %s, %v", jn, err)
		}
	}
	return nil
}

// startJob runs specific job by its name
func (sc *Scheduler) StartJob(ctx context.Context,
	cancel context.CancelFunc, jn string) error {

	if ctx == nil {
		return fmt.Errorf("context of job %s is nil", jn)
	}
	j, ok := sc.jobsPool[string(jn)]
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
	sc.wgjobs.Add(1)
	go sc.startJob(ctx, j)

	return nil
}

// TODO: check timeout in context
// startJob starts new job and handles signals from channal
func (sc *Scheduler) startJob(ctx context.Context, j *job) {
	if j.conf.Delay > 0 {
		timer := time.NewTimer(j.Conf().Delay)
		select {
		// FIXME: Can i close timer?
		case <-timer.C:
			j.handler(ctx)
		case <-ctx.Done():
			j.stop()
			sc.wgjobs.Done()
			return
		}
	}

	// run periodic ticker
	ticker := time.NewTicker(j.conf.Period)
	for {
		select {
		case <-ticker.C:
			j.handler(ctx)
		case <-ctx.Done():
			ticker.Stop()
			j.stop()
			sc.wgjobs.Done()
			return
		}
	}
}

func (j *job) stop() {
	j.status = StatStopped
	// j.done <- struct{}{}
}

// Stop stops specific job by its name
// if jobs have same CancelFuncs then it all be closed
func (sc *Scheduler) StopJob(jn string) error {
	j, ok := sc.jobsPool[string(jn)]
	if !ok {
		return fmt.Errorf("can not find job %s", jn)
	}
	if j.status != StatRunning {
		return fmt.Errorf("job %s has already stopped", j.name)
	}

	j.cancel()
	sc.wgjobs.Wait()
	// waiting finishing job
	// <-j.done

	return nil
}

// Remove removes specific job by its name
// job must not be running
func (sc *Scheduler) RemoveJob(jn string) error {
	j, ok := sc.jobsPool[string(jn)]
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

// ModifyJobConf modifies job time configuration
func (sc *Scheduler) ModifyJobConf(jn string, delay, period time.Duration) error {
	j, ok := sc.jobsPool[string(jn)]
	if !ok {
		return fmt.Errorf("can not find job %s", jn)
	}
	if j.status == StatRunning {
		return fmt.Errorf("job %s has already running", j.name)
	}

	j.SetConf(delay, period)

	return nil
}

// // RemoveAll removes all jobs in jobsPool
func (sc *Scheduler) RemoveAllJobs() {
	for _, j := range sc.jobsPool {
		sc.RemoveJob(j.name)
	}
}

// StopAllJobs stops all running jobs in jobsPool
func (sc *Scheduler) StopAllJobs() {
	for _, j := range sc.jobsPool {
		if j.status == StatRunning {
			sc.StopJob(j.name)
		}
	}
}

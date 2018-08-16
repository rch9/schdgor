package main

import "context"

// ctx
// mux
// chan

type Job struct {
	Id     int
	Name   string
	Status int
	Period int
	Do     func(ctx context.Context) error
}

func main() {

}

func (j *Job) Run() {

}

func (j *Job) Pause() {

}

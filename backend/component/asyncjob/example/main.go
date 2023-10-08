package main

import (
	"context"
	"errors"
	"log"
	"time"
	"todo/component/asyncjob"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		// log.Println("I am job 1")
		time.Sleep(time.Second * 1)

		return errors.New("job 1 error")
	})

	job1.SetRetryDurations([]time.Duration{time.Millisecond * 500})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 3)
		log.Println("I am job 2")

		return nil
	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 2)
		log.Println("I am job 3")

		return nil
	})

	group := asyncjob.NewGroup(true, job1, job2, job3)

	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}
}

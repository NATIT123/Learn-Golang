package main

import (
	"context"
	"fmt"
	"log"
	"main/common/asyncjob"
	"time"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 2)
		fmt.Println("... I am job 1")
		return nil
	}, asyncjob.WithName("Job 1"))

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 3) // Simulate some work
		fmt.Println("... I am job 2")
		return nil
	}, asyncjob.WithName("Job 2"))

	if err := asyncjob.NewGroup(true, job1, job2).Run(context.Background()); err != nil {
		log.Println(err)
	}
}

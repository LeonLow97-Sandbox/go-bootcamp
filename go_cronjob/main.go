package main

import (
	"fmt"

	"github.com/robfig/cron"
)

func main() {
	// create a new cron job
	c := cron.New()

	// schedule the job to run every minute
	c.AddFunc("@every 2s", func() {
		fmt.Println("Hello from cron job")
	})

	// start the cron scheduler
	c.Start()

	// keep the program running
	select {}
}

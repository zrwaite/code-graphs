package cron

import (
	"time"

	"github.com/go-co-op/gocron"
)

func RunCronJobs() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Hours().Do(func() {
		parseCodeData()
	})
	s.StartBlocking()
}

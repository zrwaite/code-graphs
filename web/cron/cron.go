package cron

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/zrwaite/github-graphs/api/wakatime"
)

func RunCronJobs() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Hours().Do(func() {
		wakatime.ParseCodeData()
	})
	s.StartBlocking()
}

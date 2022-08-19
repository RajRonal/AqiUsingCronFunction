package utilities

import (
	"AqiWithCron/handler"
	"gopkg.in/robfig/cron.v2"
	"log"
)

func RunCronJobs() {
	cronRunner := cron.New()
	_, err := cronRunner.AddFunc("@every 60m", func() {
		handler.InsertAqi()
	})
	if err != nil {
		log.Printf("unable to add cron job")
		return
	}
	cronRunner.Start()
}

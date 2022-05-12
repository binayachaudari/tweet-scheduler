package main

import (
	"binayachaudari/schedule-tweets/models"
	route "binayachaudari/schedule-tweets/routes"
	event_scheduler "binayachaudari/schedule-tweets/utils"
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	db := models.ConnectDB()
	db.AutoMigrate(&models.ScheduleTweet{})

	ctx, cancel := context.WithCancel(context.Background())
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	scheduler := event_scheduler.NewScheduler(db)

	scheduler.CheckEventsInInterval(ctx, time.Minute/2)

	go func() {
		for range interrupt {
			log.Println("\n❌ Interrupt received closing...")
			cancel()
		}
	}()

	r := route.Routes(db)
	r.Run()

	<-ctx.Done()
}

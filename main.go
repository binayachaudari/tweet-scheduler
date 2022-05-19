package main

import (
	"binayachaudari/schedule-tweets/models"
	route "binayachaudari/schedule-tweets/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	// db := models.ConnectDB()
	// db.AutoMigrate(&models.ScheduleTweet{})

	// ctx, cancel := context.WithCancel(context.Background())
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)

	// scheduler := event_scheduler.NewScheduler(db)

	// scheduler.CheckEventsInInterval(ctx, time.Minute/2)

	app := fx.New(
		route.Module,
		models.Module,
		fx.Invoke(route.RegisterRoutes, startServer),
	)

	app.Run()
	// go func() {
	// 	for range interrupt {
	// 		log.Println("\n❌ Interrupt received closing...")
	// 		cancel()
	// 	}
	// }()

	// <-ctx.Done()
}

func startServer(eng *gin.Engine) {
	eng.Run()
}

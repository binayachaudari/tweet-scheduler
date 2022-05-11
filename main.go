package main

import (
	"binayachaudari/schedule-tweets/models"
	route "binayachaudari/schedule-tweets/routes"
)

func main() {
	db := models.ConnectDB()
	db.AutoMigrate(&models.ScheduleTweet{})

	r := route.Routes(db)
	r.Run()
}

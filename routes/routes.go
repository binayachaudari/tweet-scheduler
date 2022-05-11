package route

import (
	controller "binayachaudari/schedule-tweets/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(d *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("dbCon", d)
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/tweets", controller.GetAllTweets)
		v1.GET("/tweets/:id", controller.GetTweetDetail)
		v1.POST("/tweets", controller.ScheduleTweet)
		v1.PUT("/tweets/:id", controller.UpdateTweet)
		v1.DELETE("/tweets/:id", controller.DeleteTweet)
	}

	return router
}

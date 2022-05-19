package route

import (
	controller "binayachaudari/schedule-tweets/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(d *gorm.DB, eng *gin.Engine) *gin.Engine {
	v1 := eng.Group("/api/v1")
	{
		v1.GET("/tweets", controller.GetAllTweets)
		v1.GET("/tweets/:id", controller.GetTweetDetail)
		v1.POST("/tweets", controller.ScheduleTweet)
		v1.PUT("/tweets/:id", controller.UpdateTweet)
		v1.DELETE("/tweets/:id", controller.DeleteTweet)
	}

	return eng
}

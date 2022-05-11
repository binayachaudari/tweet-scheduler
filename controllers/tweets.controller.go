package controller

import (
	"binayachaudari/schedule-tweets/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Request body
type scheduleTweetInput struct {
	Tweet        string    `json:"tweet"`
	ScheduleTime time.Time `json:"schedule_time"`
}

// Get all tweets
func GetAllTweets(c *gin.Context) {
	db := c.MustGet("dbCon").(*gorm.DB)

	tweets := []models.ScheduleTweet{}

	db.Find(&tweets)
	c.IndentedJSON(http.StatusOK, tweets)
}

func GetTweetDetail(c *gin.Context) {
	db := c.MustGet("dbCon").(*gorm.DB)
	id := c.Param("id")

	tweet := models.ScheduleTweet{}

	if err := db.Where("id = ?", id).First(&tweet); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "No tweet recorded with id: " + id,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, tweet)
}

// Schedule Tweet
func ScheduleTweet(c *gin.Context) {
	input := scheduleTweetInput{}
	db := c.MustGet("dbCon").(*gorm.DB)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tweet := models.ScheduleTweet{Tweet: input.Tweet, ScheduleTime: input.ScheduleTime}
	db.Create(&tweet)

	c.IndentedJSON(http.StatusOK, tweet)
}

// Update tweet
func UpdateTweet(c *gin.Context) {
	db := c.MustGet("dbCon").(*gorm.DB)
	input := scheduleTweetInput{}
	tweet := models.ScheduleTweet{}
	id := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Where("id = ?", id).First(&tweet); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "No tweet recorded with id: " + id,
		})
		return
	}

	db.Model(&tweet).Updates(input)
	c.IndentedJSON(http.StatusOK, tweet)
}

// Delete tweet
func DeleteTweet(c *gin.Context) {

}

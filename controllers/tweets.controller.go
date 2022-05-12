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
	c.JSON(http.StatusOK, tweets)
}

func GetTweetDetail(c *gin.Context) {
	db := c.MustGet("dbCon").(*gorm.DB)
	id := c.Param("id")

	tweet := models.ScheduleTweet{}

	if err := db.First(&tweet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No tweet recorded with id: " + id,
		})
		return
	}

	c.JSON(http.StatusOK, tweet)
}

// Schedule Tweet
func ScheduleTweet(c *gin.Context) {
	input := scheduleTweetInput{}
	db := c.MustGet("dbCon").(*gorm.DB)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tweet := models.ScheduleTweet{Tweet: input.Tweet, ScheduleTime: input.ScheduleTime}
	db.Create(&tweet)

	c.JSON(http.StatusOK, tweet)
}

// Update tweet
func UpdateTweet(c *gin.Context) {
	db := c.MustGet("dbCon").(*gorm.DB)
	input := scheduleTweetInput{}
	tweet := models.ScheduleTweet{}
	id := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.First(&tweet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No tweet recorded with id: " + id,
		})
		return
	}

	db.Model(&tweet).Updates(models.ScheduleTweet{Tweet: input.Tweet, ScheduleTime: input.ScheduleTime})
	c.JSON(http.StatusOK, tweet)
}

// Delete tweet
func DeleteTweet(c *gin.Context) {
	db := c.MustGet("dbCon").(*gorm.DB)
	tweet := models.ScheduleTweet{}
	id := c.Param("id")

	if err := db.First(&tweet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No tweet recored with id: " + id,
		})
		return
	}

	db.Delete(&tweet)
	c.JSON(http.StatusOK, tweet)
}

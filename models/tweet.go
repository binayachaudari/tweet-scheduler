package models

import "time"

type ScheduleTweet struct {
	ID           int       `json:"id" uri:"id"`
	Tweet        string    `json:"tweet"`
	IsPublished  bool      `json:"is_published"`
	ScheduleTime time.Time `json:"scheduled_time"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

package event_scheduler

import (
	"binayachaudari/schedule-tweets/events"
	"binayachaudari/schedule-tweets/models"
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

// List of event listeners; has attached event listeners
type Listeners map[uint]ListenFunc

// ListenFunc is call when event is fired
type ListenFunc func(string)

// Event: that we schedule; structure defined as
type Event struct {
	ID    uint
	Tweet string
}

// Event scheduler data-structure
type Scheduler struct {
	database *gorm.DB
}

// NewScheduler creates a new scheduler
func NewScheduler(db *gorm.DB) Scheduler {
	return Scheduler{
		database: db,
	}
}

// Scheduler receiver function
func (s Scheduler) checkScheduledEvents() []Event {
	events := []Event{}
	tweets := []models.ScheduleTweet{}
	if err := s.database.Model(&tweets).Where("schedule_time < ?", time.Now()).Find(&events).Error; err != nil {
		log.Print("💀 error: ", err)
		return nil
	}

	return events
}

// Call Listener: call event-listener of provided event
func (s Scheduler) callListener(event Event) {
	tweet := models.ScheduleTweet{}

	go events.Tweet(event.Tweet)

	if err := s.database.First(&tweet, event.ID).Error; err != nil {
		log.Print("💀 error: ", err)
	}

	s.database.Delete(&tweet)
}

// CheckEventsInInterval checks the event in given interval
func (s Scheduler) CheckEventsInInterval(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(duration)

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return

			case <-ticker.C:
				log.Println("⏰ Ticks Received...")
				events := s.checkScheduledEvents()

				for _, e := range events {
					s.callListener(e)
				}
			}

		}
	}()
}

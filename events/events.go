package events

import (
	"binayachaudari/schedule-tweets/twitter"
	"encoding/json"
	"log"
)

func Tweet(t string, c chan bool) {
	twitterClient := twitter.SetupTwitter()

	tweet, _, err := twitterClient.Statuses.Update(t, nil)

	if err != nil {
		log.Fatal(err.Error())
		c <- false
	}

	log.Println("\n🐦 Tweeted", decodeTweet(tweet))
	c <- true
}

func decodeTweet(t any) string {
	decodedTweet, err := json.Marshal(t)

	if err != nil {
		log.Fatal("\n⚠️ Error Occured while marshalling JSON")
	}

	return string(decodedTweet)
}

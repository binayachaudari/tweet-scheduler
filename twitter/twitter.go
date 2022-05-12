package twitter

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func SetupTwitter() *twitter.Client {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerKeySecret := os.Getenv("CONSUMER_KEY_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerKeySecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	client := config.Client(oauth1.NoContext, token)

	twitterClient := twitter.NewClient(client)

	return twitterClient
}

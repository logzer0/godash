package main

import (
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/go-ini/ini"
)

var Config *ini.File

func main() {
	// Define buttons and action function
	DashMacs["74:c2:46:84:f5:ee"] = Tweet

	// Kick it off!
	SnifferStart()
}

func NoAction() {
	log.Println("No action on click.")
}

func Tweet() {
	PostTweet("Testing now")
}

func PostTweet(tweet string) {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)

	v := url.Values{}

	_, err := api.PostTweet(tweet, v)
	if err != nil {
		log.Println("We have an error now ", err)
	}

}

package main

import (
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/go-ini/ini"
)

var Config *ini.File

const (
	consumerKey    = "your key here"
	consumerSecret = "your secret here"
	accessToken    = "your token here"
	accessSecret   = "your secret here"
	macAddress     = "mac address of the dash button here"
)

func main() {
	// Define buttons and action function
	DashMacs[macAddress] = Tweet

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

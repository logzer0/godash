package main

import (
	"log"
	"math/rand"
	"net/url"
	"time"

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

var (
	statements = []string{
		"PHP is the coolest language",
		"I love Javascript",
		"I love working with Internet Explorer",
		"Functional programming is overrated",
		"LISP is for losers",
		"Windows phone is the next big thing",
	}
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
	rand.Seed(time.Now().UTC().UnixNano())
	PostTweet(statements[rand.Intn(len(statements))])
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

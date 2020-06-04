package miniTwiter

import (
	"time"
)

type MiniTwiterModel interface {
	PostTweet(userID int, tweetText string) error
	GetTimeLine(userID int) []string
	GetNewsFeed(userID int) []Tweet
	Follow(fromUserID, toUserID int) error
	UnFollow(fromUserID, toUserID int) error
}

type Tweet struct {
	Tweet  string
	timestamp time.Time
}

type miniTwiter struct {
	tweetMap map[int][]Tweet
	friends  map[int]map[int]bool
}

// Post a twweet
func (t *miniTwiter) PostTweet(userID int, tweetText string) error {
	tweet := Tweet{
		Tweet:  tweetText,
		timestamp: time.Now(),
	}
	// Todo: error handling for userID or tweetText not right
	tweets, ok := t.tweetMap[userID]
	if ok {
		t.tweetMap[userID] = append(t.tweetMap[userID], tweet)
	} else {
		t.tweetMap[userID] = []Tweet{tweet}
	}

	return nil
}

func (t *miniTwiter) GetTimeLine(userID int) []Tweet {
	tweets, ok := t.tweetMap[userID]
	if !ok {
		return nil
	}

	var newestTweets []Tweet
	if len(tweets) > 10 {
		newestTweets = tweets[(len(tweets) - 10):]
	}else{
		newestTweets = tweets
	}

	return newestTweets
}

func (t *miniTwiter) GetNewsFeed(userID int) []Tweet {
	var newsFeed []Tweet
	
	tweets, ok := t.tweetMap[userID]
	// get the newest 10 from user's tweets
	if ok{
		
		i := 0
		for i 
		
	}
	
	// get the newest 10 from users' friends 
	
	
	// sort by the timestamp
	
	
	return newsFeed
}

func (t *miniTwiter) Follow(fromUserID, toUserID int) error {
	return nil
}

func (t *miniTwiter) UnFollow(fromUserID, toUserID int) error {
	return nil
}

func NewMniniTwitter() MiniTwiterModel {

}

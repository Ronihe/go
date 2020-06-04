package miniTwitter

import (
	"fmt"
	"sort"
	"time"
	"errors"
)

type MiniTwiterModel interface {
	PostTweet(userID int, tweetText string) error
	GetTimeLine(userID int) []Tweet
	GetNewsFeed(userID int) []Tweet
	Follow(fromUserID, toUserID int) error
	UnFollow(fromUserID, toUserID int) error
}

type Tweet struct {
	Tweet     string
	timestamp time.Time
}

type miniTwiter struct {
	tweetMap map[int][]Tweet
	friends  map[int]map[int]bool
}

// Post a twweet
func (t *miniTwiter) PostTweet(userID int, tweetText string) error {
	tweet := Tweet{
		Tweet:     tweetText,
		timestamp: time.Now(),
	}
	// Todo: error handling for userID or tweetText not right
	_, ok := t.tweetMap[userID]
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
	} else {
		newestTweets = tweets
	}

	return newestTweets
}

func (t *miniTwiter) GetNewsFeed(userID int) []Tweet {
	var newsFeed []Tweet

	tweets, ok := t.tweetMap[userID]
	// get the newest 10 from user's tweets
	if ok {
		if len(tweets) < 10 {
			newsFeed = append(newsFeed, tweets...)
		} else {
			newsFeed = append(newsFeed, tweets[(len(tweets)-10):]...)
		}
	}

	// get the newest 10 from users' friends
	friends, ok := t.friends[userID]
	if ok {
		for friendID := range friends {
			friendTweets, ok := t.tweetMap[friendID]
			if ok {
				if len(friendTweets) < 10 {
					newsFeed = append(newsFeed, friendTweets...)
				} else {
					newsFeed = append(newsFeed, friendTweets[(len(tweets)-10):]...)
				}
			}
		}
	}

	sort.Slice(newsFeed, func(i, j int) bool { return newsFeed[i].timestamp.After(newsFeed[j].timestamp) })
	if len(newsFeed) > 10 {
		return newsFeed[:10]
	}
	return newsFeed
}

func (t *miniTwiter) Follow(fromUserID, toUserID int) error {
	
	_, ok := t.friends[fromUserID]
	if !ok{
		t.friends[fromUserID] = map[int]bool{}
	}
	t.friends[fromUserID][toUserID] = true
	fmt.Println(t.friends)
	return nil
}

func (t *miniTwiter) UnFollow(fromUserID, toUserID int) error {
	
	friends, ok := t.friends[fromUserID]
	
	if !ok{
		return errors.New("user1 is not following ")
	}
	
	delete(friends, toUserID)
	return nil
}

func NewMniniTwitter() MiniTwiterModel {
	return &miniTwiter{
		tweetMap: map[int][]Tweet{},
		friends: map[int]map[int]bool{},
	}
}

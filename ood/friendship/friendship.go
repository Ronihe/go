package friendship

import "sort"

type FriendshipInterface interface {
	GetFollowers(userID int) []int
	GetFollowings(userID int) []int
	Follow(toUserID, fromUserID int) error
	UnFollow(toUserID, fromUserID int) error
}

type Friendship struct {
	followers  map[int]map[int]bool
	followings map[int]map[int]bool
}

func (f *Friendship) GetFollowers(userID int) []int {
	followers, ok := f.followers[userID]
	if !ok {
		return nil
	}
	followersList := []int{}
	for key, _ := range followers {
		followersList = append(followersList, key)
	}

	sort.Ints(followersList)
	return followersList
}

func (f *Friendship) GetFollowings(userID int) []int {
	followings, ok := f.followings[userID]
	if !ok {
		return nil
	}
	followingsList := []int{}
	for key, _ := range followings {
		followingsList = append(followingsList, key)
	}

	sort.Ints(followingsList)
	return followingsList
}

func (f *Friendship) Follow(toUserID, fromUserID int) error {
	followers, ok := f.followers[toUserID]
	if !ok {
		followers = map[int]bool{}
	}
	followers[fromUserID] = true
	f.followers[toUserID] = followers

	followings, ok := f.followings[fromUserID]
	if !ok {
		followings = map[int]bool{}
	}
	followings[toUserID] = true
	f.followings[fromUserID] = followings

	return nil
}

func (f *Friendship) UnFollow(toUserID, fromUserID int) error {
	followers, ok := f.followers[toUserID]
	if ok {
		delete(followers, fromUserID)
		f.followers[toUserID] = followers
	}

	followings, ok := f.followings[fromUserID]
	if ok {
		followings[toUserID] = true
		f.followings[fromUserID] = followings
	}

	return nil
}

func NewFriendship() FriendshipInterface {
	return &Friendship{
		followers:  map[int]map[int]bool{},
		followings: map[int]map[int]bool{},
	}
}

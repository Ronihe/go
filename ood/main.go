package main

import (
	"fmt"
	"ood/friendship"
	"ood/memcache"
	"ood/miniTwitter"
	"ood/shape"
	"time"
)

type Animal struct {
	Name string
	mean bool // private inside of the package
}

// there is no interitance in go, composition is the only choice

type cat struct {
	Basics Animal
	Meow   int
}

type Dog struct {
	Animal // embeding, unnamed struct
	Bark   int
}

func (animal *Animal) PerformNoise(strength int, sound string) {
	if animal.mean == true {
		strength = strength * 5
	}

	for voice := 0; voice < strength; voice++ {
		fmt.Printf("%s ", sound)
	}

	fmt.Println("")
}

func (dog *Dog) MakeNoise() {
	dog.PerformNoise(dog.Bark, "bark")
}

func (cat *cat) MakeNoise() {
	cat.Basics.PerformNoise(cat.Meow, "meow")
}

// When using the Dog reference we access the Animal fields directly. With the Cat reference we use the named field called Basics.

// polymorphic
type AnimalSounder interface {
	//There is a Go convention of naming interfaces with the "er" suffix when the interface only contains one method.

	MakeNoise()
}

func MakeSomeNoise(animalSounder AnimalSounder) {
	animalSounder.MakeNoise()
}

var myDog *Dog = &Dog{
	Animal: Animal{
		Name: "Lola",
		mean: false,
	},
	Bark: 5,
}

var myCat *cat = &cat{
	Basics: Animal{
		Name: "Julius",
		mean: true,
	},
	Meow: 3,
}

func main() {
	MakeSomeNoise(myDog)
	MakeSomeNoise(myCat)
	var s shape.Shape = shape.NewRect(6, 7)
	fmt.Println(s.Area())

	fmt.Println("------minitwitter----")
	minit := miniTwitter.NewMniniTwitter()
	err := minit.Follow(1, 2)
	minit.PostTweet(1, "user 1 first post")
	time.Sleep(10000)
	minit.PostTweet(1, "user 1 second post")
	time.Sleep(10000)

	minit.PostTweet(1, "user 1 third post")

	time.Sleep(1000)
	minit.PostTweet(1, "user 1 forth post")
	time.Sleep(10000)
	minit.PostTweet(1, "user 1 fifth post")
	time.Sleep(10000)
	minit.PostTweet(1, "user 1  sixth post")
	time.Sleep(10000)
	minit.PostTweet(1, "user 1 seventh post")
	time.Sleep(10000)
	minit.PostTweet(1, "user 1 eighth post")
	time.Sleep(10000)
	minit.PostTweet(1, "user 1 ninth post")
	time.Sleep(1000)
	minit.PostTweet(1, "user 1 tenth post")
	time.Sleep(1 * time.Second)
	minit.PostTweet(1, "user 1 eleventh post")
	time.Sleep(1 * time.Second)

	if err != nil {
		fmt.Println(err)
	}

	x := minit.GetNewsFeed(1)
	fmt.Println(x)

	fmt.Println("------friendship-----")
	f := friendship.NewFriendship()
	f.Follow(1, 2)
	fs := f.GetFollowers(1)
	fmt.Println("1's followers", fs)

	fmt.Println(".....memcache----")
	m := memcache.NewMemcache()
	val := m.Set(1, 222, 30)
	fmt.Println("memcahce", val)
}

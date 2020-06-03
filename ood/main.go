package main

import (
	"fmt"
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
}

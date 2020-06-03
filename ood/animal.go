
package main

type Animal struct{
	Name string
	mean bool // private inside of the package
}

type cat struct{
	Basics Animal
	Meow int
}

type Dog struct{
	Animal // embeding
	Bark int
}


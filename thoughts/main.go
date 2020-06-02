package main

import "fmt"

var h int
var i string
var j float64
var k bool

func main() {
	a := 34
	b := "Learning Golang"
	c := 5.23
	d := true
	e := "Welcome"
	f := `How are you?`
	g := 'S'

	var info string
	info = "welcome to go"

	var l, m, n string
	l = "google"

	var o, p int = 32, 34

	var q, r, s = false, 54.56, "H"
	t, u, v := false, 54.56, 'H'

	w := `this is
	multiple lines`

	fmt.Printf("Type:%T \t Value:%v \n", a, a)
	fmt.Printf("Type:%T \t Value:%v \n", b, b)
	fmt.Printf("Type:%T \t Value:%v \n", c, c)
	fmt.Printf("Type:%T \t Value:%v \n", d, d)
	fmt.Printf("Type:%T \t Value:%v \n", e, e)
	fmt.Printf("Type:%T \t Value:%v \n", f, f)
	fmt.Printf("Type:%T \t Value:%v \n", g, g)
	fmt.Println()
	fmt.Println("Default Values---")
	fmt.Println("int:", h)
	fmt.Println("string:", i) //string has empty value
	fmt.Println("float64:", j)
	fmt.Println("bool:", k)
	fmt.Println()
	fmt.Println(info)
	fmt.Println("l,m,n:", l, m, n)
	fmt.Println("o,p:", o, p)
	fmt.Println("q,r,s:", q, r, s)
	fmt.Println("t,u,v:", t, u, v)
	fmt.Println("backtricks w:", w)

}

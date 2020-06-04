package main

import "fmt"

// Golang just have one looping constuct
func basicLoop(times int) {
	for i := 0; i < times; i++ {
		fmt.Println("times", i)
	}
}

func whileLoop(sum int) {

	init := 1
	for init < sum {
		init += init
	}
	fmt.Println(init)
}

func multiInitLoop(a, b int) {
	for i, j := 0, 1; i <= a && j <= b; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

}

// for each range loop
func rangeLoop(length int) {
	nums := make([]int, length)

	for i, val := range nums {
		fmt.Println("range", i, val)
	}
}

// exit the for loop with continue, break
func breakLoop(max int) {
	for i := 0; i < max; i++ {
		if i >= 5 {
			break
		}
		
		if i%2 != 0 {
			continue
		}
		fmt.Println("even:", i)
	}
}

func main() {
	basicLoop(3)
	whileLoop(5)
	multiInitLoop(10, 18)
	rangeLoop(3)
	breakLoop(100)

}

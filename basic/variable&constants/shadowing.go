package main

import "fmt"

func main() {
	x := 10
	if true {
		x := 20 // shadowing
		fmt.Println(x) // 20
	}
	fmt.Println(x) // 10	
}
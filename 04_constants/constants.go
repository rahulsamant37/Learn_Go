package main

import "fmt"

const age = 22

func main() {
	const name string = "golang"
	// name = "javascript" if you give a value to const then it's value will not change and you can declar a value outside the function

	

	fmt.Println(name)
	fmt.Println(age)

}
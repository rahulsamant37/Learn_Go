package main

import "fmt"

// for -> only construct in go for looping
func main() {
	//while loop
	println("-----------## While loop ##-----------")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i+=1
	}

	
	//infinite loop
	// for {
		// 	println("1")
		// }
		
	// Classic For loop
	println("-----------## Clasic For Loop ##-----------")
	for i:=0; i<3; i++ {
		fmt.Println(i)
	}

	// go 1.2> have new feature which is `range`
	println("-----------## range in loop ##-----------")
	for i := range 3 {
		fmt.Println(i)
	}
}
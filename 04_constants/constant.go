package main

import "fmt"

// bahar bhi constant declare kar sktre hain
const age = 30 
// variable ko bhi declare kar skte but shorthand me ni kr skte 
// name:="yash"
// var name string = "golang"

func main(){
	// const name string = "golang"
	const name  = "golang"
	fmt.Println(age)

	// Multiple constants group
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port,host)
	
}

package main

import "fmt"

func main(){
	// var keyword  name  type = value
	// var name string = "golang"
	// fmt.Println(name)

	var name  = "golang"
	// jab type ni dete direct value dete : type infer krleti hai 
	fmt.Println(name)

	// var isAdult bool= true
	var isAdult = true
	fmt.Println(isAdult)

    // hum int hi use karte golang internally hi optmize karleti hai according to internal architecture
	var age int = 22
	fmt.Println(age)

	// Shorthand syntax
	nameShort :="golang"
	fmt.Println(nameShort)


	// Situation jahan humko poora use karne ka need hoga : jb sirf variable ko define kr rhe hain toh
	// yhan name ki value abhi assign ni karni hai badme karni hai 

	var nameLater string 
	nameLater = "Golang"

	fmt.Println(nameLater)
    
	// float 

	// var price float32 = 50.5
	// fmt.Println(price)

	// var price  = 50.5
	// fmt.Println(price)

	price:=50.5
	fmt.Println(price)

}
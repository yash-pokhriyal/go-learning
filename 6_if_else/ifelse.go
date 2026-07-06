package main

import "fmt"

func main(){
	// age:=16
	// if age >=18{
	// 	fmt.Println("Person is adult")
	// }else{
	// 	fmt.Println("Person is not adult")
	// }


	// age:=11

	// if age>=18{
	// 	fmt.Println("Person is adult")
	// }else if age>=12{
	// 	fmt.Println("Person is teen")
	// }else{
	// 	fmt.Println("Person is a kid")
	// }

	// var role = "admin"
	// var hasPermissions = false
	
	// // or operater
	// if role == "admin" || hasPermissions{
	// 	fmt.Println("Yes")
	// }
	// // and operater
	// if role == "admin" && hasPermissions{
	// 	fmt.Println("Yes")
	// }

	// Hum variable direct if ke andar declare kar skte hain 
    if age :=11;age>=18{
		fmt.Println("Person is an adult",age)
	} else if age >=12{
		fmt.Println("Person is a teenager",age)
	} else{
		fmt.Println("Person is a kid",age)
	}

}

// agar wo import package hum use ni krenge save krne pe automatic delete ho jaega : magic done by extension we install
// go deoesnot have a ternery operater : 1.22 version

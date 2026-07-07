package main

import (
	"fmt"

	"github.com/yash-pokhriyal/learning-package/auth"
	"github.com/yash-pokhriyal/learning-package/user"
	// "github.com/fatih/color"
)

// Kyu | dry priciple
// naming conflicts me help
// code better organize
// compilation time speed up
// unhi package ko recompile karta go jisme changes hue hain


func main(){
	auth.LoginWithCredentials("codersgyaan","secretpass")
	session:=auth.GetSession()
	fmt.Println("session",session)
 

	user:=user.User{
		Email: "user@email.com",
		// Name:"John Doe",

	}
	// fmt.Println(user.Email,user.Name)
	// color.Red(user.Email)

	// go mod tidy se fix
}

// terminal go mod init modulename
// go mod init github.com/codersgyan/podcast 
//  go mod init github.com/yash-pokhriyal/learning-package

// go get github.com/fatih/color
// go package se laa skte hain



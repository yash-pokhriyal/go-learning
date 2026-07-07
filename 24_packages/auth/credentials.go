package auth

import "fmt"

// package or folder ka naam same rkhte by convention

// func loginWithCredentials(username string , password string){
// 	// abhi ye sirf apne package me hi accessible
// 	fmt.Println("Login user using username and password",username,password)
// }

// bahar se accessible bnane ke liye func ko capital se likhte
func LoginWithCredentials(username string , password string){
	// abhi ye sirf apne package me hi accessible
	fmt.Println("Login user using username and password",username,password)
}
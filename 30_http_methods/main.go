package main

import (
	"fmt"
	"net/http"
)

// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Home Page")
// 	// Method check karna
// 	fmt.Println(r.Method)
// }

// func main(){
// 	http.HandleFunc("/",home)
	  
// 	fmt.Println("Server started on :8080")
// 	http.ListenAndServe(":8080",nil)

// }

// ab browser ya client server ko request bhejta hai, to uske saath ek method bhi bhejta hai.

// Sabse common methods:

// Method	Kaam
// GET	Data lena
// POST	Naya data bhejna
// PUT	Purana data update karna
// DELETE	Data delete karna


// GET Method

// Sabse common method.

// Browser me jab tum ye likhte ho:

// http://localhost:8080

// to browser by default GET request bhejta hai.


// Sirf get allow krna 


func home(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "Welcome")
}

func main(){
	http.HandleFunc("/",home)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080",nil)
}


// Multiple methods
// func home(w http.ResponseWriter, r *http.Request) {

// 	switch r.Method {

// 	case http.MethodGet:
// 		fmt.Fprintln(w, "GET Request")

// 	case http.MethodPost:
// 		fmt.Fprintln(w, "POST Request")

// 	default:
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 	}
// }



// GET
// GET /users

// ➡️ Sare users lao.

// POST
// POST /users

// ➡️ Naya user banao.

// PUT
// PUT /users/1

// ➡️ User update karo.

// DELETE
// DELETE /users/1

// ➡️ User delete karo.

// Visual
// Client

//       │

// GET /users

//       │

// Server

//       │

// Returns Users
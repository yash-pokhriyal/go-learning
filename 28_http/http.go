package main

import (
	"fmt"
	// Go ka backend package.
	"net/http"
)

// HTTP = Browser aur Backend ke beech communication ka protocol.

// Ye function tab chalega jab koi / open karega.

// r = Browser ki Request
// w = Browser ko Response bhejne ka tarika


// http.Request = Struct ⇒ Pointer (*http.Request)
// http.ResponseWriter = Interface ⇒ Direct (http.ResponseWriter)

func home(w http.ResponseWriter,r *http.Request){
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Host:", r.Host)
	fmt.Println("Browser:", r.UserAgent())
	if r.Method == "GET" {
		fmt.Fprintln(w, "GET Request Received")
	}

	fmt.Fprintln(w, "Hello Backend")
}

func main(){
    // / route ko home function se connect kar diya.
	http.HandleFunc("/",home)
	// Browser ko response bhej rahe hain.
	fmt.Println("Server is running on http://localhost:8080")
	// Server start ho gaya.
	http.ListenAndServe(":8080",nil)

}

// Request
// Jab browser backend se kuch maangta hai, usse Request kehte hain.
// Browser
//    │
//    │ GET /
//    ▼
// Backend


// r *http.Request

// Is r ke andar browser ki saari information hoti hai.

// func home(w http.ResponseWriter, r *http.Request) {
//     // r me request ki details hain
// }

//1. w :Browser ko response bhejne ke liye.
//2. r: Browser se jo request aati hai, uski saari information r ke andar hoti hai.
// Jaise:

// Method
// URL
// Headers
// Body
// Cookies

// 3. r.Method


// Ye batata hai kis type ki request aayi hai.

// Examples:

// GET
// POST
// PUT
// DELETE

// 4. r.URL.Path


// Ye batata hai kaunsa route hit hua.

// Examples:

// /          -> Home
// /about     -> About
// /login     -> Login
// /users/10  -> User 10


// Method	Kaam
// GET	Data lena 📥
// POST	Data  or data create/submit 📤


// Browser
//    │
//    │ GET /products
//    ▼
// Backend
//    │
//    ▼
// Products bhej do


// w (http.ResponseWriter) browser ko response bhejne ke liye use hota hai. r (*http.Request) browser se aayi request ki saari information rakhta hai, jaise Method, URL, Headers aur Body. Request ek bada struct hai, isliye uski copy na bane aur performance achhi rahe, isliye pointer (*) pass kiya jata hai. ResponseWriter ek interface hai, isliye uske saath pointer use nahi karte.
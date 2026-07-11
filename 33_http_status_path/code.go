// package main

// import( "net/http"
// "fmt"
// )


// Status Code kya hota hai?

// Jab client request bhejta hai:

// Client ─────► Server

// Server sirf data nahi bhejta.

// Wo ye bhi batata hai:

// Request successful hui ya fail?

// Ye batane ke liye HTTP Status Code use hota hai.

// |                        Status | Kab use karte hain?         |
// | ----------------------------: | --------------------------- |
// |                    **200 OK** | Request successful          |
// |               **201 Created** | Naya resource create hua    |
// |           **400 Bad Request** | Client ki request galat hai |
// |             **404 Not Found** | Resource nahi mila          |
// | **500 Internal Server Error** | Server ke andar error aaya  |

// // 200
// w.WriteHeader(http.StatusOK)

// // 201
// w.WriteHeader(http.StatusCreated)

// // 400
// http.Error(w, "Bad Request", http.StatusBadRequest)

// // 404
// http.Error(w, "Not Found", http.StatusNotFound)

// // 500
// http.Error(w, "Internal Server Error", http.StatusInternalServerError)

// Jab success hota hai:
// w.WriteHeader(...)
// Jab error hota hai:
// http.Error(...)

// Kyun?

// w.WriteHeader() → Sirf status (aur phir tum khud response likh sakte ho).
// http.Error() → Status + error message dono ek saath bhej deta hai.

// Path parameter

// Step 1

// Abhi tak tum ye use karte the:

// GET /user?id=10

// Isme id kahan hai?

// 👉 URL ke baad ? ke through.

// Isko kehte hain Query Parameter.

// Step 2

// REST APIs me zyada common hota hai:

// GET /users/10

// Yahan 10 URL ka hi part hai.

// Isko kehte hain Path Parameter.

// Difference

// Query Parameter:

// /users?id=10

// Path Parameter:

// /users/10
// Real APIs

// GitHub:

// /users/octocat

// User ID:

// /users/10

// Product:

// /products/25

// Ye sab Path Parameters use karte hain.

package main

import (
	"fmt"
	"net/http"
	"strings"
)

func getUser(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/users/")

	fmt.Fprintln(w, "User ID:", id)
}

func main() {
	http.HandleFunc("/users/", getUser)
	http.ListenAndServe(":8000", nil)
}



// Step 3 (Important)

// ⚠️ Standard net/http package me automatic path parameters nahi hote.

// Agar tum likho:

// http.HandleFunc("/users/", getUser)

// Aur browser kholo:

// http://localhost:8080/users/10

// To request getUser me aa jayegi.

// Lekin 10 automatically alag nahi milega.

// Step 4

// Uske liye hum URL se value nikalte hain.

// id := r.URL.Path

// Agar URL hai:

// /users/10

// To:

// fmt.Println(id)

// Output:

// /users/10
// Step 5

// Ab sirf number chahiye.

// Go me package hai:

// import "strings"

// Aur:

// id := strings.TrimPrefix(r.URL.Path, "/users/")

// Agar URL hai:

// /users/10

// To output:


// Browser:

// http://localhost:8080/users/25

// Output:

// User ID: 25
// 🧠 Important Note

// Ye tarika standard net/http ka hai.

// Real projects me log routers use karte hain jaise:

// chi
// gorilla/mux (legacy)
// gin
// echo

// Unme ye bahut simple hota hai.

// Example (chi):

// id := chi.URLParam(r, "id")

// Lekin pehle standard library seekhna sahi decision hai. Isi se HTTP ka base strong hota hai.



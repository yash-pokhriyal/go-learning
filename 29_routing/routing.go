package main

import (
	"fmt"
	"net/http"
)

// Routing kya hoti hai?

// Jab user kisi URL par request bhejta hai, to server decide karta hai kaunsa function chalana hai.

// func home(w http.ResponseWriter , r *http.Request){
// 	fmt.Fprintln(w,"Welcome to Home page")
// }
// func about(w http.ResponseWriter , r *http.Request){
// 	fmt.Fprintln(w,"Welcome to About page")
// }
// func contact(w http.ResponseWriter , r *http.Request){
// 	fmt.Fprintln(w,"Welcome to Contact page")
// }

// func main(){
// 	http.HandleFunc("/",home)
// 	http.HandleFunc("/about",about)
// 	http.HandleFunc("/contact",contact)
    
// 	fmt.Println("Server started on :8080")
// 	http.ListenAndServe(":8080",nil)

// }


// Browser

//         |

// localhost:8080/about

//         |

// HandleFunc()

//         |

// about()

//         |

// Response

//         |

// Browser


// Query Parameter kya hota hai?

// URL ke through server ko extra data bhejna.

// http://localhost:8080?name=Yash

// /search?q=golang
// /products?category=laptop
// /users?page=2

// URL?key=value

// key   → name
// value → Yash


func home(w http.ResponseWriter , r *http.Request){

	name := r.URL.Query().Get("name")
	city := r.URL.Query().Get("city")


	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "%s lives in %s", name,city)

}
func about(w http.ResponseWriter , r *http.Request){
	fmt.Fprintln(w,"Welcome to About page")
}
func contact(w http.ResponseWriter , r *http.Request){
	fmt.Fprintln(w,"Welcome to Contact page")
}

func main(){
	http.HandleFunc("/",home)
	http.HandleFunc("/about",about)
	http.HandleFunc("/contact",contact)
    
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080",nil)

}
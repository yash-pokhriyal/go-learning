package main

import (
	"fmt"
	"io"
	"net/http"
)


func home(w http.ResponseWriter,r *http.Request){

	body,err:= io.ReadAll(r.Body)
	if err!=nil{
		http.Error(w,"Cannot Read Body",http.StatusBadRequest)
		return
	}
	fmt.Println(string(body))
	fmt.Fprintln(w, "Body Received")

}
func main(){
    http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)


}

// Request Body

// Abhi tak hum data URL se bhej rahe the:

// GET /?name=Yash

// Lekin agar user signup kare to?

// Kya hum ye karenge?

// /signup?name=Yash&email=yash@gmail.com&password=123456

// ❌ Bilkul nahi.

// Isliye POST Request ke andar data Request Body me bheja jata hai.


// Request ke 3 major parts hote hain:

// URL
// Headers
// Body ✅ (Aaj ka topic)

// Body Read Kaise Kare?

// Go me request body r.Body me hoti hai.




// Browser address bar sirf GET request bhejta hai.

// POST request bhejne ke liye use karte hain:

// Postman
// curl
// Thunder Client (VS Code)

// Abhi hum curl use karenge kyunki extra install nahi karna.


// curl -X POST http://localhost:8080 \
// -d "Hello Go Backend"

// Body sirf String nahi hoti

// Real applications me body mostly JSON hoti hai.

// r.Body

// ➡️ POST request ka data.


// 1. r.Body

// Yaad hai r kya tha?

// func home(w http.ResponseWriter, r *http.Request)

// r = Request jo client ne bheji.

// Agar client ne ye request bheji:

// POST / HTTP/1.1

// Hello Yash

// To r.Body ke andar hoga:

// Hello Yash

// Agar JSON bheji:

// {
//     "name":"Yash",
//     "age":21
// }



// body,err:= io.ReadAll(r.Body)
// To r.Body ke andar wahi JSON hogi.

// Important: r.Body directly string nahi hota.

// Uska type hota hai:

// io.ReadCloser
// 2. io.ReadAll()

// r.Body ek stream hai.

// Socho tumhare paas ek pipe hai.

// Client

// Hello Yash
// ──────────────► Pipe (r.Body)

//                   ▼

//                Server

// Pipe ko seedha print nahi kar sakte.

// Pehle usme se saara data nikalna padega.

// Ye kaam karta hai:

// io.ReadAll()

// Matlab:

// "Pipe ke andar jo bhi data hai, sab padh lo."

// 3. body
// body, err := io.ReadAll(r.Body)

// body ka type hota hai

// []byte

// Yaani bytes ka slice.

// Agar request thi:

// Hello Yash

// to internally

// [72 101 108 108 111 ...]

// store hota hai.

// Isliye hum likhte hain:

// fmt.Println(string(body))

// Output:

// Hello Yash

// Kyuki bytes ko string me convert kiya.

// 4. err

// Reading fail bhi ho sakti hai.

// Isliye Go return karta hai

// body, err

// Agar sab sahi hai

// err == nil

// Agar problem hai

// err != nil

// Tab

// http.Error(w, "Cannot Read Body", http.StatusBadRequest)
// return
// Visual Flow

// Client request:

// POST /

// Hello Yash

// ↓

// Server

// r.Body

// ↓

// Pipe

// ↓

// io.ReadAll(r.Body)

// ↓

// []byte

// ↓

// string(body)

// ↓

// Hello Yash
// Ek analogy 🚰

// Socho tumhare ghar me water pipe hai.

// Pipe
//  │
//  ▼
// Water

// r.Body = Pipe

// io.ReadAll() = Pipe ka saara paani bucket me bhar lo.

// body = Bucket

// string(body) = Bucket ka paani dekhna.

// Is line ka matlab ek sentence me
// body, err := io.ReadAll(r.Body)

// 👉 Client ne request body me jo bhi data bheja hai, use poora padhkar body variable me store kar do. Agar padhne me koi problem aaye to err me bata do.

// Curl command ka breakdown
// curl

// 👉 HTTP request bhejne ka tool.

// -X POST

// 👉 Method ko POST bana do.

// http://localhost:8080

// 👉 Kis server ko request bhejni hai.

// -d "Hello Go Backend"
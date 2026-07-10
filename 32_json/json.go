package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

// JSON hota kya hai?

// {
//     "name": "Yash",
//     "age": 21
// }

// Ye sirf data ko represent karne ka format hai.


// type User struct{
// 	Name string
// 	Age int
// }

// func main(){
// 	data :=[]byte(`{
// 	    "Name":"Yash",
// 		"Age":22
// 	}`)
// 	var user User
// 	// unmarshal byte leta hai data ko user mai dera
// 	err := json.Unmarshal(data, &user)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(user)
// }

// json tags 
// Tumne notice kiya hoga hum JSON me likh rahe the:

// {
//     "Name":"Yash"
// }

// Lekin real APIs me JSON aise hoti hai:

// {
//     "name":"Yash"
// }

// 🤔 Agar struct me field ka naam Name hai aur JSON me key name hai, to Go ko kaise pata chalega ki dono same field hain?

// 👉 Isi problem ko solve karte hain JSON Tags. Ye har Go backend developer roz use karta hai.


// type User struct{
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// }

// func main(){
// 	data :=[]byte(`{
// 	    "name":"Yash",
// 		"age":22
// 	}`)
// 	var user User
// 	// unmarshal byte leta hai data ko user mai dera
// 	err := json.Unmarshal(data, &user)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(user)
// }


// marshal struct -> json
// type User struct{
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// }

// func main(){
// 	user := User{
// 	Name: "Yash",
// 	Age: 22,
// 	}

// 	data, err := json.Marshal(user)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(data))

// }

// Go ka encoding/json package bytes ([]byte) ke saath kaam karta hai.

// Unmarshal ➜ []byte leta hai
// Marshal ➜ []byte return karta hai

// Isliye Go backend me sab naturally fit ho jata hai.




type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	user := User{
		Name: "Yash",
		Age: 22,
	}

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}


// Mini Revision

// Ye line dekho:

// json.NewEncoder(w).Encode(user)

// Isme:

// w = Response ka raasta (client tak)
// NewEncoder(w) = JSON bhejne ke liye encoder banata hai
// Encode(user) = Struct → JSON → Client
// Go Struct ko JSON me convert karke client ko bhej do.


// Go Struct

// ↓

// Marshal

// ↓

// JSON

// ↓

// w.Write()

// ↓

// Client

// json.NewEncoder(w)

// w ke through JSON response bhejne ke liye encoder banata hai.

// Encode(user)

// user struct ko JSON me convert karke client ko bhej deta hai.
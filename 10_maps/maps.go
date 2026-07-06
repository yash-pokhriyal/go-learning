package main

import (
	"fmt"
	"maps"
)

// maps -> associative data structure hota hai
// maps:- hash , object , dictionary

func main(){

	// creating map
	// map[keytype]valuetype
	// m:=make(map[string]string)

	// setting element
	// m["name"]="golang"
	// m["area"]="backend"

	// getting element
	// fmt.Println(m["name"],m["area"])

	// If key does not exist in map : returns 0 value
	// fmt.Println(m["phone"])

	// m:= make(map[string]int)
	// m["age"] = 30
	// m["price"]=3000


	// fmt.Println(m["age"])
	// fmt.Println(m["phone"])
	// fmt.Println(len(m))
    
	// to delete from map
	// fmt.Println(m)
	// delete(m,"price")
	// fmt.Println(m)

	// map ko khali krne ke liye : clear(m)
	// clear(m)
	// fmt.Println(m)


	// map witout using make function 

	// m:=map[string]int{"price":3000,"phones":3}
	// fmt.Println(m)

	// to chec element in map
    // go ke andar multiple cheezein return hoti hain
	// pehle wale ke andar value dusra wala boolean deta
	// v,ok :=m["price"]
	// fmt.Println(v)

	// if ok{
	// 	fmt.Println("All ok")
	// }else{
	// 	fmt.Println("Not ok")
	// }


	m1:=map[string]int{"price":3000,"phones":3}
	m2:=map[string]int{"price":3000,"phones":3}

	fmt.Println(maps.Equal(m1,m2))




 
}

package main

import "fmt"

// Jo bhi variable ya data structure hum store karte hain apne computer memory ke andar pointers basically  uss variable ki jo memory location hai uska ye address hota hai


// by value pass hojata hai
// func changeNum(num int){
// 	num=5
// 	// ye copy change hori naaki source variable 
// 	fmt.Println("In changeNum",num)
// }

// by reference
func changeNum(num *int){
	// dereference : means is address ke value ko change karna 
	*num=5

	fmt.Println("In changeNum",*num)
}

func main(){
   // num:=1
   // changeNum(num)
   // fmt.Println("After changing num",num)
   // In changeNum 5
   // After changing num 1

    // kyu hmara number change ni hua
    // solution abv aese number pass krne ki jageh uska reference pass krenge function me

   num:=1
   fmt.Println("Memory address",&num)
   //Memory address 0x62bef616020
   //& ye lgake memory address get karte hain
   changeNum(&num)
   fmt.Println("After changing num",num)
   
}

// & = "Address do."
// * = "Address se value lao."

// *int → "pointer to an int" (address store karega)
// *num → "us address par jo value hai, usko access ya modify karo"
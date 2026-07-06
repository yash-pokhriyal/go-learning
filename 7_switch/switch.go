package main

import (
	"fmt"
	// "time"
)

func main(){
	// simple switch 
	// i :=3
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
    // default:
	// 	fmt.Println("Other")

	// }


	// multiple condition switch 
	// switch time.Now().Weekday(){
	// case time.Saturday,time.Sunday:
	// 	fmt.Println("Its weekend")
	// default:
	// 	fmt.Println("Its workday")
	// }


	// type switch 
	// first class function 
	// interface{} :- means any type of value
	// whoAmI:=func(i interface{}){
    //    switch t:=i.(type){
	//    case int :
	// 	fmt.Println("Its an integer")
	//    case string:
	// 	fmt.Println("Its a string")
	//    case bool :
	// 	fmt.Println("Its a boolean")
	//    default:
	// 	fmt.Println("Other",t)
	//    }
	// }

	// whoAmI("golang")
	// whoAmI(6)
	// whoAmI(6.999)

	// Agar value use ni karni to direct ye krlo
	whoAmI:=func(i interface{}){
       switch i.(type){
	   case int :
		fmt.Println("Its an integer")
	   case string:
		fmt.Println("Its a string")
	   case bool :
		fmt.Println("Its a boolean")
	   default:
		fmt.Println("Other")
	   }
	}

	whoAmI("golang")
	whoAmI(6)
	whoAmI(6.999)
}

// Baaki languages me break likhna pdta par go automatically behind the scenes handle karleta
// default case is optional
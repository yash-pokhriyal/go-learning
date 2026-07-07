package main

import "fmt"

// enumerated types
// we can create our custom type in go:- type Mytype string

// type OrderStatus int
type OrderStatus string

const(
	// Recieved OrderStatus = iota
	// Confirmed 
	// Prepared
	// Delivered

	// for string
	Recieved OrderStatus = "recieved"
	Confirmed ="confirmed"
	Prepared = "prepared"
	Delivered = "delivered"
)


func changeOrderStatus(status OrderStatus){
	fmt.Println("Changing order status to",status) 
}
func main(){
	// pehle hmne func type string krke kara but uske kai dikkat hai 
	// changeOrderStatus("recieved")
	// changeOrderStatus(Recieved)
	changeOrderStatus(Prepared)
 
}

// we implement enums in go using const

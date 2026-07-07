package main

// structs -> custom data structure hote hain
// classes of go

// order struct
// type structname struct

// works as a blueprint

import (
	"fmt"

	"time"
)

// type order struct{
// 	id string
// 	amount float32
// 	status string
// 	createdAt time.Time //nanosecond precission

// }

// // constructor in go func new/New(Struct name)
// func newOrder(id string,amount float32,status string)*order{
// 	// initial setup goes here ....
// 	myOrder:=order{
// 		id:id,
// 		amount: amount,
// 		status: status ,

// 	}

// 	return &myOrder
// }

// method
// func (reciver type:ise se hi struct ke saath jodte : yhan struct ka pehla letter likhte:structname) changeStatus(status string)

// func (o order) changeStatus(status string){
//     o.status = status
// }
// func (o *order) changeStatus(status string){
// 	// struct hmare liye dereferencing ka kaam automatically krta hai
//     o.status = status
// }
// func (o *order) getAmount() float32{
// 	return o.amount
// }
// yhan hum bina * ke  bhi kar skte hain uske bina bhi kaam krega
// * | use karo jab value modify krni ho


// struct embedding 

type customer struct{
	name string
	phone string
}
// composition
type order struct{
	id string
	amount float32
	status string
	createdAt time.Time //nanosecond precission
	customer

}

func main(){

	// struct ka instance 
	// var order order 
	// myOrder:=order{
	// 	id:"1",
	// 	amount: 50.00,
	// 	status: "recieved" ,

	// }
	// myOrder.createdAt = time.Now()
    
	// fmt.Println(myOrder.status)
	// fmt.Println("Order struct",myOrder)

	// myOrder2 :=order{
	// 	id: "2",
	//     amount: 60.00,
	// 	status: "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder2.status = "paid"

	// fmt.Println("Order struct",myOrder2)

	// myOrder2.changeStatus("confirm")
	// fmt.Println(myOrder2)
	// fmt.Println(myOrder2.getAmount())

	// agar amount ki value set ni krenge to zero value aegi
	//  if you dont set any field then the default value is the zero value


	// myOrder := newOrder("1",30.50,"recieved")
	// fmt.Println(myOrder.amount)


	// agar struct bs ek baar hi use hora ho
	// language :=struct{
	// 	name string
	// 	isGood bool

	// }{"golang",true}

	// fmt.Println(language)
    
	// newCustomer :=customer{
	// 	name:"yash",
	// 	phone:"09899",
	// }
	newOrder :=order{
		id:"1",
		amount: 50.00,
		status: "recieved" ,
		// customer: newCustomer,
		customer: customer{
			name:"yash",
		phone:"09899",
		},
	}
    newOrder.customer.name = "robin"
	fmt.Println(newOrder)
	fmt.Println(newOrder.customer)
	// {1 50 recieved {0 0 <nil>} { }}

 
}

package main

import "fmt"

// type payment struct{}


// // violating open close principle :- classes methods they should be open for extension but closed for modification
// func (p payment) makePayment(amount float32){
// 	// razorPayPaymentGw := razorPay{}
	
// 	// razorPayPaymentGw.pay(amount)

// 	stripePaymentGw := stripe{}
// 	stripePaymentGw.pay(amount)

   
// }

// type razorPay struct{}

// func (r razorPay) pay(amount float32){
// 	// logic to make payment
// 	fmt.Println("Making payment using razorpay",amount)
// }

// type stripe struct{}

// func (s stripe) pay(amount float32){
// 	fmt.Println("Making payment using stripe",amount)
// }

// func main(){
// 	newPayment:=payment{}
// 	newPayment.makePayment(100)
 
// }



// Improving code 

// type payment struct{
// 	// gateway stripe 
// 	// gateway razorPay
// 	gateway fakePayment
// }



// func (p payment) makePayment(amount float32){

// 	p.gateway.pay(amount)

   
// }

// type razorPay struct{}

// func (r razorPay) pay(amount float32){
// 	// logic to make payment
// 	fmt.Println("Making payment using razorpay",amount)
// }

// type stripe struct{}

// func (s stripe) pay(amount float32){
// 	fmt.Println("Making payment using stripe",amount)
// }

// // testing ke time fake payment gateway pass krte hain
// type fakePayment struct{}

// func (f fakePayment)pay(amount float32){
// 	fmt.Println("Making payment using fake gateway for testing purpose",amount)
// }

// func main(){
// 	// stripePaymentGw := stripe{}
// 	// razorPayPaymentGw :=razorPay{}
// 	fakePaymentGw:=fakePayment{}
// 	newPayment:=payment{
// 		// gateway: stripePaymentGw,
// 		// gateway: razorPayPaymentGw,
// 		gateway: fakePaymentGw,
// 	}
// 	newPayment.makePayment(100)
 
// }
// isme bhi har baar ek hi hora hai : - iska solution hai interfaces



// Interfaces :- contracts hote
// convention :- er at last
type paymenter interface {
	// ek method deni jo bhii struct isko implement krega uske andar ek method honi chahiye
	pay(amount float32) //iske aage return value de skte
	refund(amount float32,account string)
}

type payment struct{
	gateway paymenter
}



func (p payment) makePayment(amount float32){

	p.gateway.pay(amount)

   
}

type razorpay struct{}

func (r razorpay) pay(amount float32){
	// logic to make payment
	fmt.Println("Making payment using razorpay",amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32){
// 	fmt.Println("Making payment using stripe",amount)
// }

// testing ke time fake payment gateway pass krte hain
type fakepayment struct{}

func (f fakepayment)pay(amount float32){
	fmt.Println("Making payment using fake gateway for testing purpose",amount)
}

type paypal struct{

}
func (p paypal) pay(amount float32){
	fmt.Println("Making payment using paypal",amount)
}

func (p paypal) refund(amount float32 , account string){

}

func main(){
	// stripePaymentGw := stripe{}
	// razorPayPaymentGw :=razorpay{}
	// fakePaymentGw:=fakepayment{}
	paypalPaymentGw:=paypal{}
	newPayment:=payment{
		// gateway: stripePaymentGw,
		// gateway: razorPayPaymentGw,
		// gateway: fakePaymentGw,
		gateway: paypalPaymentGw,
	}
	newPayment.makePayment(100)
 
}

// dusri languages ke andar agar kisi bhi intrerfaces ko implement krna hai to hum classes ke upar implement likhte hain
// aur interface ka naam likhte hai and go ke andar ke implicitly hota hai 
// interface ke andar jo method uski signature jo param type 
// same hum bnaenge method struct ke upar wo automatically implement hojata


// hmne dependency inversion kia 

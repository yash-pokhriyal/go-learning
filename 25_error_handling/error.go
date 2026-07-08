package main

import "fmt"

// Go ka philosophy hai:

// Errors are values.

func divide(a,b int) (int,error){
	if b==0{
		return 0,fmt.Errorf("Cannot divide by zero")
		// fmt.Erro() : error object bna deta hai
	}
	return a/b ,nil
}
func main(){

	ans,err:=divide(10,0)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(ans)


}

// Function Call
//       │
//       ▼
// (value,error)
//       │
//       ▼
// err == nil ?
//    │
//  ┌─┴──────┐
//  │        │
// Yes      No
//  │        │
// Use     Handle
// Value   Error


// errors.New() → Fixed/simple error message.
// fmt.Errorf() → Formatted message bana sakte ho (%d, %s) aur baad mein %w se error wrapping bhi kar sakte ho.


// nil ka matlab hai:

// Koi error nahi hui. ✅


// Defer stack ki tarah kaam karta hai (LIFO - Last In, First Out).



// 💡 Hint: defer function ko baad mein chalata hai, lekin arguments usi waqt evaluate ho jaate hain jab defer likha jata hai.

// defer normal function call mein arguments turant evaluate karta hai, lekin deferred anonymous function variables ki latest value function ke end mein read karta hai.

// defer fmt.Println(x) → value capture hoti hai immediately.
// defer func(){ fmt.Println(x) }() → variable baad mein read hota hai.

// panic program ko immediately stop kar deta hai aur stack unwind karna shuru kar deta hai.

// panic aane ke baad bhi saare deferred functions execute hote hain, uske baad program crash hota hai.


// recover ka kaam hai panic ko catch karna, taaki program crash na ho.




// import "fmt"

// func test() {

// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered:", r)
// 		}
// 	}()

// 	panic("Something went wrong")

// 	fmt.Println("Hello")
// }

// func main() {
// 	test()
// 	fmt.Println("Program continues...")
// }

// output 
// Recovered: Something went wrong
// Program continues...


// test()
//    │
//    ▼
// panic()
//    │
//    ▼
// defer chalta hai
//    │
//    ▼
// recover() panic ko catch karta hai
//    │
//    ▼
// test() return ho jata hai
//    │
//    ▼
// main() continue karta hai
//    │
//    ▼
// Program continues...

// panic → Program ko crash karta hai.
// recover → Sirf defer ke andar kaam karta hai aur panic ko crash hone se bacha leta hai.
// <nil> direct recover() likhenge to



// func main() {

	// defer fmt.Println("A")

	// defer func() {
	// 	fmt.Println("B")
	// 	recover()
	// }()

	// panic("Error")

	// fmt.Println("C")
// }

// Execution
// defer fmt.Println("A") register ho gaya.
// Dusra defer bhi register ho gaya.
// panic("Error") aaya.
// Ab defer LIFO mein chalenge.


// Pehle:

// defer func() {
//     fmt.Println("B")
//     recover()
// }()

// Output:

// B

// recover() panic ko handle kar leta hai.

// Phir dusra defer:

// fmt.Println("A")

// Output:

// A

// Program crash nahi karega kyunki panic recover ho chuka hai.

// Final Output
// B
// A


// r := recover()
// fmt.Println(r)

// To output hota:

// B
// Error
// A

// Kyuki recover() panic ki value return karta hai ("Error"), aur panic ko stop bhi kar deta hai.


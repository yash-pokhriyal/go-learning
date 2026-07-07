package main

import "fmt"

func counter() func() int{
	var count int = 0
	return func() int{
		count +=1
		return count
	}
}

func main(){
  increment:=counter()
  fmt.Println(increment())
  fmt.Println(increment())
  fmt.Println(increment())
}

// Humhe pata hai jese bhi koi function execute hoti call stack pe jaati hai
// ek baar execution hote hi hatt jati hai 
// but yhan is case me function return hori hai 
// Closure (Go): Function jo apne outer variable ko yaad rakhta hai, isliye uski value har call ke baad bhi save rehti hai.

// Closure in Go is a function that captures and remembers variables from its outer scope, even after the outer function has returned. This allows the function to preserve state between multiple calls.

// "Jaise counter() function me count ek local variable hai. Normally function return hone ke baad local variables destroy ho jate hain, lekin returned anonymous function count ko use karta hai, isliye Go us variable ko preserve rakhta hai. Har baar increment() call karne par wahi count update hota hai (1, 2, 3...), isliye closure state maintain karta hai."

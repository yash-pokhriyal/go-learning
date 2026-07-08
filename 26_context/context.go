package main


// 1. context.Context
// Ek line

// Context request ko control karta hai (cancel, timeout, deadline).
// Imagine:

// Client
//    │
//    ▼
// API
//    │
// Database (10 sec query)

// Agar client 2 second mein browser band kar de.

// Database ko bhi ruk jana chahiye.

// Isi liye Context.



// Level 2 - Background()

// Har context ki shuruaat yahan se hoti hai.



// import (
// 	"context"
// 	"fmt"
// )

// func main() {

// 	ctx := context.Background()

// 	fmt.Println(ctx)
// }

// Background()
//       │
//       ▼
// Sab naye Context isi se banenge



// import (
// 	"context"
// 	"fmt"
// )

// func main() {

// 	ctx, cancel := context.WithCancel(context.Background())

// 	fmt.Println(ctx.Err())
//     // nil
// 	// No Error
//     // Context Active
	
// 	cancel()
// 	// Context band ho gaya.
// 	// Output context canceled

// 	fmt.Println(ctx.Err())

// }






import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {

	for {

		select {

		case <-ctx.Done():

			fmt.Println("Worker stopped")

			return

		default:

			fmt.Println("Working...")

			time.Sleep(time.Second)

		}

	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(3 * time.Second)

	cancel()

	time.Sleep(2 * time.Second)

}


// Step by Step

// Worker start hua.

// Working...

// 1 second baad

// Working...

// 2 second baad

// Working...

// 3 second baad

// Main function

// cancel()

// Call karta hai.

// Ab

// ctx.Done()

// signal de deta hai.

// Worker bolta hai

// Worker stopped

// Aur function return.

// Final Output
// Working...
// Working...
// Working...
// Worker stopped


// ctx.Done() kya hai?

// Ye ek channel hai.

// Jab context cancel hota hai

// ctx.Done()

// automatically close ho jata hai.

// Isliye

// case <-ctx.Done():

// execute ho jata hai.

// Yehi wajah hai ki context aur channels ka relation bahut strong hai.



// WithTimeout()

// Ab maan lo API ko maximum 5 second hi dene hain.

// ctx, cancel := context.WithTimeout(
// 	context.Background(),
// 	5*time.Second,
// )

// defer cancel()

// Agar worker 5 sec tak complete nahi hua

// Go khud

// cancel()

// kar dega.

// Example

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {

// 	ctx, cancel := context.WithTimeout(
// 		context.Background(),
// 		3*time.Second,
// 	)

// 	defer cancel()

// 	<-ctx.Done()

// 	fmt.Println(ctx.Err())

// }

// Output

// context deadline exceeded



// WithDeadline()

// Difference sirf itna hai.

// Timeout

// 3 seconds

// Deadline

// 2:30 PM tak

// Example

// deadline := time.Now().Add(5*time.Second)

// ctx, cancel := context.WithDeadline(context.Background(), deadline)



// WithValue()

// Suppose login hua.

// User ID pass karni hai.

// ctx := context.WithValue(
// 	context.Background(),
// 	"userID",
// 	101,
// )

// fmt.Println(ctx.Value("userID"))

// Output

// 101



// | Function         | Purpose                |
// | ---------------- | ---------------------- |
// | `Background()`   | Root Context           |
// | `WithCancel()`   | Manual Cancel          |
// | `WithTimeout()`  | Auto cancel after time |
// | `WithDeadline()` | Cancel at exact time   |
// | `WithValue()`    | Request-scoped values  |



// 🔥 Sabse Important Rule (Backend)

// Har backend function mein tum ye dekhoge:

// func GetUser(ctx context.Context, id int) error {
// 	// DB Query
// 	return nil
// }

// Aur HTTP handler mein:

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	// Pass context to DB/API calls
// }

// Rule: Context ko hamesha function ka pehla parameter rakha jata hai aur usse neeche wale functions tak pass kiya jata hai, taaki agar request cancel ya timeout ho jaye to poori chain ko signal mil sake.


// ctx.Done() ek channel hai jo context cancel ya timeout hote hi close ho jata hai. Is channel par wait karne wali sabhi goroutines ko turant cancellation ka signal mil jata hai.


// Boolean polling (baar-baar check) karna padta, jabki ctx.Done() channel bina CPU waste kiye sabhi goroutines ko ek saath signal de deta hai. 🔥

// Channel event-driven hai, boolean polling-based hai. Isliye Go channels use karta hai.



// 🎯 Context Revision (Interview Notes)
// context.Context

// Request ki lifetime ko control karta hai (cancel, timeout, deadline, values).

// Background()
// Root context.
// WithCancel()
// Manual cancel.
// cancel() call karte hi ctx.Done() close.
// WithTimeout()
// Given time ke baad automatically cancel.
// WithDeadline()
// Specific time par cancel.
// WithValue()
// Request-scoped values (UserID, RequestID).
// ctx.Done()
// Channel hai.
// Context cancel/timeout hote hi close hota hai.
// Is par wait karne wali sabhi goroutines ko signal mil jata hai.
// ctx.Err()
// nil → Context active.
// context canceled → Manual cancel.
// context deadline exceeded → Timeout/Deadline.

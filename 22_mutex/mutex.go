package main

import (
	"fmt"
	"sync"
)

// jb bhi hum multthreading krte tab race condition se bchne ke liye utex use krte
// race condition
// jb bhi multiple processes ek same resource ko modify kar rhe hain to vhan pe jo modification hai wo atomic ni hogi
// like 1 ne jo change kia to 2 waliu aake override kre

// type post struct {
// 	views int

// }

// // increment krna views
// func (p *post) inc(){
// 	p.views+=1
// }

// func main(){
// 	myPost:=post{views:0}
// 	// real app me ye concurrently hote hain

// 	myPost.inc()
// 	myPost.inc()
// 	fmt.Println(myPost.views)

// }


// type post struct {
// 	views int

// }

// // increment krna views
// func (p *post) inc(wg *sync.WaitGroup){
// 	defer wg.Done()
// 	p.views+=1
// }


// func main(){
// 	var wg sync.WaitGroup
// 	myPost:=post{views:0}
// 	// simulate multiple krte to kya hoga
// 	for i:=0;i<100;i++{
// 		// aese me dikkt ni hogi jb concurrently krenge tb hogi
//     //    myPost.inc()
// 	   wg.Add(1)
// 	   go myPost.inc(&wg)
// 	//    ab synchronization chahiye hoga saare jo go routine hai uska wait krna pdega khtm hone ka

	
// 	}
  

// 	wg.Wait()
	
// 	fmt.Println(myPost.views)

 
// }

// ab kbhi 100 ,96, etc etc dikhara
// as sab ek saath modify kr rha

// to solve this we use mutex 

type post struct {
	// iske andar hi rkha jae
	views int
	mu sync.Mutex

}

// increment krna views
func (p *post) inc(wg *sync.WaitGroup){
	defer func(){
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.views+=1
	// p.mu.Unlock()
	// good practice defer me rkho
	// good practice ki poora ko lock ni karo sirf usko karo jisme modification hora hai| taaki utni hi jageh bottle nech bne
}


func main(){
	var wg sync.WaitGroup
	myPost:=post{views:0}
	// simulate multiple krte to kya hoga
	for i:=0;i<100;i++{
		// aese me dikkt ni hogi jb concurrently krenge tb hogi
    //    myPost.inc()
	   wg.Add(1)
	   go myPost.inc(&wg)
	//    ab synchronization chahiye hoga saare jo go routine hai uska wait krna pdega khtm hone ka

	
	}
  

	wg.Wait()
	
	fmt.Println(myPost.views)

 
}



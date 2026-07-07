package main

import (
	"fmt"
	"sync"
	// "time"
)

// go routine lightweight threads hote hain
// jabhi bhi multithredding krni hai concurrently cheezopn ko run karna hai wo hum karte hain go routine ki madad se

// func task(id int){
// 	fmt.Println("doing task",id)
// }
// func main(){
// 	for i:=0;i<=10;i++{
// 		// task(i)
// 		// abhi tak ye function blocking hai ab mai chahta hu task function parllely run ho ek saath mai hi run ho jae
// 		go task(i)
// 		// ye rightaway ekdun ru  ni hoti hai wo scheduler pe jaati hai phir schedule isko handle karti hai
// 		// ab aage kuch ni tha to main prpgramme exit hojata hai | to go programme khtm ho jata hai
// 	}

// 	// ab ye likhne pe run ho jaega
// 	time.Sleep(time.Second*2)

// }
// ab cheezein blocking ni hai ye cheezein concurrently run hore
// doing task 6
// doing task 2
// doing task 0
// doing task 8
// doing task 10
// doing task 9
// doing task 7
// doing task 4
// doing task 3
// doing task 1
// doing task 5

// func main(){
// 	for i:=0;i<=10;i++{
// 		go func(i int){
//            fmt.Println(i)
// 		}(i)
// 		// good practice recieve int i and pass on (i)
// 	}

// 	// hum isko aese ni krte hain iske liye go ke andar mechanism hai
// 	time.Sleep(time.Second*2)

// }

// Fix Waitgroups: hmara go routines ko synchronize karne ka ek mechanism hai


func task(id int,w *sync.WaitGroup){
	// go routine end hone ke baad jo add kia tha waitgroup se nikaal dena hai 
	// defer keyword jo bhi hum likhenge wo function execution khtm hone ke baad run hota hai
	// is se hum bolre wait grp ko ki jo ye add kia tha counter is waitgrp se -1 kardega
	defer w.Done()
	fmt.Println("doing task",id)
}

func main(){
	// creating waitgroup
	var wg sync.WaitGroup
	for i:=0;i<=10;i++{
		// kyunki wg pointer hai
		wg.Add(1)
		go task(i,&wg)
	}

	wg.Wait()
    
	

}
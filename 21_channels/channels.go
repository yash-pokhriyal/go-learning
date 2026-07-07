package main

import (
	"fmt"

	// "math/rand"
	// "time"
	// "time"
)

// ye ek pipe ki tarah hota hai | jahan ek side se data ko send krte hain aur dusre side se data ko recieve karte hain
// iska main concept ye hai ki jabhi bhi hmre multiple go routines run hote hain to ye concurrently hore hote hain to uske beech ka matlab ek go routine se dusre go routine ke andar kuch data send karna hai ye recieve karna hai usko hum karte hain channels ke madad se

// go routine ke beech ka communication go routine ki madad se kia jata hai

// func main(){
// 	// make(chan datatype )
// 	messageChan := make(chan string)

// 	// sending data to channel
// 	messageChan <- "ping" //blocking hai

// 	// recieving data
// 	// <- messageChan
// 	// storing in variable

// 	msg:= <-messageChan

// 	fmt.Println(msg)
// }

// fatal error: all goroutines are asleep - deadlock!
// deadlock ek operating system ka concept hai jab bhi hmari multiple process run hori hoti hain concurrently
// to ye processes hai kisi resource ko hold karleti hain jo dusri processes hain wo wait kr hi hai ki kab hmhe miljaega
// par jo pehli wali process hai wo usko release ni kar rhi hai
// Kyunki wo jo hai wo aur kisi resource ka wait kr rhi hai
// infinite waiting me fse hain

// func processNum(numChan chan int){
// 	fmt.Println("processing number",<-numChan)

// }
// func main(){

// 	numChan:=make(chan int)

// 	go processNum(numChan)

// 	numChan<-5

// 	time.Sleep(time.Second*2)
// }

// sending

// func processNum(numChan chan int){
// 	for num:=range numChan{
// 		// range jab use krte tab arrow use krne ki need ni hai
// 		fmt.Println("processing number",num)
// 		time.Sleep(time.Second)

// 	}

// }
// func main(){

// 	numChan:=make(chan int)

// 	go processNum(numChan)

// 	for {
// 		numChan<-rand.Intn(100)
// 	}

// }

// channels ki madat se ek go routine se dusri go routine me data send kr skte hain

// recieving

// func sum(result chan int,num1 int ,num2 int){
// 	numResult:=num1+num2
// 	result <- numResult

// }
// func main(){

// 	result :=make(chan int )
// 	go sum(result,5,4)

// 	// receiving result
// 	res:=<-result //blocking tbhi yhan time sleep lgane ki need ni pdi (partially sach )

// 	fmt.Println(res)

// }

// go routine synchronizer
// func task( done chan bool){
// 	defer func(){done<-true}()
// 	fmt.Println("processing...")
// 	// done ke andar true value send krdi but error ane pe ni chlegi
// 	// done<-true

// }

// channels used for synchronization of go routines
// func main(){
//     // unbuffered channel
// 	//wait grp wala same kaam kar skte hain
// 	done:=make(chan bool )
// 	go task(done)

// 	// ab isko rok ke rkhna hai
// 	// recieving or sending part blocking hota
// 	// value use ni kr rhe hain bas receive kr rhe hain

// 	<-done

// }
// single go routine me channel ka use
// multiple go routine ke liye wait grp


// func emailSender(emailChan chan string, done chan bool){
// 	defer func(){done<-true}()
//
// 	// ye infinite chlega vhi neecheka ka done wait krta rhega->deadlock
// 	for email:=range emailChan{
// 		fmt.Println("sending email to",email)
// 		time.Sleep(time.Second)
// 	}
// }

// func main(){
// 	// real life me struct share kr rhe honge
// 	// second parameter size
// 	// ye hai buffered channel ye blocking ni hoga 
// 	emailChan:=make(chan string,100)
// 	// 100 ka mtlb hi hai itna data hum send kr skte without blocking

// 	done := make(chan bool)

// 	// emailChan<-"1@example.com"
// 	// emailChan<-"2@example.com"

// 	// fmt.Println(<-emailChan)
// 	// fmt.Println(<-emailChan)
    
// 	go emailSender(emailChan,done)
// 	for i:=0;i<5;i++{
// 		emailChan<-fmt.Sprintf("%d@gmail.com",i) //blocking ni hai
// 	}

// 	fmt.Println("Done sending")

// 	// dead lock dikkt khtm| this is important
// 	close(emailChan)

// 	<-done

// 	// normal wale me hum yhan tk pahuch hi ni pae 
// }



// Ek hi saathme multiple channel ke upr listen krna hai
func main(){

	chan1:=make(chan int)
	chan2:=make(chan string)
   
	// yhan hum closure use kr rhe hain
	go func(){
         chan1<-10
	}()
	go func(){
         chan2<-"pong"
	}()

	for i:=0;i<2;i++{
		select{

		case chan1Val:=<-chan1:
			fmt.Println("recieved  data from chan1",chan1Val)
		
		case chan2Val:=<-chan2:
			fmt.Println("recieved  data from chan2",chan2Val)

		}
	
	}

}

// ab emailchan sirf recieve hi kr skte hain aur done send only(data channel pe sirf send kr skte)
func emailSender(emailChan <-chan string, done chan<- bool){
	defer func(){done<-true}()
	
	// ye infinite chlega vhi neecheka ka done wait krta rhega->deadlock
	for email:=range emailChan{
		fmt.Println("sending email to",email)
		time.Sleep(time.Second)
	}
}
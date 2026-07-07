package main

import "fmt"

// ye go ke andar ni the 1.18 me introduced kia gya isko

// func printSlice(items []int){
// 	for _,item:=range items{
// 		fmt.Println(item)
// 	}
// }
// func printStringSlice(items []string){
// 	for _,item:=range items{
// 		fmt.Println(item)
// 	}
// }

// func main(){
// 	printSlice([]int{1,2,3})
// 	names:=[]string{"golang","ts","js"}
// 	printStringSlice(names)


// }

// yhan kaafi duplication hora kyunki bas naam aur type hi change hora hai 

// use krenge generics :- convention T any ki jageh interface{}

// func printSlice[T any](items []T){
// 	for _,item:=range items{
// 		fmt.Println(item)
// 	}
// }

// not good practice ki sbke liye krde
// isiliye aese isme bas int ya string hi pass krte
// func printSlice[T int|string|bool](items []T){
// 	for _,item:=range items{
// 		fmt.Println(item)
// 	}
// }


// func main(){
// 	printSlice([]int{1,2,3})
// 	names:=[]string{"golang","ts","js"}
// 	printSlice(names)
	


// }


// struct

// LIFO
// type stack struct{
// 	elements []int
// }

// func main(){
// 	myStack:=stack{
// 		elements: []int{1,2,3},
// 	}
// 	fmt.Println(myStack)
// }

// ab same stack use krna ho string ke liye to ni kar paenge

// type stack[T any] struct{
// 	elements []T
// }

// func main(){
// 	// [btana pdega T kis type ka hai]
// 	myStack:=stack[string]{
// 		// elements: []int{1,2,3},
// 		elements: []string{"go","js","ts"},
// 	}
// 	fmt.Println(myStack)
// }

// comparable bhi use kar skte hain

// func printSlice[T comparable,V string](items []T,name V){
// 	for _,item:=range items{
// 		fmt.Println(item,name)
// 	}
// }


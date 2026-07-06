package main

import "fmt"

// variadic function
func sum(nums ...int)int{
	total:=0
	for _,num:=range nums{
		total = total+num

	}
	return total

}

// kisi bhi type ka recieve ho jae
// func sum(nums ...interface{})int{

// }


func main(){
	// yhan kitne bhi parameters add kr skte 
	// ye jo function hai isko go mai bolte variadic function
	// fmt.Println(1,2,3)
	// result:=sum(3,4,5,6,7)
	// fmt.Println(result)

	fmt.Println(1,2,3)
	nums:=[]int{3,4,5,6,7}
	result:=sum(nums...)
	fmt.Println(result)
}

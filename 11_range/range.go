package main

import "fmt"

// iterating over data structures
func main(){
     
	// Slice
	// nums:=[]int{6,7,8}
	// for i:=0;i<len(nums);i++{
	// 	fmt.Println(nums[i])
       
	// }
    // sum :=0
	// for index,num:=range nums{
	// 	sum = sum+num
	// 	fmt.Println(num,index)
	// }
	// fmt.Println(sum)


	// Map
	// m:= map[string]string{"fname":"john","lname":"doe"}

	// for k , v :=range m{
	// 	fmt.Println(k,v)
	// }

	// agar single return value deta hu
	// for k :=range m{
	// 	fmt.Println(k)
	// 	// return key only
	// }


	// range on string 
	// c is unicode code point rune
	// i is not index it is starting byte of rune
	// ascii till 255
	// agar 255->1 byte
	for i,c:=range "golang"{
		fmt.Println(i,c)
		// fmt.Println(i,string(c))
	}

 
}

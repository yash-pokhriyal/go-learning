package main

import "fmt"

// arrays: number sequence of specific length

func main(){
	// var nums [4]int 
	// nums[0] = 1
	// fmt.Println(nums[0])
	// fmt.Println(nums)
	// [1 0 0 0]
	// jahan value assign ni hui vhan zeroed value add ho jaega
	

	// getting length
	// fmt.Println(len(nums))

	// var vals [4]bool
	// vals[2] = true
	// fmt.Println(vals)
	// [false false false false]

	// var names [3]string
	// names[0] = "raju"
	// fmt.Println(names)

	// int -> 0 string ->"" bool->false
    
	// to declare it in single line
	// numbers:=[3]int{1,2,3}
	// fmt.Println(numbers)

	// 2d arrays 
	nums:=[2][2]int{{3,4},{5,6}}
	fmt.Println(nums)
	
}

// Useful jab advance me pta ho kitne elements aane wale ho 
// Uselful 
// - fixed size that is predictable
// - memory optimization
// - constant time access

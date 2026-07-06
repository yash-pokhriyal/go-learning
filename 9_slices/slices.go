package main

import (
	"fmt"
	// "slices"
)

// slice :- dynamic arrays
// most used construct in go
// + useful methods

func main(){
	// uninitialized slice is nil 
	// var nums []int 
	// fmt.Println(nums)
	// fmt.Println(nums==nil)
	// fmt.Println(len(nums))

	// hum chahte ye nil na ho to : make()
	// make([]type initialsize capacity)
	// var nums = make([]int,2,5)
	// iske karan 2 initial 0 bnre hai aur normaly initial size 0 rkhte hain
	// fmt.Println(nums==nil)
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// capacity :- maximum number of elemet can fit 
	// yhan automatic resize hora , yhan cap aur len change hogi
	// fmt.Println(cap(nums))


	// to add elements 
	// nums = append(nums,1 )
	// nums = append(nums,2 )
	// nums = append(nums,3 )
	// nums = append(nums,4 )
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// [0 0 1 2 3 4]
    // 10
	// current capacity se double kardeta hai
	// nums = append(nums,1 )
	// nums = append(nums,2 )
	// nums = append(nums,3 )
	// nums = append(nums,4 )
	// nums = append(nums,5 )
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// [0 0 1 2 3 4 1 2 3 4 5]
    // 20
    
	// var nums1 = make([]int,0,5)
	// // iske karan 2 initial 0 bnre hai aur normaly initial size 0 rkhte hain
	// nums1 = append(nums1, 1)
	// nums1 = append(nums1, 2)
	// fmt.Println(nums1)
	// fmt.Println(len(nums1))
	// fmt.Println(cap(nums1))

	// directly 
	// var nums2 = []int{}
	// nums2 = append(nums2, 1)
	// nums2 = append(nums2, 2)

	// can also use indexes
	// var nums2 = make([]int,2,5)
	// nums2[0] = 3
	// nums2[1] = 5
	// fmt.Println(nums2)
	// fmt.Println(len(nums2))
	// fmt.Println(cap(nums2))

	// copy function 
	// var nums = make([]int, 0,5)
	// var nums2 = make([]int,len(nums))
	// nums = append(nums,2)
	
	// fmt.Println(nums,nums2)

	// copy(nums2,nums)
	// fmt.Println(nums,nums2)
	// yhan pe bhi copy ni hoga kyunki des ka len abhi bhi 0 


	// var nums = make([]int, 0,5)
	// nums = append(nums,2)
	// // ab len 1 hogi

	// var nums2 = make([]int,len(nums))
	

	// copy(nums2,nums)
	// fmt.Println(nums,nums2)


	// Slice operator 
	// var nums = []int{1,2,3}
	// fmt.Println(nums[0:2])
	// fmt.Println(nums[:1])
	// fmt.Println(nums[1:])
	// start with zero to till index 2 where index 2 is excluded
	// mention ni krenge to starting se start krega 
	// [from : to ]


	// slices packes
	// var nums1 = []int{1,2,3}
	// var nums2 = []int{1,2,4}


	// compare
	// fmt.Println(slices.Equal(nums1,nums2))
	// increasingly fashion me compare


	// 2d slices
	var nums = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums)






}

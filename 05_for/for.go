package main

import "fmt"

// go ke andar sirf for hai : looping ke liye
// for -> only construct in go for looping
func main(){
    // while loop fashion
	i:=1
	for i<=3{
		fmt.Println(i)
		i = i+1
	}

	// infinite loop
	// for {
	// 	println(1)
	// }
	// ctrl c dbake cancel 


	// Classic for loop 
	for i:=0;i<3;i++ {
		// fmt.Println(i)
		// break :-stop the loop 
		// continue :-skip the current iteration
		if i == 2{
			continue
		}
		fmt.Println(i)
	}

	// 1.22 range 
	for i:=range 3{
		fmt.Println(i)
	}
	// 0 
	// 1
	// 2

}

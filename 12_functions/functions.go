package main

import "fmt"

// creating function
// func name(param type,parm2 type) type konsi value return hori {}
// func add(a int,b int) int {
// 	return a +b

// }

// agar sab ka type same to aese kar skte hain
func add(a ,b int) int {
	return a +b

}

// go ke andar functions multiple values return kar skti hain

func getLanguages()(string ,string ,bool){
	return "golang","js",true
}

// first class decent function 
// func processIt(fn func(a int)int){
// 	fn(1)
// }

func processIt()func(a int)int{
	return func(a int) int {
		return 4
	}
}


func main(){
	// result:=add(6,5)
	// fmt.Println(result)
	// fmt.Println(getLanguages())
	// lang1,lang2,lang3:=getLanguages()
	// to compress _ hum ye use kr rhe taaki error ni de
	// lang1,lang2,_:=getLanguages()
	// fmt.Println(lang1,lang2)

	// fn := func(a int)int{
	// 	return 2
	// }

	// processIt(fn)

	fn:=processIt()
	fmt.Println(fn(6))


}
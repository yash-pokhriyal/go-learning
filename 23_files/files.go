package main

import (
	// "fmt"
	// "bufio"
	// "fmt"
	"fmt"
	"os"
)

func main(){

	// f,err:=os.Open("examples.txt")
	// // go ke nadar error aati to return krte hain
	// if err!=nil{
	// 	// log the error 
	// 	panic(err)
	// }
	// fileInfo,err:=f.Stat()
	// if err!=nil{
	// 	// log the error 
	// 	panic(err)
	// }
	// fmt.Println("file name:",fileInfo.Name())
	// fmt.Println("Is directory:",fileInfo.IsDir())
	// fmt.Println("file size:",fileInfo.Size())
	// fmt.Println("file permission:",fileInfo.Mode())
	// fmt.Println("file modified at:",fileInfo.ModTime())


	// Read file
	// f,err:=os.Open("examples.txt")
	// if err!=nil{
	// 	panic(err)
	// }
	// defer f.Close()

	// // array of bytes:buffer create krenge
	// buf:=make([]byte,12)
	// d,err:=f.Read(buf)
	// if err!=nil{
	// 	panic(err)
	// }
	// for i:=0;i<len(buf);i++{
    //  fmt.Println("data",d,string(buf[i]))
	// }


	// data,err:=os.ReadFile("examples.txt")
	// if err!=nil{
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// hmesha use ni krna chahiye kyunki ye ek baar mai memory me load krti hai

	// read folders
	// dir,err:= os.Open(".")
	// dir,err:= os.Open("../") //rootfolder me jaare
	// if err!=nil{
	// 	panic(err)
	// }
	// defer dir.Close()
	// // fileInfo,err:=dir.ReadDir(2)
	// fileInfo,err:=dir.ReadDir(-1)

	// for _,fi:=range fileInfo{
	// 	fmt.Println(fi.Name())
	// }



	// create a file
	// f,err := os.Create("example2.txt")
	// if err!=nil{
	// 	panic(err)
	// }
	// defer f.Close()

	// f.WriteString("Hi go")
	// f.WriteString("Nice language")

	// task aese karo ki hi go hatt jae aur nice language aae bas

	// bytes := []byte("Hello Golang")

	// f.Write(bytes)


	// read and write to another file (streaming fashion)

	// sourceFile,err:=os.Open("examples.txt")
	// if err!=nil{
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile,err:=os.Create("example2.txt")
	// if err!=nil{
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader:=bufio.NewReader(sourceFile)
	// writer:=bufio.NewWriter(destFile)

	// for{
	// 	b,err:=reader.ReadByte()
	// 	if err!=nil{
	// 		if err.Error()!="EOF"{
    //             panic(err)
	// 		}
	// 		break
			
	// 	}
	// 	e:=writer.WriteByte(b)
	// 	if e!=nil{
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()
	// fmt.Println("Written to new file successfully")


	// deleting file
	
	err:=os.Remove("example2.txt")
	if err!=nil{
		panic(err)
	}
	fmt.Println("File deleted successfully")




}

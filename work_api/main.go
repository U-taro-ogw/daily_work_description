package main

//import "fmt"

func main() {
	//fmt.Println("hello")

	a := App{}
	a.Initialize()
	a.Run(":8080")
}

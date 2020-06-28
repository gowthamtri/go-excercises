package main

import "fmt"

func main() {

	var a float64
	var b int64
	fmt.Println("enter the floating number")
	fmt.Scanf("%f", &a)
	b = int64(a)

	fmt.Printf("%v %T", b, b)
}

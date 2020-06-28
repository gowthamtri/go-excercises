package main

import "fmt"

func main() {
	var x float32;

	num, err := fmt.Scan(&x);

	if num > 0 && err == nil {
			fmt.Println(int32(x));
	}
}
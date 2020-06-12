package main

import (
	"bufio"
  	"os"
	"fmt"
	"encoding/json"
)

type Person struct {
    Name string
    Address string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	firstName := scanner.Text()

	scanner.Scan()
	address := scanner.Text()
	fmt.Println(firstName, address);
	p := Person{firstName, address};
	b, _ := json.Marshal(p);
	fmt.Println(b);
	os.Stdout.Write(b)
}
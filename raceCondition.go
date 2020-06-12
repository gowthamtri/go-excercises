package main

import (
	"fmt"
	"time"
)

/* This function prints the given value 'n' times */

func Print(value string, count int) {
	for i:=0; i < count; i++ {
		fmt.Printf(value + " ")
		time.Sleep(100) // Wait for 100 milliseconds
	}
}

func main() { 
	go Print("A", 5) // Triggers a go routine
	go Print("B", 5) // Triggers a go routine
	go Print("C", 5) // Triggers a go routine
	Print("D", 5) // Normal function call. This gets run on the main thread.

	/*

	Test Runs, 
	1: D B A C D B A B D C A D A B C A C D B C
	2: D B A C C A B D D A C B B C A D C D A B
	3: D A B C B D A C D C B A C A D B A B C D

	On Every run, the order of alphabets change. We can see that 'D' gets printed at start.
	So the function call on the main thread gets first slot for execution. But the other
	go routines are executed at different order in each run.
	*/
}
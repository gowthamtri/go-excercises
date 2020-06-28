package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"strings"
)

type Name struct {
	FirstName string
	LastName string
}

func main() {
	fmt.Println("Enter the file name")
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	fileName := scanner.Text()

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var names []Name

	for fileScanner.Scan() {
		txtLine := fileScanner.Text()
		nameData := strings.Fields(txtLine)
		names = append(names, Name{FirstName: nameData[0], LastName: nameData[1]})
	}

	file.Close()
 
	for _, name := range names {
		fmt.Println(name.FirstName, name.LastName)
	}
}
package main

import (
	"bufio"
  	"os"
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Bird struct {
}


func (self Bird) Speak() {
	fmt.Println("peep")
}

func (self Bird) Eat() {
	fmt.Println("worms")
}

func (self Bird) Move() {
	fmt.Println("fly")
}

type Cow struct {
}


func (self Cow) Speak() {
	fmt.Println("moo")
}

func (self Cow) Eat() {
	fmt.Println("grass")
}

func (self Cow) Move() {
	fmt.Println("walk")
}

type Snake struct {
}


func (self Snake) Speak() {
	fmt.Println("hsss")
}

func (self Snake) Eat() {
	fmt.Println("mice")
}

func (self Snake) Move() {
	fmt.Println("slither")
}



func main() {
	animalMap := map[string]Animal { }

	for {
		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		lower := strings.TrimSpace(strings.ToLower(input));
		data := strings.Fields(lower)

		if len(data) != 3 {
			fmt.Println("Insufficient input")
			continue;
		}

		switch data[0] {
			case "newanimal":
				newAnimal := CreateAnimal(data[2])
				if newAnimal != nil {
					animalMap[data[1]] = newAnimal
					fmt.Println("Animal Created : " + data[1])
				}
			case "query":
				ExecuteAction(animalMap[data[1]], data[2])
		}
	}
}

func CreateAnimal(animalType string) Animal {
	if animalType == "cow" {
		return new(Cow)
	} else if animalType == "snake" {
		return new(Snake)
	} else if animalType == "bird" {
		return new(Bird)
	} else {
		return nil
	}
}

func ExecuteAction(animal Animal, actionName string) {
	if actionName == "eat" {
		animal.Eat()
	} else if actionName == "speak" {
		animal.Speak()
	} else if actionName == "move" {
		animal.Move()
	} else {
		fmt.Println("Wrong input")
	}
}
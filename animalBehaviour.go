package main

import (
	"bufio"
  	"os"
	"fmt"
	"strings"
)

type Animal struct {
	name string
	food string
	locomotion string
	noise string
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func main() {
	animalMap := map[string]Animal {
		"cow": Animal { "cow", "grass", "walk", "moo" },
		"bird": Animal { "bird", "worms", "fly", "peep" },
		"snake": Animal { "snake", "mice", "slither", "hsss" },
	}

	for {
		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		lower := strings.TrimSpace(strings.ToLower(input));
		data := strings.Fields(lower)

		if len(data) != 2 {
			fmt.Println("Wrong input")
			continue;
		}

		if data[1] == "eat" {
			animalMap[data[0]].Eat()
		} else if data[1] == "speak" {
			animalMap[data[0]].Speak()
		} else if data[1] == "move" {
			animalMap[data[0]].Move()
		} else {
			fmt.Println("Wrong input")
		}
	}
}
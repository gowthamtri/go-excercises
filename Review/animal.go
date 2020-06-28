package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// --------- type Animal -------------

type Animal struct {
	food, locomotion, sound string
}

func (self Animal) Eat() {
	fmt.Println(self.food)
}

func (self Animal) Move() {
	fmt.Println(self.locomotion)
}

func (self Animal) Speak() {
	fmt.Println(self.sound)
}

// -----------------------------------

func main() {
	var animal, action string
	var a Animal

	cow := Animal{food: "grass", locomotion: "walk", sound: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", sound: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", sound: "hsss"}

	in := bufio.NewReader(os.Stdin) // I don't use Scan because of white spaces in the input string

	for {
		fmt.Print("> ")
		s, err := in.ReadString('\n')
		if err == nil {

			s = strings.TrimLeftFunc(s, func(r rune) bool { return unicode.IsSpace(r) })  // rempve leading spaces
			s = strings.TrimRightFunc(s, func(r rune) bool { return unicode.IsSpace(r) }) // remove trailing spaces
			request := strings.Split(s, " ")

			animal = strings.ToLower(request[0])
			for i := 1; i < len(request); i++ { // Just in case there is more than one space separating animal and action
				if request[i] != "" {
					action = strings.ToLower(request[i])
					break
				}
			}

			switch animal {
			case "cow":
				a = cow
			case "bird":
				a = bird
			case "snake":
				a = snake
			default:
				continue
			}

			switch action {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				continue
			}
		}
	}
}

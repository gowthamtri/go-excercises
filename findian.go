package main

import (
	"bufio"
  	"os"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
	input := scanner.Text()
	lower := strings.TrimSpace(strings.ToLower(input));

	if strings.HasPrefix(lower, "i") && strings.HasSuffix(lower, "n") && strings.Contains(lower, "a") {
		fmt.Println("Found !");
	} else {
		fmt.Println("Not Found!");
	}
}
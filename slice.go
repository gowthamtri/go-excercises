package main

import (
	"fmt"
	"sort"
)

func main() {
	var x int;
	sli := make([]int, 0, 100);

	for {
		num, err := fmt.Scan(&x);

		if num > 0 && err == nil {
			sli = append(sli, x);
			sort.Ints(sli);
			fmt.Println(sli);
		} else {
			break;
		}
	}
}
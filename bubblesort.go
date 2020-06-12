package main

import (
	"fmt"
)

func main() {
	total := 10;
	var tempInt int;
	sli := make([]int, 0, 10);
	for i := 0; i < total; {
		fmt.Printf("Enter number %d : ", i + 1);
		num, err := fmt.Scan(&tempInt);
		if num > 0 && err == nil {
			sli = append(sli, tempInt);
			i++;
		} else {
			continue;
		}
	}

	fmt.Printf("Input List : ")
	fmt.Println(sli);
	BubbleSort(sli);
	fmt.Printf("Sorted List : ")
	fmt.Println(sli);
}

func BubbleSort(dataSli []int) {
	for i := 0; i < len(dataSli); i++ {
		for j := 0; j < len(dataSli) - 1; j++ {
			if dataSli[j] > dataSli[j + 1] {
				Swap(j, dataSli);
			}
		}
	}
}

func Swap(index int, dataSli []int) {
	temp := dataSli[index];
	dataSli[index] = dataSli[index + 1];
	dataSli[index + 1] = temp;
}
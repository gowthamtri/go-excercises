package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(j int, arr []int) {
	temp := arr[j]
	arr[j] = arr[j+1]
	arr[j+1] = temp
}

//BubbleSort Takes Slice As input and Sorts Array
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(j, arr)
			}
		}
	}
}

func main() {
	arr := make([]int, 10)
	i := 0
	reader := bufio.NewReader(os.Stdin)
	for i < 10 {
		fmt.Printf("Enter %d number : ", i+1)
		num, _ := reader.ReadString('\n')
		val, _ := strconv.Atoi(strings.TrimSpace(num))
		arr[i] = val
		i++
	}
	fmt.Println("Before Sorting : ")
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println("After Sorting : ")
	fmt.Println(arr)
}

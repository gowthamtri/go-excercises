package main

import (
	"bufio"
  	"os"
	"fmt"
	"strings"
	"sync"
	"math"
	"strconv"
)

func main() {
	fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		fields := strings.Fields(strings.TrimSpace(input));

		data := make([]int, 0, len(fields))
		for _, v := range(fields) {
			if len(strings.TrimSpace(string(v))) > 0 {
				val, _ := strconv.Atoi(v)
				data = append(data, val)
			}
		}

		fmt.Println("DATA : ", data)

		n := int(math.Max(math.Ceil(float64(len(fields)) / float64(4)), 1))

		var wg sync.WaitGroup
		for i:=0; i < 4;i++ {
			from := int(math.Min(float64(i * n), float64(len(data))))
			to := int(math.Min(float64((i+1) * n), float64(len(data))))

			wg.Add(1)

			go sortArr(data[from:to], &wg)
		}

		wg.Wait()

		sortArr(data, nil)

		fmt.Println("SORTED : ", data)
}

func sortArr(dataSli []int, wg *sync.WaitGroup) {
	fmt.Println("Sorting :", dataSli)
	for i := 0; i < len(dataSli); i++ {
		for j := 0; j < len(dataSli) - 1; j++ {
			if dataSli[j] > dataSli[j + 1] {
				Swap(j, dataSli);
			}
		}
	}

	if wg != nil {
		wg.Done()
	}
}

func Swap(index int, dataSli []int) {
	temp := dataSli[index];
	dataSli[index] = dataSli[index + 1];
	dataSli[index + 1] = temp;
}
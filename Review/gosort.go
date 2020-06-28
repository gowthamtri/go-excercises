package main
import(
	"fmt"
	"sync"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the values in the format 1 2 3 4 5")
	a,_ := reader.ReadString('\n')
	str := strings.Split(strings.TrimSpace(a)," ")

	array := make([]int,0,len(str))

	for _, s := range str{
		n, _ := strconv.Atoi(s)
		array = append(array, n)
	}

	arr_len := len(array)
	part := int(arr_len/4)
	fmt.Println(part)

	var wg sync.WaitGroup

	for i:=0;i<4;i++{
		start := int(i*(part))
		end := int((i+1)*part)
		if arr_len%2 != 0 && i == 3{
			end++;
		}

		wg.Add(1)

		go func(arr []int){
			fmt.Println("Sorting :",arr)
			sort.Ints(arr)
			wg.Done()
		}(array[start:end])
		
	}
	wg.Wait()
	sort.Ints(array)
	fmt.Println(array)
}
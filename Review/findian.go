package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
  "strconv"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    line := scanner.Text()
    
    
    p:=strconv.FormatBool(strings.HasPrefix(line, "i"))
    q:=strconv.FormatBool(strings.HasSuffix(line, "n"))
    r:=strconv.FormatBool(strings.Contains(line, "a"))
    
    if p=="true" && q=="true" && r=="true" {fmt.Printf("Found!") 
        
    }else {
        fmt.Printf("Not Found!")} 
    
    
    
   
}
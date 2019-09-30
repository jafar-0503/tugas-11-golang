package main

import "fmt"
import "strconv"

func main(){
  var str1 = "50"
  var str2 = 25
  var num, err = strconv.Atoi(str1)

  if err == nil{
    fmt.Println(num)
    fmt.Println(num + str2)
    fmt.Println(num - str2)
    fmt.Println(num / str2)
    fmt.Println(num * str2)
  }
}

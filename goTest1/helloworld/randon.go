package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	  num1 :=rand.Intn(100)
	  fmt.Println(num1)
	  rand.Seed(1)
	  num2 :=rand.Intn(100)
	  fmt.Println(num2)
	  t1 :=time.Now()
	  fmt.Println(t1)
	  fmt.Printf("%T\n",t1)
	  hour :=t1.Hour()
	  fmt.Println(hour)
	  randon()
	  num3 :=rand.Intn(46)+3
	  fmt.Println(num3)
}
func randon(){
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++{
		fmt.Println(rand.Intn(100))
	}
}

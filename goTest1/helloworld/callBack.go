package main

import "fmt"

func main(){
	fmt.Println(oper(1,2,add))
}
//加法
func add(a,b int)int{
	return a+b
}
func oper(a,b int,fun func(int,int)int)int{
	return fun(a,b)
}

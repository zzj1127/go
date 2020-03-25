package main

import "fmt"

func main(){
	//fmt.println的内容不可以为空
	//定义一个函数变量
	var res func(int,int)
	fmt.Println(res)
	//给函数赋值
	res1 := fun2
	fmt.Println(res)

	fmt.Println(res1(1,2))
}
func fun1(a,b int){
	fmt.Println(a,b)
}
func fun2(a,b int)(int){
	return a+b
}

package main

import "fmt"

func main(){
	//defer延迟执行
	//defer 栈结构，先进后出，
	//defer 延迟执行并不延迟取值
	//defer fun1("111")
	//fmt.Println("12345")
	//defer fun1("1111")
	//fmt.Println("王二狗")
	a := 2
	defer fun2(a)
	 a++
	fmt.Println(a)
}
func fun2(a int){
	fmt.Println(a)
}
func fun1(s string){
	fmt.Println(s)
}

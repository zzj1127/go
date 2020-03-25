package main

import "fmt"

func main(){
	//go语言支持函数编程
	fun1()
	fun2 := fun1
	fun2()
	//匿名函数
	//匿名函数只能使用一次无法调用
	//可以把匿名函数赋值给变量，可以反复调用
	fun3 :=func(){
		fmt.Println("222")
	}
	fun3()

	//带参数
	func(a,b int){
		fmt.Println(a,b)
	}(1,2)
	//带参数返回值
	func(a,b int)(int){
		return a+b
	}(1,2)
}
func fun1(){
	fmt.Println("11111")
}

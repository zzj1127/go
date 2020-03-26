package main

import "fmt"

func main(){
	var a func()
	a = potint
	a()
	fmt.Println(fuc2())
	arr2 :=fuc3()
	fmt.Println(&arr2)
}
func potint(){
	fmt.Println("111")
}
func fuc2()[4]int{
	return [4]int{1,2,3,4}
}
//指针函数
func fuc3()*[4]int{
	arr :=[4]int{2,3,4,5}
	fmt.Println(&arr)
	return &arr
}

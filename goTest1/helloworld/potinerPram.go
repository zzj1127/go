package main

import "fmt"

func  main(){
	//函数传值是通过把值赋给参数，意味着调用的参数和本身函数的参数不是同一个
	//函数传指针是通过把地址传给参数，以为这修改哪里都可以修参数的值
	//
	arr1 :=[4]int{1,2,3,4}
	fmt.Println(arr1)
	fun3(arr1)
	fmt.Println(arr1)
	fmt.Println("----------------------- ")
	fmt.Println(arr1)
	fun4(&arr1)
	fmt.Println(arr1)

}
func fun4(p2 *[4]int){
	fmt.Println(p2)
	p2[0] = 200
	fmt.Println(p2)
}
func fun3(arr [4]int){
	fmt.Println(arr)
	arr[0] =100
	fmt.Println(arr)
}

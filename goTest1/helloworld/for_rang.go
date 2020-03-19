package main

import "fmt"

func main(){
	arr1 :=[5]int{1,2,3,4,5}
	//1
	for i:=1;i<len(arr1);i++{
		fmt.Println(arr1[i])
	}
	fmt.Print("----------\n")
	//2 取得下标和数值
	for index,value :=range arr1{
		fmt.Println(index,value)
	}
	fmt.Print("----------\n")
	for _,value1 :=range arr1{
		fmt.Println(value1)
	}
}

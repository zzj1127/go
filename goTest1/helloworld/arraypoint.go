package main

import "fmt"

func main(){
	//数组指针
	arr1 :=[4]int{1,2,3,4}
	fmt.Println(arr1)
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)
	fmt.Printf("%p\n",p1)//数组地址
	fmt.Printf("%p\n",&p1)//数组指针地址

	(*p1)[0] = 100
	fmt.Println(arr1)
	p1[0] =200//简化写法
	fmt.Println(arr1)
	//指针数组
	a :=1
	b :=2
	c :=3
	d :=4
	arr2 :=[4]int{a,b,c,d}
	arr3 :=[4]*int{&a,&b,&c,&d}
	fmt.Println(arr2)
	fmt.Println(arr3)


}

package main

import "fmt"

func main(){
	a :=10
	fmt.Println(a)
	fmt.Println(&a)
	//创建指针
	var p1 *int
	fmt.Println(p1)
	p1 = &a
	//p1地址
	fmt.Println(&p1)
	fmt.Println(*p1)
	//指针的指针
	var p2 **int
	fmt.Println(p2)
	p2 =&p1

	fmt.Println(p2)

	fmt.Println(*p2)
	fmt.Println(**p2)
}

package main

import "fmt"

func main(){
	//const1()
	const2()
}
func const2(){
	//iota:特殊的常量，可以被编译器自动修的常量
	const a,b,c = iota,iota,iota
	const (
		q = iota
		w = iota
		e = iota
	)
	fmt.Print(a,b,c)
	fmt.Print(q,w,e)
}
func const1(){
//	固定的数值
	const PATH string  = "www.baid.com"
	const PI   = 3.14
	//fmt.Print(PATH,PI)
	const(
		a int = 100
		b
		c string = "111"
		d
	)
	fmt.Printf("a:%T,%d\n",a,a)
	fmt.Printf("b:%T,%d\n",b,b)
	fmt.Printf("c:%T,%s\n",c,c)
	fmt.Printf("d:%T,%s\n",d,d)

}
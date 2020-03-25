package main

import "fmt"

func main(){
	//连续开辟俩个内存
	//指向的地址是里面函数的地址
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment()())
	fmt.Println(increment()())
	//rest1开辟内存，调用俩次rest1函数
	res1 :=increment()
	fmt.Println(res1())
	fmt.Println(res1())
}
//外层函数
func increment()func()int{
	i := 0
	//内层函数
	return  func ()int{
		i++
		return i
	}

}

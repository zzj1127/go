package main

import "fmt"
//全局变量（不能使用简短定义）
func main(){
	var2()
}
func var2(){
	var num int
	num = 100
	fmt.Printf("num的数值：%d,地址是：%p\n",num,&num)

	var name string
	name = "111"
	fmt.Printf("name的值：%s,地址是%p\n",name,&name)

	num,name,sex := 99,"222","男"
	fmt.Printf("num的数值：%d,地址是：%p\n",num,&num)
	fmt.Printf("nameg的值：%s,地址是%p\n",name,&name)
	fmt.Printf("sex的值：%s,地址是%p\n",sex,&sex)
}
func var1(){
	var num1 int
	num1 = 30
	fmt.Printf("num1的值：%d\n",num1)
	var num2 int = 30
	fmt.Printf("num2的值：%d\n",num2)
	var name="111"
	fmt.Printf("数据类型：%T\n",name)
	num3 :=  "111"
	fmt.Printf("num3的值：%s\n",num3)
	//多个变量申明
	var a,b,c int
	a = 1
	b = 2
	c = 3
	fmt.Print(a,b,c)

	var  m,n int = 200,300
	fmt.Print(m,n)

	var w,s,d=1,3.12,"abc"
	fmt.Print(w,s,d)
	var (
		studentName = "李小龙"
		age = 18
		sex = "男"
	)
	fmt.Printf("名字：%s,年龄：%d岁，性别：%s\n",studentName,age,sex)
}
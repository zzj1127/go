package main

import "fmt"

func main(){
	//创建
	var map1 map[int]string//map没有初始化 nil
	var map2  = make(map[int]string)//make创建自带初始化
	var map3 = map[int]string{1:"go",2:"java"}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	if map1 == nil{
		map1 =make(map[int]string)
		fmt.Println("初始化")
	}
	//存储
	map1[1]="hello"
	map1[2]="world"
	//读取
	fmt.Println(map1)
	fmt.Println(map1[2])
	value,ok :=map1[2]
	fmt.Println(value)
	fmt.Println(ok)
	//修改
	map1[2] = "hello world"
	fmt.Println(map1)
	//删除
	delete(map1,1)
	fmt.Println(map1)
	//获取长度
	fmt.Println(len(map1))
	//map遍历
	map2[1] = "111"
	map2[2] = "222" 
	map2[3] = "333"
	map2[4] = "444"

	for k , v := range map2{
		fmt.Println(k,v)
	}
	//sort包进行排序、、

}

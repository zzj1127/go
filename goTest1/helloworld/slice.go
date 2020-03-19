package main

import "fmt"

func main(){
	/*
	切片与数组相似，存储相同数据类型的数据，但是切片可以自己扩容
	对于数组遍历方法同样使用于切片
	
	 */
	arr :=[4]int{1,2,3,4}
	fmt.Println(arr)
	//切片
	var s1 []int
	fmt.Println(s1)
	//make(类型，长度，容量)
	s3 :=make([]int,3,8)
	fmt.Println(s3)
	fmt.Printf("容量：%d，长度：%d\n",cap(s3),len(s3))
	for i:=0;i<len(s3);i++{
		s3[i]=i+1
		fmt.Println(s3[i])
	}
	s3 =append(s3,1,2)
	fmt.Println(s3)
}

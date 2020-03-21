package main

import "fmt"

func main(){
	s1 :=[]int{1,2,3,4}
	//深copy
	s2 :=make([]int,0)
	for i:=0;i<len(s1);i++{
		s2 = append(s2,s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("=============================")
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("=============================")
	s3 :=[]int{7,8,9}
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("=============================")
	copy(s2,s3)//将s3中的元素拷贝到s2中
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("=============================")
	copy(s2,s1[:2])
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("=============================")
}

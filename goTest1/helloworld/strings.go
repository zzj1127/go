package main

import (
	"fmt"
	"strings"
)

func main(){
	s1 :="helloWorld"
	//查看s1中是否存在相应的字符
	fmt.Println(strings.Contains(s1,"a"))
	//是否s1中的任意一个字符
	fmt.Println(strings.ContainsAny(s1,"adcd"))
	//统计相应字符在s1中的次数
	fmt.Println(strings.Count(s1,"h"))
	s2 := "2020年.java"
	//判断字段开头
	fmt.Println(strings.HasPrefix(s2,"2020"))
	//判断字段结尾
	fmt.Println(strings.HasSuffix(s2,".java"))
	//返回在string中的位置，不存在返回-1
	fmt.Println(strings.Index(s1,"h"))
	//返回在string中的位置的最小值,不存在返回-1
	fmt.Println(strings.IndexAny(s1,"abdh"))
	//返回在string中的最后的位置,不存在返回-1
	fmt.Println(strings.LastIndex(s1,"o"))
	//string数组之间的拼接
	ss1 :=[]string{"11","222"}
	 fmt.Println(strings.Join(ss1,"-"))
	//切割
	s3 :="123,234,234,435"
	ss2 :=strings.Split(s3,",")
	fmt.Println(ss2)
	fmt.Printf("%T\n",ss2)
	//拼接
	fmt.Println(strings.Repeat("hello",5))
	//替换 strings.Replace中替换所有的为-1
	fmt.Println(strings.Replace(s1,"o","0",-1))
	//字符的大小写转换
	s4 :="helloWorld***123"
	fmt.Println(strings.ToLower(s4))
	fmt.Println(strings.ToUpper(s4))
	//截取字串substring
	fmt.Println(s4)
	fmt.Println(s4[0:5])
	fmt.Println(s4[5:])
}

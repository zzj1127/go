package main

import (
	"fmt"
	."strconv"
)

func main(){
	fmt.Println("aa"+FormatInt(100,10))
	b1,err := ParseBool("true")
	fmt.Println(b1,err)
	fmt.Println(ParseInt("100",2,64))
	//itoa()//转为string,atoi()转为int类型
	i1,err := Atoi("-31")
	s3 :=Itoa(-42)
	fmt.Printf("%T,%T",i1,s3)
	fmt.Println(i1,s3,err)
	i2,_ :=Atoi("10")
	fmt.Println(getNum(i2))
	s1 :=[]int{1,2,3,4}
	randonParm(s1 ...)
}
func getNum(n int)(m int){
	sum :=0
	for i:=1;i<=n;i++{
		sum += i
	}
	return sum
}
func randonParm(num ... int){
	sum := 0
	for i := 0;i<len(num);i++{
		sum +=num[i]
	}
	fmt.Println(sum)
}

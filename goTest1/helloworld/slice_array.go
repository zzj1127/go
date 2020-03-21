package main

import "fmt"
//s :=a[开始下标：结束下标]
func main(){
	a := [10]int{1,2,3,4,5,6,7,8,9,10}
	s1 :=a[:5]  //[1 2 3 4 5] len 5 cap10
	s2 :=a[3:8]//[4 5 6 7 8]len 5 cap7
	s3 :=a[5:]//[6 7 8 9 10]len 5 cap5
	s4 :=a[:]//[1 2 3 4 5 6 7 8 9 10]len10 cap10
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)


}

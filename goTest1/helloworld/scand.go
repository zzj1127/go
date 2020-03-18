package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//io读取
	reader :=bufio.NewReader(os.Stdin)
	s1,_ := reader.ReadString('\n')
	fmt.Println(s1)
	//scanf
	var x int
	var y float64
	fmt.Scanf("%d,%f",&x,&y)
	fmt.Println(x,y)
	//scanln
	var a string
	var b float64
	fmt.Scanln(&a,&b)
	fmt.Println(a,b)

}

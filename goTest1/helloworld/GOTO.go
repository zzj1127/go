package main

import "fmt"

func main(){
	var a = 10
	togo(a)
}
func togo(a int){

Loop://标记
	for a < 20 {
		if a == 16 {
			a += 1
			goto Loop
		}
		fmt.Println(a)
		a++
	}

}
func togo2(){

}

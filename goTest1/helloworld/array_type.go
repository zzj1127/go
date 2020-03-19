package main

import "fmt"

func main(){
	arr1 :=[3]int{1,2,3}
	arr2 :=[3]float64{1.2,2.3,3.4}
	arr3 :=[3]string{"q","w","e"}
	arr4 :=[2]int{1,2}
	fmt.Printf("%T\n",arr1)
	fmt.Printf("%T\n",arr2)
	fmt.Printf("%T\n",arr3)
	fmt.Printf("%T\n",arr4)
}

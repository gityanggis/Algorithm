package main

import (
	"fmt"
)


func main(){
	var a=[]int{3,9,1,4,5,7,2}

	//冒泡排序
	inserSort(a)

	for i:=0;i<len(a) ;i++  {
		fmt.Print(a[i]," ")
	}
}





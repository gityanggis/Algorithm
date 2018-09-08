package main

import "fmt"

//选择排序 （选出最大的放到前面）
func SelectSort(a []int){
	len:=len(a)
	for i:=0;i<len ;i++  {
		max:=i
		for j:=i+1;j<len ;j++  {
			if a[j]>a[i] {
				max=j
			}
		}
		a[i],a[max]=a[max],a[i]

		fmt.Print(a[i],",")
	}
}
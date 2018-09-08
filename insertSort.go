package main

//插入排序
//每次都将当前元素插入到左侧已经排序的数组中，使得插入之后左侧数组依然有序。
func inserSort(a []int){
	len:=len(a)
	for i := 1; i < len; i++ {
		for j:=i;j >0;j--{
			if a[j-1]>a[j] {
				a[j-1],a[j]=a[j],a[j-1]
			}
		}
	}
}

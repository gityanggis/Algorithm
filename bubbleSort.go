package main


//冒泡排序，不断交换找到最小的数据到最后
func BubbleSort(a []int){
	len:=len(a)
	flag:=true
	for i := len-1; i >0&&flag;i--  {
		flag=false
		for j := 0; j < i; j++ {
			if a[j+1] >a[j]{
				flag=true
				a[j+1],a[j]=a[j],a[j+1]
			}
		}
	}
}

package main



//快速排序通过一个切分元素将数组分为两个子数组，左子数组小于等于切分元素，右子数组大于等于切分元素，
//将这两个子数组排序也就将整个数组排序了。
func QuickSort(a []int){
	len:=len(a)
	quickSort(a,0,len-1)
}

func sort(a []int,l int,r int) int{
	t:=a[l]
	j:=l
	for i:=l+1;i<=r ;i++  {
		if t>a[i] {
			a[i],a[j+1]=a[j+1],a[i]
			j++
		}
	}

	a[l],a[j]=a[j],a[l]
	return j
}

func quickSort(a []int,l int,r int){
	if l>=r {
		return
	}
	p:=sort(a,l,r)
	quickSort(a,l,p-1)
	quickSort(a,p+1,r)
}


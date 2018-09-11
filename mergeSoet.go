package main



func MergeSort(a []int){
	len:=len(a)
	if len!=0 {
		mergeSort(a,0,len-1)
	}
}


func mergeSort(a []int,f int,l int){
	if f>=l {
		return
	}

	mid:=f+(l-f)/2

	mergeSort(a,f,mid)
	mergeSort(a,mid+1,l)
	merge(a,f,mid,l)
}

func merge(a []int,f int,mid int,l int){
	i:=f
	j:=mid+1

	//b :=a[f:]
	b:=make([]int,l-f+1)
	nn:=f
	for mm:=0;mm<l-f+1;mm++{
		b[mm]=a[nn]
		nn++
	}

	for k:=f;k<=l ;k++  {
		if i> mid{
			a[k]=b[j-f]
			j++
		}else if j>l {
			a[k]=b[i-f]
			i++
		}else if b[i-f]<b[j-f] {
			a[k]=b[i-f]
			i++
		}else {
			a[k]=b[j-f]
			j++
		}
	}


}
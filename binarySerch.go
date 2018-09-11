package main



func BinarySerch(a []int,t int)int{
	len:=len(a)

	l:=0
	r:=len-1
	for l<=r {
		mid:=(r-l)/2

		if a[mid]==t {
			return mid
		}
		if a[mid]<t {
			r=mid-1
		}else {
			t=mid+1
		}
	}
	return -1
}


package selection

func Sort(arr []int) {
	//i int
	for i:=0;i<len(arr)-1;i++ {
		k:=i
		for j:=i+1;j<len(arr);j++ {
			if arr[j]<arr[k] {
				k=j
			}
		}
		if k!=i {
			arr[k],arr[i]=arr[i],arr[k]
		}
	}
}

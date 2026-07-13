package main 

func Bubble(arry [420]int) [420]int{
	for i := 0; i < len(arry); i++ {
		for j := 0; j < len(arry)-1; j++ {

			// increasing order 
			if arry[j]>arry[j+1] {
				temp :=arry[j+1]
				arry[j+1] = arry[j]
				arry[j] = temp
			}
		}
	}

	return arry
}
package algo

import "sort"
func Bubble(arry [420]int) [420]int {
	for i := 0; i < len(arry); i++ {
		for j := 0; j < len(arry)-1; j++ {

			// increasing order
			if arry[j] > arry[j+1] {
				temp := arry[j+1]
				arry[j+1] = arry[j]
				arry[j] = temp
			}
		}
	}

	return arry
}

func Selection(arry [420]int) [420]int {
	n:=len(arry)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i+1; j < n; j++ {
			if arry[j] < arry[min] {
				min = j
			}
		}
		//swap
		if min != arry[i]{
			arry[i],arry[min]=arry[min],arry[i]
		}

	}
	return arry
}
//Insertion

func Insertion(arry [420]int) [420]int {
 n:=len(arry)
 for i:=0; i<n; i++{
	current:= arry[i]
	previous:=i-1
	for previous >=0 && current < arry[previous]{
		arry[previous+1]= arry[previous]
		previous -- 
	}
	arry[previous+1]= current

 }
	return arry	
}

func Defult(arry [420]int)[420]int{
	sort.Ints(arry[:])
	return arry 
}

package main

import (
	"fmt"
)

func recursion(x int){
	
	//base case
	if x ==1 {
		fmt.Println(x)
		return
	}
	recursion(x-1)
	fmt.Println(x)
	
	
}

func factorial(n int) int{
	if n==0 {
		return 1;
	}
	fn := n* factorial(n-1)
	return fn
}

	func rmDuplicate(arry [26]bool,word string,idx int,newword string){
		if idx == len(word){
			fmt.Println(newword)
			return;	
		}
		//again call
		if arry[word[idx] - 'a'] == true {
			rmDuplicate(arry,word,idx+1,newword)
		} else{
			newword+=string(word[idx])
			arry[word[idx]-'a'] = true
			rmDuplicate(arry,word,idx+1,newword)
		}
	}

	func main()  {
		// recursion(10)
		arr1 := [26]bool{}
		// fmt.Println(factorial(20))
		rmDuplicate(arr1,"appnnacollege",0,"")
	}




package main

import "fmt"







func fastExpow(x int)int{
	var ans int =1
	var n int = x
	for n>0{
		if n & 1 == 1 {
			ans = ans * x
			x=x*x
		}
		n=n>>1
	}
	return ans
}
func swap(a,b int)(int,int){
	// return b,a 
	a=a^b
	fmt.Println(a)
	b=a^b
	fmt.Println(b)
	a=a^b
	fmt.Println(a)
	return a,b
}

func addOneBit(x int) int{
	return -(^x)
}
func upperToLower(x string)string{
	result :=make([]byte,len(x))
	for  i := 0; i < len(x);i++ {
		result[i]=x[i] | ' '
	}
	return string(result)
}

func main()  {
	// Question 1 :What is the value of x^x for any valueof x?
	fmt.Println("value of x^x is : ",fastExpow(10))
	fmt.Println(swap(5,10))
	fmt.Println(addOneBit(-1))
	fmt.Println(upperToLower("ASFJSJD"))
}

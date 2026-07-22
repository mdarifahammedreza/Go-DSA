package main

import "fmt"

// func main()  {
// 	path:="WNEENESENNN"
// 	x:=0;
// 	y:=0;
// 	for _,ch:=range path {
// 		if 87 == ch {
// 			y--;
// 		}else if 78 == ch {
// 			x++;
// 		} else if 69 == ch {
// 			y++
// 		}else {x--}
// 	}
// 	ans := math.Sqrt((math.Pow(float64(x),2))+(math.Pow(float64(y),2)))
// 	fmt.Println(ans)
// }

//Binary Number system

// func main(){
// 	fmt.Println("number system")
// 	// and operator
// 	x := 33
// 	y := 22
// 	fmt.Printf("%b <-x y-> %b \n",x,y)
// 	fmt.Printf("AND operator %b \n",x & y)
// 	//or operator
// 	fmt.Printf("or operator %b \n",x | y)
// 	//xor
// 	fmt.Printf("xor operator %b \n",x ^ y)
// 	//one's compilement ~
// 	//binary left shift <<
// 	//binary Right shift >>

// 	var number int = 2
// 	var bitMask int = 1
// 	if number & bitMask==0 {
// 		fmt.Println("Even number")
// 	} else {fmt.Println("Odd number")}

// 	var numberbit int = 6
// 	var bit int = 3
// 	var bitMask1 int = 1 << bit
// 	if numberbit & bitMask1 == 0 {
// 		fmt.Println(0)
// 	} else {fmt.Println(1)}

// }

// func main(){
// 	number:=10
// 	bit:=2
// 	bitMask:=1<<bit

// 	bit = number | bitMask
// 	fmt.Printf("%b\n",bit)
// 	fmt.Println(bit)

// }


func main(){
	// number:=10
	// bit:=1
	// bitMask:=^1<<bit

	// bit = number & bitMask
	// fmt.Printf("%b\n",bit)
	// fmt.Println(bit)
	// fmt.Println(clearRangeOfBits(10,2,4))
	// fmt.Println(powerOfTwo(8))
	// fmt.Println(setBits(15))
	fmt.Println(fastExpotential(5,5))
}

func clearRangeOfBits(n,i,j int)int{
	//genrate 111100000
	a:= (^0)<< j+1
	b:=(1<<i)-1 //00000001111
	bitMask := a | b


return n & bitMask;	
}

func powerOfTwo(n int)bool{
	if n & (n-1) == 0{
		return true
	}
	return false

}

func setBits(n int)int{
	count:=0;
	for n>0{
	if n & 1 != 0 {
		count++
	}
	n = n>>1;
}
return count
}

func fastExpotential(n int,pow int)int{
	ans:=1
	for pow >0 {
		if pow & 1 == 1 {
			ans= ans*n
			n=n*n
		}
		pow=pow>>1
	}


	return ans;
}
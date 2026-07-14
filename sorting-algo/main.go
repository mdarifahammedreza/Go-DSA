// bubble sort practicing
package main

import (
	"Go-DSA/algo"
	"Go-DSA/data"
	"fmt"
	"time"
)
func main(){
	arry:= data.Arry
	// buble
	start := time.Now()
	fmt.Println("bubble sort -->",algo.Bubble(arry))
	elapsed := time.Since(start)
	fmt.Println("Bubble Sort Runtime:", elapsed)
	// Selection
	start = time.Now()
	fmt.Println("Selection sort -->",algo.Selection(arry))
	elapsed = time.Since(start)
	fmt.Println("Selection Sort Runtime:", elapsed)

	//insertion
	start = time.Now()
	fmt.Println("Insertion sort-->",algo.Insertion(arry))
	elapsed  = time.Since(start)
	fmt.Println("Insertion sort runtime:",elapsed)

	//Defult
	start = time.Now()
	fmt.Println("defult sort-->",algo.Defult(arry))
	elapsed  = time.Since(start)
	fmt.Println("defult sort runtime:",elapsed)
	// 	start = time.Now()
	// elapsed  = time.Since(start)
	// fmt.Println("Insertion sort runtime:",elapsed)

}



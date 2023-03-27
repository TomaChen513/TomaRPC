package main

import (
	"fmt"
	// "strconv"
	"strings"
)

func main() {
	// var str string
	// var num int
	// var float float64
	// fmt.Scanf("%d",&num)
	// fmt.Scanf("%s",&str)
	// fmt.Scanf("%f",&float)


	// n
	// 123 213 213
	
	// fmt.Println(str)
	// fmt.Println(num)
	// fmt.Println(float)
	inputDemo()
	
}

func inputDemo(){
	// EgTn61O2Hi
	var n int
	fmt.Scan(&n)
	var line string
	fmt.Scan(&line)
	res:=strings.Split(line," ")
	fmt.Println(res)
}
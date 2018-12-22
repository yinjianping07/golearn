package main

import "fmt"

func main(){
	fmt.Println("hello go!")
	var a []int = []int{1,2,1,5}
	key := 1
	a = test(a,key)
	fmt.Println(a)
}

func test(a []int,key int)([]int){
	i := 0
	for i<len(a) {
		if a[i] != key {
			i++
		} else {
			for j := i;j<len(a)-1;j++ {
				a[j] = a[j+1]
			}
			a = a[:len(a)-1]
		}
	}
	return a
}

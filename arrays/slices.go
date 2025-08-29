package arrays

import "fmt"

func main(){
	var slice = []int{1, 2, 3, 4}
	var myslice = []int{}
	slice = append(slice, 10)
	// fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println("myslice", myslice)
	printArray(slice)
}

func printArray(list []int){
	for _, item := range list {
		fmt.Printf("%d\t", item)
	}
}
package main

import "fmt"

func catch(arr[]int) ([]int) {
	fmt.Println(arr)
	var j = 0
	for i := 0; i < len(arr);i++ {
		if(arr[i] != 0 && arr[j] == 0){
			arr[i], arr[j] = arr[j], arr[i]
		}
		if(arr[j] != 0){
			j++
		}
	}
	fmt.Println(arr)
	return arr
}

func main(){
	array := []int{1, 0, 3, 0, 4, 5, 6, 0, 6, 0, 4, 0, 4}
	catch(array)
}


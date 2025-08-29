package main

import "fmt"

func main(){
	fmt.Println(removeEvenNumbersFromArry([]int{1, 2, 3, 4, 5}))
}

// retorno nomeado em go
func removeEvenNumbersFromArry(arr []int) (oddNumbers []int) {
	for i := 0;i < len(arr); i++ {
		if arr[i]%2 != 0 {
			oddNumbers = append(oddNumbers, arr[i])
		}
	}  
	return
}
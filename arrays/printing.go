

// AULA 15 PRINTANDO ELEMENTOS DE UM ARRAY

import "fmt"

func main() {

	var arr = []int{21, 3, 432, 432, 5543, 5643}
	printArray(arr)
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d \t", arr[i])
	}

	for _, value := range arr {
		fmt.Printf("%d \t", value)
	}
}
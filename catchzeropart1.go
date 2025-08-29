import "fmt"

func catchZerosAndSendToEnd(arr[]int) ([]int) {
	var zeros = []int{}
	var newArr = []int{}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		if arr[i] == 0 {
			zeros = append(zeros, arr[i])
		}else {
			newArr = append(newArr, arr[i])
		}
		
	}
	fmt.Println("Zeros", zeros) 
	newArr = append(newArr, zeros...)
	fmt.Println("newArr", newArr)
	return newArr
}

func catch(){
	var array = []int{0, 1, 2, 0, 3, 4, 0, 5, 0, 6, 7}
	catchZerosAndSendToEnd(array)
}


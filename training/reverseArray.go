
import "fmt"

func reverseArray(arr[]int) ([]int){
	// atribui 0 a variavel que da inicio
	start := 0
	// atribui o tamanho da array a variavel end
	end := len(arr) - 1

	// while do golang enquanto start for menor que end
	for start < end {
		// atribui o inicio a variavel fim e o fim a variavel inicio
		arr[start], arr[end] = arr[end], arr[start]
		// aumenta o start a cada iteracao
		start++
		// diminui o end a cada iteracao
		end-- 
	}
	fmt.Println(arr)
	return arr
}

func main(){
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	reverseArray(array)
}
package main 

import "fmt"

func findOdd(seq []int) int {
	count := make(map[int]int)
	for i := 0; i < len(seq); i++ {
		value, exists := count[seq[i]]
		if exists {
			count[seq[i]] = value + 1	
		} else {
			count[seq[i]] = 1
		}
	}	
	var result int
	for key, value := range count {
		if value % 2 != 0 {
			result = key
		} 
	}

return result
}

func main() {
	fmt.Println(findOdd([]int{7}))
}

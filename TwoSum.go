package main

func TwoSum(numbers []int, target int) [2]int {
	for i, v := range numbers {
		for j, vj := range numbers {
			if i != j && (v+vj) == target {
				return [2]int{i, j}
			}
		}

	}
	return [2]int{}
}

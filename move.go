package main

func Move(input []int) []int {
	w, r := len(input)-1, len(input)-1
	for r >= 0 {
		if input[r] == 0 {
			r--
		} else {
			input[r], input[w] = input[w], input[r]
		}
	}
	return input
}

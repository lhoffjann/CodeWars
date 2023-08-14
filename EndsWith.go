package main

import "math"
func solution(str, ending string) bool {
	end_of_str := str[int(math.Max(float64(len(str)-len(ending)),0)):]
	return end_of_str == ending
}


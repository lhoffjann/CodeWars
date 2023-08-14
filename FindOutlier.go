package main

func FindOutlier(input []int) int  {
	var result bool
	q := []bool{}
	for _, v := range input {
		q = append(q, v&1==0)	   
	}
	for _, v := range q {
		result ^= v	
	}	
	for k,v := range q{
		if v == result {
			return input[k]
		}
	}
	return -1
}

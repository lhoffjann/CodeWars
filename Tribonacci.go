package main

func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0 {
		return make([]float64, 0)
	} else if n < 3 {
		return signature[:n]
	}
	
	float_slice := signature[:]
	for i := 0; i < n-3; i++ {
		float_slice = append(float_slice, sumSlice(float_slice[len(float_slice)-3:]))	
	}
return float_slice		
	
}

func sumSlice(nums []float64) float64 {
    sum := 0.0 
   for _, num := range nums {
        sum += num
    }
    return sum
}



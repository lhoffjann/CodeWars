package main

func binarySearch2(a []int, key int) int {
	lo, hi := 0, len(a)-1
	var mid int
	for lo <= hi {
		mid = (hi-lo)/2 + lo
		if a[mid] == key {
			return mid
		}
		if a[mid] > key {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

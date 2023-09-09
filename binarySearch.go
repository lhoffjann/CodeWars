package main

func bsearch(a []int, key int) int {
	lo, hi := 0, len(a)-1
	for lo <= hi {
		mid := lo + ((hi - lo) / 2)
		if a[mid] == key {
			return mid
		}
		if a[mid] < key {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

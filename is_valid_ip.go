package main

import ("strings"
	"strconv")

func Is_valid_ip(ip string) bool {
	numbers := strings.Split(ip, ".")
	if len(numbers) != 4 {
		return false
	}
	for _, i := range numbers {
		if len(i) != 1 && i[0] == '0'{
			return false
		}
		v, err := strconv.Atoi(i)
		if err != nil {
			return false
		}
		if !(v >= 0 && v <= 255) {
			return false
		}
	}
	return true
}

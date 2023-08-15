package main

import ("strings"
	"strconv"
	"fmt")

func StockList(listArt []string, listCat []string) string {
	 if len(listArt) == 0 || len(listCat) == 0{
		return ""
		}
	categories := make(map[string]int)
	
	for _, v := range listCat {
		categories[v] = 0
	}
	for _, v := range listArt {
		entry := strings.Split(v, " ")
	if _, ok := categories[string(entry[0][0])]; ok {
		s, _ := strconv.Atoi(entry[1])
		categories[string(entry[0][0])] += s
		}
	}
	result := ""
	for _, value := range listCat {
		result += fmt.Sprintf("(%s : %d)", value, categories[value])
		result += " - "
	}
	result = result[:len(result)-3] 
	return result
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	valMaps := make(map[string]int)
	for idx := range arrayA {
		valMaps[arrayA[idx]] = idx
	}
	for idx := range arrayB {
		valMaps[arrayB[idx]] = idx
	}
	var unique []string
	for k := range valMaps {
		unique = append(unique, k)
	}
	return unique
}

func MunculSekali(angka string) []int {
	strArray := strings.Split(angka, "")
	strCounter := make(map[string]int)
	for idx := range strArray {
		strCounter[strArray[idx]] += 1
	}
	var sekali []int
	for k, v := range strCounter {
		if v == 1 {
			i, _ := strconv.Atoi(k)
			sekali = append(sekali, i)
		}
	}
	return sekali
}

func PairSum(arr []int, target int) []int {
	sumMapping := make(map[int]int)
	for idx, val := range arr {
		if sumMapping[target-val] == 0 {
			sumMapping[val] = idx + 1
		} else {
			return []int{sumMapping[target-val] - 1, idx}
		}
	}
	return []int{}
}

func main() {
	// Problem 1: Array Merge
	fmt.Println(ArrayMerge([]string{"kazuya", "jin", "lee"}, []string{"kazuya", "feng"}))
	fmt.Println(ArrayMerge([]string{"lee", "jin"}, []string{"kazuya", "panda"}))
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))

	// Problem 2: Angka Muncul Sekali
	fmt.Println(MunculSekali("76523752"))
	fmt.Println(MunculSekali("1122"))
	fmt.Println(MunculSekali("1234123"))
	fmt.Println(MunculSekali("12345"))
	fmt.Println(MunculSekali("1122334455"))
	fmt.Println(MunculSekali("0872504"))

	// Problem 3: Pair with Target Sum
	fmt.Println(PairSum([]int{1, 2, 3, 4, 5, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}

package main

import "fmt"

func SimpleEquations(a, b, c int) {}

func MoneyCoins(money int) []int {
	coinFractions := []int{
		10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1,
	}
	var coinsUsed []int
	for money > 0 {
		if money >= coinFractions[0] {
			coinsUsed = append(coinsUsed, coinFractions[0])
			money -= coinFractions[0]
		} else if money >= coinFractions[1] {
			coinsUsed = append(coinsUsed, coinFractions[1])
			money -= coinFractions[1]
		} else if money >= coinFractions[2] {
			coinsUsed = append(coinsUsed, coinFractions[2])
			money -= coinFractions[2]
		} else if money >= coinFractions[3] {
			coinsUsed = append(coinsUsed, coinFractions[3])
			money -= coinFractions[3]
		} else if money >= coinFractions[4] {
			coinsUsed = append(coinsUsed, coinFractions[4])
			money -= coinFractions[4]
		} else if money >= coinFractions[5] {
			coinsUsed = append(coinsUsed, coinFractions[5])
			money -= coinFractions[5]
		} else if money >= coinFractions[6] {
			coinsUsed = append(coinsUsed, coinFractions[6])
			money -= coinFractions[6]
		} else if money >= coinFractions[7] {
			coinsUsed = append(coinsUsed, coinFractions[7])
			money -= coinFractions[7]
		} else if money >= coinFractions[8] {
			coinsUsed = append(coinsUsed, coinFractions[8])
			money -= coinFractions[8]
		} else if money >= coinFractions[9] {
			coinsUsed = append(coinsUsed, coinFractions[9])
			money -= coinFractions[9]
		} else if money >= coinFractions[10] {
			coinsUsed = append(coinsUsed, coinFractions[10])
			money -= coinFractions[10]
		}
	}
	return coinsUsed
}

func AscBubbleSort(arr []int) []int{
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if (arr[j] > arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			 }
		}	
	}
	return arr
}

func DescBubbleSort(arr []int) []int{
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if (arr[j] < arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			 }
		}	
	}
	return arr
}

func DragonOfLoowater(dragonHead, knightHead []int) {
	// Assumption:
	// 1) 1 <= dragonHead[i] 
	// 2) 1 <= knightHead[i]
	dragonHeadLen := len(dragonHead)
	knightHeadLen := len(knightHead)
	
	if knightHeadLen < dragonHeadLen {
		fmt.Println("knight fall")
		return
	}

	var dragonHeadSum int
	var selectedKnightHead = make([]int, dragonHeadLen)
	for i := 0; i < dragonHeadLen; i++ {
		dragonHeadSum += dragonHead[i]
		for j := 0; j < knightHeadLen; j++ {
			if knightHead[j] >= dragonHead[i] {
				if (selectedKnightHead[i] == 0) || (knightHead[j] < selectedKnightHead[i]) {
					selectedKnightHead[i], knightHead[j] = knightHead[j], selectedKnightHead[i]
				} 
			}
		}
	}

	var selectedKnightHeadSum int
	for _, val := range selectedKnightHead {
		selectedKnightHeadSum += val
	}
	if dragonHeadSum <= selectedKnightHeadSum {
		fmt.Println(selectedKnightHeadSum)
	} else {
		fmt.Println("knight fall")
	}
}

func BinarySearch(array []int, x int) {
	begIdx := 0
	endIdx := len(array)-1
	
	for begIdx <= endIdx {
		midIdx := int((begIdx+endIdx)/2)
		if array[midIdx] < x {
			begIdx = midIdx + 1
		} else if array[midIdx] > x {
			endIdx = midIdx -  1
		} else {
			fmt.Println(midIdx)
			return
		}
	}
	fmt.Println(-1)
}

func FiboTopDown(n int) {
	previous := 0
	current := 0
	for i := 1; i <= n; i++ {
		if current == 0 {
			current = 1
		} else {
			tmp := current
			current += previous
			previous = tmp
		}
	}
	fmt.Println(current)
}

func FiboBottomUp(n int) {
	f := make([]int32, n+1)
	if len(f) > 0 {
		f[0] = 0
	}
	if len(f) > 1 {
		f[1] = 1
	}
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	
	fmt.Println(f[n])
}

func Abs(num int) int {
	if num < 0 {
		return num * -1 
	}
	return num
}

func Frog(jumps []int) int {
	jumpsNum := len(jumps)
	if jumpsNum < 4 {
		return Abs(jumps[0] - jumps[jumpsNum-1])
	}

	var current, step1, step2, totalCost int
	var i int
	for i = jumpsNum-1; i > jumpsNum-2; {
		current = jumps[i]
		step1 = jumps[i-1]
		step2 = jumps[i-2]
		if Abs(current-step1) < Abs(current-step2) {
			totalCost += Abs(current-step1)
			i--
		} else {
			totalCost += Abs(current-step2)
			i -= 2
		}
	}
	totalCost += Abs(jumps[i]-jumps[0])
	return totalCost
}

func main()  {
	// PART 1: BRUTE FORCE & GREEDY

	// PROBLEM 1: SIMPLE EQUATIONS
	// SimpleEquations(1, 2, 3)
	// SimpleEquations(6, 6, 14)

	// PROBLEM 2: MONEY COINS
	// fmt.Println(MoneyCoins(123))
	// fmt.Println(MoneyCoins(432))
	// fmt.Println(MoneyCoins(543))
	// fmt.Println(MoneyCoins(7752))
	// fmt.Println(MoneyCoins(15321))
	
	// PROBLEM 3: DRAGON OF LOOWATER
	// DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})
	// DragonOfLoowater([]int{5, 10}, []int{5})
	// DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})
	// DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})
	
	// PROBLEM 4: BINARY SEARCH ALGORITHM
	// BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)
	// BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)
	// BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)
	// BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100)

	// PART 2: DYNAMIC PROGRAMMING

	// PROBLEM 1: FIBONACCI NUMBER TOP-DOWN
	FiboTopDown(0)
	FiboTopDown(1)
	FiboTopDown(2)
	FiboTopDown(3)
	FiboTopDown(5)
	FiboTopDown(6)
	FiboTopDown(7)
	FiboTopDown(9)
	FiboTopDown(10)

	// PROBLEM 2: FIBONACCI NUMBER BOTTOM-UP
	// FiboBottomUp(0)
	// FiboBottomUp(1)
	// FiboBottomUp(2)
	// FiboBottomUp(3)
	// FiboBottomUp(5)
	// FiboBottomUp(6)
	// FiboBottomUp(7)
	// FiboBottomUp(9)
	// FiboBottomUp(10)

	// PROBLEM 3: FROG
	// fmt.Println(Frog([]int{10, 30, 40, 20}))
	// fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
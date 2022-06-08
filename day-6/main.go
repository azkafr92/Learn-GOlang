package main

import (
	"fmt"
	"sort"
	"strconv"
)

func PrimeNumber(i int) bool {
	if (i == 2) || (i == 3) || (i == 5) || (i == 7) {
		return true
	} else if (i > 1) && (i%2 != 0) && (i%3 != 0) && (i%5 != 0) && (i%7 != 0) {
		return true
	}
	return false
}

func Pow(x, n int) int {
	a := x
	for i := 1; i < n; i++ {
		a *= x
	}
	return a
}

func PrimeX(number int) int {
	var prime int
	start := 2
	for i := 1; i <= number; {
		if PrimeNumber(start) {
			prime = start
			i++
		}
		start += 1
	}
	return prime
}

func Fibonacci(number int) int {
	previous := 0
	current := 0
	for i := 1; i <= number; i++ {
		if current == 0 {
			current = 1
		} else {
			tmp := current
			current += previous
			previous = tmp
		}
	}
	return current
}

func PrimaSegiEmpat(wide, high, start int) {
	var (
		total   int
		column  int
		numbers string
	)
	for row := 1; row <= high; {
		start += 1
		if PrimeNumber(start) {
			total += start
			numbers += strconv.Itoa(start)
			numbers += " "
			column += 1
		}
		if column == wide {
			column = 0
			row++
			fmt.Println(numbers)
			numbers = ""
		}
	}
	fmt.Println(total)
}

func MaxSequence(arr []int) int {
	var (
		sum     int = arr[0]
		highest int = arr[0]
	)
	for _, val := range arr[1:] {
		if sum+val < val {
			sum = val
		} else {
			sum += val
		}
		if highest < sum {
			highest = sum
		}
	}
	return highest
}

func FindMinAndMax(arr []int) string {
	var (
		minNum int = arr[0]
		minIdx int = 0
		maxNum int = arr[0]
		maxIdx int = 0
	)

	for idx, val := range arr {
		if val < minNum {
			minNum = val
			minIdx = idx
		} else if val > maxNum {
			maxNum = val
			maxIdx = idx
		}
	}

	return fmt.Sprintf("min: %d index: %d max: %d index: %d", minNum, minIdx, maxNum, maxIdx)
}

func MaximumBuyProduct(money int, productPrice []int) {
	sort.Ints(productPrice)
	var item int
	for _, val := range productPrice {
		money -= val
		if money < 0 {
			break
		} else {
			item += 1
		}
	}
	fmt.Println(item)
}

func PlayingDomino(cards [][]int, deck []int) interface{} {
	var (
		deckCardA  int = deck[0]
		deckCardB  int = deck[1]
		highest    int
		highestArr []int
		handCardA  int
		handCardB  int
	)
	if deckCardA > deckCardB {
		deckCardB, deckCardA = deckCardA, deckCardB
	}
	for _, card := range cards {
		handCardA, handCardB = card[0], card[1]
		if handCardA > handCardB {
			handCardB, handCardA = handCardA, handCardB
		}
		if (handCardA == deckCardA) || (handCardB == deckCardB) {
			if (handCardA + handCardB) > highest {
				highest = handCardA + handCardB
				highestArr = []int{card[0], card[1]}
			}
		}
	}
	if highestArr != nil {
		return highestArr
	}
	return "tutup kartu"
}

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	var counter = make(map[string]int, len(items))
	for _, item := range items {
		counter[item] += 1
	}
	
	var pairs []pair
	for k, v := range counter {
		pairs = append(pairs, pair{name: k, count: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count < pairs[j].count {
			return true
		} else if pairs[i].name < pairs[j].name {
			return true
		}
		return false
	})
	return pairs
}

func main() {
	// TIME COMPLEXCITY & SPACE COMPLEXCITY

	// PROBLEM 1: BILANGAN PRIMA
	// fmt.Println(PrimeNumber(1500450271))
	// fmt.Println(PrimeNumber(1000000007))
	// fmt.Println(PrimeNumber(13))
	// fmt.Println(PrimeNumber(17))
	// fmt.Println(PrimeNumber(20))
	// fmt.Println(PrimeNumber(35))

	// PROBLEM 2: FAST EXPONENTIATION
	// fmt.Println(Pow(2, 3))
	// fmt.Println(Pow(5, 3))
	// fmt.Println(Pow(10, 2))
	// fmt.Println(Pow(2, 5))
	// fmt.Println(Pow(7, 3))

	// RECURSIVE, NUMBER THEORY, SORTING, SEARCHING, AND REGEX

	// PROBLEM 1: PRIMA KE-X
	// fmt.Println(PrimeX(1))
	// fmt.Println(PrimeX(5))
	// fmt.Println(PrimeX(8))
	// fmt.Println(PrimeX(9))
	// fmt.Println(PrimeX(10))

	// PROBLEM 2: FIBONACCI (RECURSIVE)
	// fmt.Println(Fibonacci(0))
	// fmt.Println(Fibonacci(2))
	// fmt.Println(Fibonacci(9))
	// fmt.Println(Fibonacci(10))
	// fmt.Println(Fibonacci(12))

	// PROBLEM 3: PRIMA SEGI EMPAT
	// PrimaSegiEmpat(2, 3, 13)
	// PrimaSegiEmpat(5, 2, 1)

	// PROBLEM 4: TOTAL MAKSIMUM DARI DERET BILANGAN
	// fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	// fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	// fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	// fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))

	// PROBLEM 5: FIND MIN AND MAX NUMBER
	// fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	// fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	// fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	// fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	// fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))

	// PROBLEM 6: MAXIMUM BUY PRODUCT
	// MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	// MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
	// MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	// MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})
	// MaximumBuyProduct(0, []int{10000, 30000})

	// PROBLEM 7: PLAYING DOMINO
	// fmt.Println(PlayingDomino(
	// 	[][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}},
	// 	[]int{4, 3},
	// ))
	// fmt.Println(PlayingDomino(
	// 	[][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}},
	// 	[]int{3, 6},
	// ))
	// fmt.Println(PlayingDomino(
	// 	[][]int{{6, 6}, {2, 4}, {3, 6}},
	// 	[]int{5, 1},
	// ))

	// PROBLEM 8: MOST APPEAR ITEM
	// fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// fmt.Println(MostAppearItem([]string{"footbal", "basketball", "tenis"}))
}

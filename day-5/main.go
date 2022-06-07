package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	w := []string{a, b}
	if len(w[0]) > len(w[1]) {
		w[0], w[1] = w[1], w[0]
	}
	if strings.Contains(w[1], w[0]) {
		return w[0]
	}
	return ""
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func Caesar(offset int, input string) string {
	r := []rune(input)
	o := abs(int64(offset)) % 26
	if offset < 0 {
		o = -o
	}
	for i := range r {
		r[i] = (r[i] + int32(o))
		if r[i] > 122 {
			r[i] = (r[i] % 122) + 96
		} else if r[i] < 97 {
			r[i] += 26
		}
	}
	return string(r)
}

func Swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func GetMinMax(numbers ...*int) (min, max int) {
	min, max = *numbers[0], *numbers[0]
	for _, val := range numbers {
		if *val < min {
			min = *val
		}
		if *val > max {
			max = *val
		}
	}
	return min, max
}

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	sum := 0
	for _, val := range s.score {
		sum += val
	}
	return float64(sum / len(s.score))
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]
	for idx := range s.score {
		if s.score[idx] < min {
			min = s.score[idx]
			name = s.name[idx]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for idx := range s.score {
		if s.score[idx] > max {
			max = s.score[idx]
			name = s.name[idx]
		}
	}
	return max, name
}

type student struct {
	name string
	nameEncoded string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	return Caesar(len(s.name), s.name)
}

func (s *student) Decode() string {
	return Caesar(-len(s.nameEncoded), s.nameEncoded)
}

func main() {
	// PROBLEM 1: COMPARE STRING
	// fmt.Println(Compare("AKA", "AKASHI"))
	// fmt.Println(Compare("KANGOORO", "KANG"))
	// fmt.Println(Compare("KI", "KIJANG"))
	// fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	// fmt.Println(Compare("ILALANG", "ILA"))

	// PROBLEM 2: CAESAR CIPHER
	// fmt.Println(Caesar(3, "abc"))
	// fmt.Println(Caesar(2, "alta"))
	// fmt.Println(Caesar(10, "alterraacademy"))
	// fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	// fmt.Println(Caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

	// PROBLEM 3: SWAP TWO NUMBER USING POINTER
	// var (
	// 	a int = 10
	// 	b int = 20
	// )
	// Swap(&a, &b)
	// fmt.Println(a, b)

	// PROBLEM 4: MIN AND MAX USING POINTER
	// var a1, a2, a3, a4, a5, a6 int
	// fmt.Scan(&a1)
	// fmt.Scan(&a2)
	// fmt.Scan(&a3)
	// fmt.Scan(&a4)
	// fmt.Scan(&a5)
	// fmt.Scan(&a6)
	// min, max := GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	// fmt.Println("Nilai min ", min)
	// fmt.Println("Nilai max ", max)

	// PROBLEM 5: STUDENTS SCORE
	// student := Student{
	// 	name: []string{
	// 		"Rizky", "Kobar", "Ismail", "Umam", "Yopan",
	// 	},
	// 	score: []int{
	// 		80, 75, 70, 75, 60,
	// 	},
	// }
	// fmt.Println("Average student score is", student.Average())
	// scoreMin, nameMin := student.Min()
	// fmt.Println("Min student score is", scoreMin, "by", nameMin)
	// scoreMax, nameMax := student.Max()
	// fmt.Println("Max student score is", scoreMax, "by", nameMax)

	// PROBLEM 6: SUBSTITUTION CIPHER
	var (
		s = student{}
		menu int
		c Chiper = &s
	)

	fmt.Printf("[1] Encrypt \n[2] Decypt \nChoose your menu? ")
	fmt.Scan(&menu)
	
	switch menu {
	case 1:
		fmt.Print("\nInput student's name: ")
		fmt.Scan(&s.name)
		fmt.Print("\nEncoded student's name ", s.name, " is ", c.Encode(), "\n")
	case 2:
		fmt.Print("\nInput student's encode name: ")
		fmt.Scan(&s.nameEncoded)
		fmt.Print("\nDecoded student's name ", s.nameEncoded, " is ", c.Decode(), "\n")
	default:
		fmt.Println("Wrong input name menu")
	}
}

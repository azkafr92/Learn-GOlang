package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func LuasPermukaanTabung(T, r float64) float64 {
	// T = tinggi tabung dalam centimeter
	// r = jari-jari tabung dalam centimeter
	return 2 * 3.14 * r * (r + T)
}

func CariGradeNilai(nilai int) string {
	var grade string
	switch {
	case nilai < 35:
		grade = "E"
	case nilai < 50:
		grade = "D"
	case nilai < 65:
		grade = "C"
	case nilai < 80:
		grade = "B"
	case nilai <= 100:
		grade = "A"
	default:
		grade = "Nilai invalid"
	}
	return grade
}

func CariFaktorBilangan(bilangan int) []int {
	var faktorBilangan []int
	batasPembagi := math.Pow(float64(bilangan), 0.5)
	for i := 1; i <= int(math.Floor(batasPembagi)); i++ {
		if bilangan%i == 0 {
			faktorBilangan = append(faktorBilangan, i)
			nilaiPembagi := bilangan / i
			if nilaiPembagi != i {
				faktorBilangan = append(faktorBilangan, nilaiPembagi)
			}
		}
	}
	sort.Ints(faktorBilangan)
	return faktorBilangan
}

func PrimeNumber(number int) string {
	if (number == 2) || (number == 3) || (number == 5) || (number == 7) {
		return fmt.Sprintf("%d adalah bilangan prima", number)
	} else if (number > 1) && (number%2 != 0) && (number%3 != 0) && (number%5 != 0) && (number%7 != 0) {
		return fmt.Sprintf("%d adalah bilangan prima", number)
	}
	return fmt.Sprintf("%d bukan bilangan prima", number)
}

func Palindrome(text string) string {
	lowerText := strings.ToLower(text)
	reverseString := func(text string) string {
		str := []rune(text)
		for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
			str[i], str[j] = str[j], str[i]
		}
		return string(str)
	}
	var result string
	switch lowerText == reverseString(lowerText) {
	case true:
		result = fmt.Sprintf("%s is palindrome", text)
	case false:
		result = fmt.Sprintf("%s is not palindrome", text)
	}
	return result
}

func Pangkat(bilangan, pangkat int64) int64 {
	var result int64 = 1
	var i int64 = 0
	for i < pangkat {
		result *= bilangan
		i++
	}
	return result
}

func PlayWithAsterix(baris int) {
	for i := 1; i <= baris; i++ {
		for j := 1; j <= baris-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k != (2*i - 1); k++ {
			if k%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func CetakPerkalian(number int) [][]string {
	arr := make([][]string, number)
	for i := range arr {
		for j := range arr {
			arr[i] = append(arr[i], strconv.Itoa((i + 1) * (j + 1)))
		}
		fmt.Println(strings.Join(arr[i], " "))
	}
	return arr
}

func main() {
	// fmt.Println("Problem 1: Menghitung luas permukaan tabung")
	// fmt.Print("Tinggi tabung (cm): ")
	// var T float64
	// fmt.Scanln(&T)
	// fmt.Print("Jari-jari tabung (cm): ")
	// var r float64
	// fmt.Scanln(&r)
	// fmt.Printf("luas permukaan tabung: %v cm kubik \n", LuasPermukaanTabung(T, r))

	// fmt.Println("Problem 2: Grade nilai")
	// fmt.Print("Masukkan nilai ujian mahasiswa: ")
	// var studentScore int
	// fmt.Scanln(&studentScore)
	// fmt.Printf("Grade nilai mahasiswa: %s \n", CariGradeNilai(studentScore))

	// fmt.Println("Problem 3: Faktor bilangan")
	// fmt.Printf("Bilangan: ")
	// var bilangan int
	// fmt.Scanln(&bilangan)
	// fmt.Printf("Faktor bilangan %d adalah %v \n", bilangan, CariFaktorBilangan(bilangan))

	// fmt.Println("Problem 4: Bilangan prima")
	// fmt.Printf("Bilangan: ")
	// var bilPrima int
	// fmt.Scanln(&bilPrima)
	// fmt.Printf("%s \n", PrimeNumber(bilPrima))

	// fmt.Println("Problem 5: Palindrome")
	// fmt.Printf("Kata: ")
	// var kata string
	// fmt.Scanln(&kata)
	// fmt.Printf("%s \n", Palindrome(kata))

	// fmt.Println("Problem 6: Exponentiation")
	// fmt.Printf("Bilangan: ")
	// var x int64
	// fmt.Scanln(&x)
	// fmt.Printf("Pangkat: ")
	// var n int64
	// fmt.Scanln(&n)
	// fmt.Printf("%d pangkat %d: %d \n", x, n, Pangkat(x, n))

	// fmt.Println("Problem 7: Play with asterix")
	// fmt.Printf("Berapa baris? ")
	// var baris int
	// fmt.Scanln(&baris)
	// PlayWithAsterix(baris)

	fmt.Println("Problem 8: Cetak Tabel Perkalian")
	CetakPerkalian(9)
}

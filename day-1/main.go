package main

import "fmt"

func IsPrimeNumber(i int) string {
	if (i == 2) || (i == 3) || (i == 5) || (i == 7) {
		return fmt.Sprintf("%d is prime number", i)
	} else if (i > 1) && (i%2 != 0) && (i%3 != 0) && (i%5 != 0) && (i%7 != 0) {
		return fmt.Sprintf("%d is prime number", i)
	}
	return fmt.Sprintf("%d is not prime number", i)
}

func LampAndButton(N int) string {
	state := "lampu mati"

	changeLampState := func(currentState string) string {
		if (currentState == "lampu mati") {
			return "lampu menyala"
		}
		return "lampu mati"
	}

	for i := 1; i <= N; i++{
		if (N%i == 0 ) {
			state = changeLampState(state)
		}
	}
	return state
}

func main() {
	fmt.Println("hello world!")

	fmt.Println(IsPrimeNumber(3))
	fmt.Println(IsPrimeNumber(7))
	fmt.Println(IsPrimeNumber(10))
	fmt.Println(IsPrimeNumber(-1))

	fmt.Println(LampAndButton(5))
	fmt.Println(LampAndButton(4))
}

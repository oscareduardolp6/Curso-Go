package main

import "fmt"

func main() {
	division(10, 2)
	division(50, 5)
	division(50, 0)
	division(50, 25)

}

func division(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recuperandome del pánico %v \n", r)
		}
	}()

	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("Ya valió :(")
	}
}

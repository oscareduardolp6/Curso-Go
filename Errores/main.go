package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("archivo.txt") // Para forzar el error
	if err != nil {
		fmt.Printf("Error: %v", err)
		return //Esto finaliza la ejecución del programa
	}
	fmt.Println(string(content))

	resultado, err := div(10, 2)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println(resultado)
	nums := []int{2, 4, 5, 6, 7, 8, 4}
	resultado2 := filter(nums, func(num int) bool {
		if num <= 10 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(resultado2)
	x := hello("Alejandro")()
	fmt.Println(x)
}

/* Funciones variaticas, funciones con número de parametros no especificado */
func sum(nums ...int) int {
	result := 0
	for _, value := range nums {
		result += value
	}
	return result
}

/* Funciones que regresan funciones */
func hello(name string) func() string {
	return func() string {
		return "Hello " + name
	}
}
func filter(nums []int, callback func(int) bool) []int {
	result := []int{}
	for _, value := range nums {
		if callback(value) {
			result = append(result, value)
		}
	}
	return result
}

/* Parametros y returns nombrados */
func div2(dividendo, divisor int) (resultado int, err error) {
	if divisor == 0 {
		err = errors.New(("No se puede dividir entre cero")) // A la hora de nombrarlos Go los inicializa, por lo tanto devovería 0 y el error
		return
	}
	resultado = dividendo / divisor // err regresaría cómo nil
	return
}

func div(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede divir entre 0")
	}
	return dividendo / divisor, nil
}

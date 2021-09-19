package main

import (
	"fmt"
	"strings"
)

func main() {
	emoji := "Valor inicial"
	change(&emoji)
	hello("oscar", "lopez")
	result := sum(2, 3)
	fmt.Println(result)
	min, may := convert("Oscar")
	fmt.Println(min, may)
	/* Funciones anónimas */
	x := func() {
		fmt.Println("Función anónima")
	} //() Función anónima autoejecutada
	x()
}

func hello(name string, lastname string) {
	fmt.Printf("Hello %s %s", name, lastname)
}

/* Por defecto las funciones reciben parametros por valor */
func change(valor *string) {
	*valor = "cambio de valor"
}

func sum(num1, num2 int) int {
	return num1 + num2
}

/* Funcion con que regresa varios valores */
func convert(texto string) (string, string) {
	min := strings.ToLower(texto)
	max := strings.ToUpper(texto)
	return min, max
}

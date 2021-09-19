package main

import "fmt"

func main() {

	for i := 10; i < 10; i++ {

	}

	/*For continuo,más o menos cómo el while */
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	/* For ever, un ciclo infinito */
	// for {

	// }

	nums := []uint8{2, 4, 6, 8}

	for index, value := range nums {
		fmt.Println("Cualquier cosa")
		fmt.Println(index, value) //Index y value son los valores en el arreglo, no muta el arreglo
		nums[i] *= 2              // Esto sí cambia el valor en el slice
	}

	sports := map[string]string{"basquet": "baloncesto", "futball": "futbol", "baseball": "base"}
	for key, value := range sports {
		fmt.Println(key, value)
	}

	hello := "Hello"
	for index, value := range hello {
		/* Value nos va a devolver el valor de char, o sea, un número
		Lo vamos a tener que castear para que sea letra */
		fmt.Println(index, string(value))
	}

}

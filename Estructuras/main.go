package main

import "fmt"

func main() {
	/* Las estructuras son similares a las clases */
	type course struct {
		Name      string
		Professor string
		Country   string
	}

	db := course{
		Name:      "Bases de datos",
		Professor: "Alexis",
		Country:   "Colombia",
	}

	/* Creación de una estructura literal */
	git := course{"Git", "Beto", "Bolivia"}

	/*Puedes no darle el valor a cada una de las propiedas, éstas se van a inicializar en el tipo
	inicial de su tipo "", 0, nil, etc */

	fmt.Printf("%+v", db) //El verbo '%+v nos permite imprimir los campos de la estructura
	fmt.Printf("%+v", git)

	/* Acceder a propiedades */
	var _ = db.Country

	/* Crear apuntadores a estructuras */
	p := &db
	/* Cambiar valores dentro de la estructura, desde el apuntador */
	// (*p).Professor = "Beto"
	p.Professor = "Beto"

}

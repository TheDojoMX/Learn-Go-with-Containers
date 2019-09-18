package main

import (
	"fmt"
)

func main() {
	name := "Héctor"
	age := 30
	w := 90.5

	// Funciones que empiezan con S devuelven el valor
	// fmt.Println(fmt.Sprintf("Hola %s", name))

	// Funciones que terminan con f -> formatean
	fmt.Printf("Hola %v\n", name)
	fmt.Printf("Tiene %v años\n", age)
	fmt.Printf("Pesa %v kilos\n", w)
	// fmt.Printf("Hola %s", name)

	// Funciones que terminan con ln imprimen y agregan nueva línea
}

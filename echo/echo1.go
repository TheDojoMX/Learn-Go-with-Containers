package main

import (
	"fmt"
	"os"
)

// Recibir parámetros de consola e imprimirlos de regreso

func main() {
	// Leer argumentos del programa
	// os.Args[0] <- Nombre del program
	var imprimir, sep string
	// sep := ""
	for i := 2; i < len(os.Args); i++ {
		imprimir += sep + os.Args[i]
		sep = os.Args[1]
	}
	fmt.Println(imprimir)
	// fmt.Printf("%v"+sep, os.Args[i]) // Printf
}

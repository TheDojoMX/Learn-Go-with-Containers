package main

import (
	"fmt"
	"os"
)

// Recibir parámetros de consola e imprimirlos de regreso

func main() {
	// Leer argumentos del programa
	// os.Args[0] <- Nombre del program
	// var imprimir, sep string
	// sep := ""
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
		// fmt.Printf("%v"+sep, os.Args[i]) // Printf
	}
}

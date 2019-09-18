package main

import (
	"fmt"
	"os"
	"strings"
)

// Recibir parámetros de consola e imprimirlos de regreso

func main() {
	res := strings.Join(os.Args[1:], " ")
	fmt.Println(res)
}

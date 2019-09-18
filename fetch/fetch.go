package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://google.com"
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error trayendo %s, el error fue: %s", url, err)
		return
	}

	// Cerrando el cuerpo después de terminar la función
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error leyendo la respuesta de %s, error: %s", url, err)
	}
	fmt.Printf("%s", body)
}

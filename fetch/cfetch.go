package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	// "time"
)

func main() {
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Levantar procesos independientes
	}

	// Recibe tantos mensajes como argumentos se reciban / o sitios se visiten
	for range os.Args[1:] {
		res := <-ch // Bloquea hasta que reciba un mensaje
		fmt.Printf(res)
	}

}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
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
	ch <- fmt.Sprintf("%s -> %s  - %d \n", url, r.Status, len(body))
}

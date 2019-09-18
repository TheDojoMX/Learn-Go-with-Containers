package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	for _, url := range os.Args[1:] {
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
		fmt.Printf("%s -> %s  - %d \n", url, r.Status, len(body))
	}
	diff := time.Since(startTime).Seconds()
	fmt.Printf("El proceso tomó %.2f", diff)
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]
	fmt.Println(fetch(urls))
}

func fetch(urls []string) (result string) {
	for _, url := range urls {
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
		result += fmt.Sprintf("%s -> %s  - %d \n", url, r.Status, len(body))
	}
	return result
}

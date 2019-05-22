package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/zerofelx/IRCLoL/API"
	"github.com/zerofelx/IRCLoL/css"
	"github.com/zerofelx/IRCLoL/index"
)

// Función principal donde se crea el servidor
func main() {
	fmt.Print("Corriendo servidor!")

	http.HandleFunc("/", handler)

	http.HandleFunc("assets/main.css", func(arg2 http.ResponseWriter, f *http.Request) {
		io.WriteString(arg2, css.Main())
	})

	http.ListenAndServe(":8000", nil)
}

// El manejador del servidor
func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, indexHTML())
}

// Este es el index de la página donde se recibe la API y se inserta en el index
func indexHTML() string {

	return index.Header() + `
			<div class="Campeones">
				` + campeones() + `
			</div>
		` + index.Footer()

}

func campeones() string {
	var indexar bytes.Buffer

	for i := 0; i < 7; i++ {
		indexar.WriteString(crearIndex(i))
	}

	return indexar.String()
}

func crearIndex(i int) string {
	Campeones := API.API()

	return `<div id="` + Campeones[i].Nombre + `">` + `
	<h1>` + Campeones[i].Nombre + `</h1>` + `
	<h2>` + Campeones[i].Subnombre + `</h2>
	</div>
	`
}

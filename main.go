package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/zerofelx/IRCLoL-Github/API"
	"github.com/zerofelx/IRCLoL-Github/css"
	"github.com/zerofelx/IRCLoL-Github/index"
)

// Función principal donde se crea el servidor
func main() {
	saludo()

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
	var cantidadDeCampeones = len(API.API())

	for i := 1; i < cantidadDeCampeones; i++ {
		indexar.WriteString(crearIndex(i))
	}

	return indexar.String()
}

func crearIndex(i int) string {
	Campeones := API.API()
	C := Campeones[i]
	return `<div style="background-image: url('')" id="` + C.Nombre + `" class="bloqueChamp"  onclick="VerCampeones(` + strconv.Itoa(C.Numero) + `)">` + `
	<h1>` + C.Nombre + `</h1>` + `
	<h2>` + C.Subnombre + `</h2>
	</div>
	`
}

func saludo() {
	var champs = len(API.API())
	fmt.Println("Actualmente existen " + strconv.Itoa((champs - 1)) + " campeones en la base de datos!")
	fmt.Println()
	fmt.Println("Hola!")
	fmt.Println("Bienvenido a IRCgo!")
	fmt.Println("Yo soy el servidor local donde está almacenada la información! :D")
	fmt.Println("Si me cierras no podrás hacer ninguna consulta! :)")
	fmt.Println()
	fmt.Println("Puedes también evitar cerrarme mientras juegas!")
	fmt.Println("Puedes cerrar la ventana donde haces las consultas y dejarme minimizado")
	fmt.Printf("Cuando vayas a volver a hacer consultas entra a la carpeta GUI e inicia 'ircgo.exe' :D ")
	fmt.Print("\n")
	fmt.Print("No olvides cerrarme cuando termines! :)")
}

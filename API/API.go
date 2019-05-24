package API

import (
	"encoding/json"
	"net/http"
	"time"
)

// Estructura del JSON de mi firebase
type campeonesFirebase struct {
	Nombre    string `json:"nombre"`
	Subnombre string `json:"subnombre"`
	Numero    int    `json:"numero"`
	P         string `json:"P"`
	Q         string `json:"Q"`
	W         string `json:"W"`
	E         string `json:"E"`
	R         string `json:"R"`
	Counters  []int  `json:"counters"`
	CountersB []int  `json:"countersB"`
}

// Llama la API y devuelve el JSON con los campeones
func API() []campeonesFirebase {
	var cliente = &http.Client{Timeout: 10 * time.Second}
	var url = "https://irc-lol.firebaseio.com/campeonesFirebase.json"

	resp, err := cliente.Get(url)

	if err != nil {
		panic(err.Error())
	}

	var Campeon []campeonesFirebase

	err = json.NewDecoder(resp.Body).Decode(&Campeon)

	if err != nil {
		panic(err.Error())
	}

	return Campeon
}

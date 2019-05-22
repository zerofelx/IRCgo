package index

import "github.com/zerofelx/IRCLoL/css"

// Header de la página inicial
func Header() string {
	return `<!DOCTYPE html>` +
		`<html lang="es">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Champs</title>
		<link rel="stylesheet" href="assets/main.css" type="text/css">
		<style>` + css.Main() + `</style>
	</head>
	<body>
		<div class="main">
			
	`
}

// Footer de la página inicial
func Footer() string {
	return `</div>` + `
	</body>
	</html>`
}

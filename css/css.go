package css

// Main es la función que retorna el código main del css
func Main() string {
	return `body {
		font-size: 'Roboto';
		margin: 0;
	}
	.center {
		text-align: center;
	}
	body::-webkit-scrollbar {
		display: none;
	}
	`
}

package css

// Main es la función que retorna el código main del css
func Main() string {
	return `body {
		font-size: 'Roboto';
		margin: 0;
		background: #FCFCFC;
	}
	body::-webkit-scrollbar {
		display: none;
	}
	.center {
		text-align: center;
	}

	#buscador {
		margin-bottom: 20px;
		width: 100%;
		text-align: center;
		background: white;
		padding: 10px 0;
	}
	#buscarCampeon {
		padding: 10px 5px;
		width: 80%
	}


	.Campeones {
		display: flex;
		flex-wrap: wrap;
		text-align: center;
		justify-content: center;
	}
	.bloqueChamp {
		min-width: 20%;
		border: 2px solid white;
		padding: 4px;
		margin: 5px;
		background: white;
		box-shadow: 1px 1px 2px grey;
	}
	.bloqueChamp:hover {
		border: 2px solid grey;
		box-shadow: 1px 1px 2px grey;
	}
	.bloqueChamp:active {
		box-shadow: 1px 1px 2px green;
		border: 2px solid green;
	}
	`
}

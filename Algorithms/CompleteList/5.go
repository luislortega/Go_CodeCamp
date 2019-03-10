package main

import "fmt"

func main() {
	const size = 26
	/* entradas */
	var people [size]int
	var prom [20]int
	entradas(&people)
	proceso(&people, &prom)
	salidas(&people, &prom)
}

func entradas(people *[26]int) {
	for i := 1; i < 26; i++ {
		fmt.Println("Ingresa la calificacion: ", i)
		fmt.Scan(&people[i])
	}
}

func proceso(people *[26]int, prom *[20]int) {
	var compuerta bool = false
	var acumulador int = 0
	var contadorProm int = 0
	for i := 1; i < 26; i++ {
		acumulador += people[i]

		if i%5 == 0 {
			compuerta = true
		} else {
			compuerta = false
		}

		if compuerta {
			contadorProm++
			prom[contadorProm] = acumulador / 5
			acumulador = 0
		}
	}
}

func salidas(people *[26]int, prom *[20]int) {
	var compuerta bool = false
	var contadorProm int = 0
	for i := 1; i < 26; i++ {
		fmt.Print("| ", people[i], " |")

		if i%5 == 0 {
			compuerta = true
		} else {
			compuerta = false
		}

		if compuerta {
			contadorProm++
			fmt.Println("PROM: ", prom[contadorProm])
		}
	}
}

package main

import "fmt"

func main() {
	fmt.Printf("Cular el área de un cuadrado \n\n")
	fmt.Printf("Ingresa el tamaño de uno de sus lados: ")

	var lado float64

	fmt.Scanf("%f", &lado)

	area := lado * 4

	fmt.Print("El área del cuadrado es: ", area)
}

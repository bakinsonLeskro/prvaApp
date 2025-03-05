package main

import (
	"fmt"
)

type Dio string

func printAssemblyLine(lista []Dio) {
	for i := 0; i < len(lista); i++ {
		trenDio := lista[i]
		trenBroj := i + 1
		fmt.Println("Dio broj", trenBroj, "je", trenDio)
	}
}
func main() {
	djelovi := []Dio{"Karburator", "Akumulator", "Pumpa"}
	printAssemblyLine(djelovi)
	djelovi = append(djelovi, "Cilindar", "Hladnjak")
	printAssemblyLine(djelovi)
	djelovi = djelovi[3:]
	printAssemblyLine(djelovi)

}

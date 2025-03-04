package main

import "fmt"

func pozdrav() string {
	return "Cao!"
}
func pozdravTebi(ime string) {
	fmt.Printf("Cao, %s! \n", ime)
}
func saberiTri(a, b, c int) int {
	return a + b + c
}
func main() {
	ime := "Balsa"
	pozdravTebi(ime)
	fmt.Println(pozdrav())
	fmt.Println(saberiTri(1, 2, 3))
}

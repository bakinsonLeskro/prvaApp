package main

import "fmt"

type pravougaonik struct {
	duzina int
	sirina int
}

func obim(pr pravougaonik) int {
	return 2 * (pr.duzina + pr.sirina)
}
func povrsina(pr pravougaonik) int {
	return pr.duzina * pr.sirina
}
func main() {
	kvadrat := pravougaonik{5, 5}
	fmt.Println(obim(kvadrat))
	fmt.Println(povrsina(kvadrat))
	kvadrat.duzina *= 2
	kvadrat.sirina *= 2
	fmt.Println(obim(kvadrat))
	fmt.Println(povrsina(kvadrat))
}

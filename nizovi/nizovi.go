package main

import "fmt"

type proizvod struct {
	cijena float64
	ime    string
}

func pregledajListu(lista [4]proizvod) {
	ukupnaCijena := 0.0
	brojNaListi := 0

	for i := 0; i < len(lista); i++ {
		// Ako je ime prazno, znači da smo stigli do kraja validnih unosa
		if lista[i].ime == "" {
			break
		}
		brojNaListi++
		ukupnaCijena += lista[i].cijena
	}

	// Poslednji validni element (onaj na indeksu brojNaListi-1)
	if brojNaListi > 0 {
		fmt.Println("Poslednji clan niza je:", lista[brojNaListi-1].ime)
	}

	fmt.Printf("Broj clanova na shopping listi je: %d i njihova ukupna cijena je: %.2f\n", brojNaListi, ukupnaCijena)
}

func main() {
	jaja := proizvod{2.5, "Radanovic"}
	mlijeko := proizvod{1.1, "Kravica"}
	hleb := proizvod{0.8, "Inpek"}
	cokolada := proizvod{3.2, "Najljepse zelje"}

	// Pravi se niz od 4 elementa, ali samo prva 3 su inicijalizovana
	shoppingLista := [4]proizvod{
		jaja,
		mlijeko,
		hleb,
	}

	fmt.Println("\nPrvo ispitivanje liste:")
	pregledajListu(shoppingLista)

	// Dodajemo četvrti proizvod u niz
	shoppingLista[3] = cokolada

	fmt.Println("\nNakon dodavanja cokolade:")
	pregledajListu(shoppingLista)
}

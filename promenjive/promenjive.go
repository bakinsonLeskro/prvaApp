package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func jeValidanDatum(datum string) bool {
	pattern := regexp.MustCompile(`^\d{2}[.-]\d{2}[.-]\d{4}[.-]$`)
	if !pattern.MatchString(datum) {
		return false
	}

	// Izdvajanje dijelova datuma (dd, mm, yyyy)
	separator := string(datum[2]) // Uzima prvi separator
	parts := strings.Split(datum, separator)
	if len(parts) != 4 { // Split će vratiti 4 dijela jer string završava separatorom
		return false
	}
	danStr, mjesecStr, godinaStr := parts[0], parts[1], parts[2]

	// Pretvaranje u brojeve
	dan, err := strconv.Atoi(danStr)
	if err != nil {
		return false
	}
	mjesec, err := strconv.Atoi(mjesecStr)
	if err != nil {
		return false
	}
	godina, err := strconv.Atoi(godinaStr)
	if err != nil {
		return false
	}

	// Provjera mjeseca: 1-12
	if mjesec < 1 || mjesec > 12 {
		return false
	}

	// Provjera dana: 1-31 (osnovna provjera)
	if dan < 1 || dan > 31 {
		return false
	}

	// Provjera godine: npr. 1900-2100
	if godina < 1900 || godina > 2100 {
		return false
	}

	return true
}

func main() {
	var datumRodjenja string
	fmt.Print("Unesite Vas datum rodjenja u dd.mm.yyyy. formatu:")
	fmt.Scan(&datumRodjenja)
	if jeValidanDatum(datumRodjenja) != true {
		fmt.Println("Pogresno unijet datum") // Ako je datum validan, samo onda se radi konverzija u dane
	} else {
		brDana, brDana2 := 0, 0
		today := time.Now().Format("02.01.2006.")
		duzina := []rune(datumRodjenja) // U duzinu ide broj runa sadrzanih u stringu
		for i := 0; i < len(duzina); i++ {
			r, _ := utf8.DecodeRuneInString(datumRodjenja[i:]) // uzima prvi rune iz splittovanog stringa datuma

			if r == '.' {
				continue
			}
			b, err := strconv.Atoi(string(r))
			if err != nil {
				fmt.Println("Greska pri konverziji:", err)
				return
			}
			if b >= 0 && b <= 9 { // Svaki mjesec mnozim sa 30 dana, svaku godinu sa 365 dana, itd...
				switch i {
				case 0:
					brDana += b * 10
				case 1:
					brDana += b
				case 3:
					brDana += b * 10 * 30
				case 4:
					brDana += b * 30
				case 6:
					brDana += b * 1000 * 365
				case 7:
					brDana += b * 100 * 365
				case 8:
					brDana += b * 10 * 365
				case 9:
					brDana += b * 365
				}

			}

		}
		duzina2 := []rune(today)
		for i := 0; i < len(duzina2); i++ {
			s, _ := utf8.DecodeRuneInString(today[i:])
			if s == '.' {
				continue
			}
			a, err := strconv.Atoi(string(s))
			if err != nil {
				fmt.Println("Greska pri konverziji:", err)
				return
			}
			if a >= 0 && a <= 9 {
				switch i {
				case 0:
					brDana2 += a * 10
				case 1:
					brDana2 += a
				case 3:
					brDana2 += a * 10 * 30
				case 4:
					brDana2 += a * 30
				case 6:
					brDana2 += a * 1000 * 365
				case 7:
					brDana2 += a * 100 * 365
				case 8:
					brDana2 += a * 10 * 365
				case 9:
					brDana2 += a * 365
				}

			}

		}
		fmt.Println("Stari ste :", (brDana2-brDana)*24, "sati.")
	}
}

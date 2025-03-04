package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bacaj(brk, brs, brb int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	poeni := 0
	for i := 0; i != brb; i++ {
		bacanja := make([]int, brk)
		zbir := 0
		snakeEyes := true
		fmt.Printf("\nBacanje broj %d", i+1)
		for j := 0; j != brk; j++ {
			bacanja[j] = 1 + rng.Intn(brs)
			fmt.Printf("\nKocka broj %d je %d", j+1, bacanja[j])
			zbir += bacanja[j]
			if bacanja[j] != 1 {
				snakeEyes = false
			}
		}
		fmt.Printf("\nZbir ovog bacanja: %d", zbir)
		if zbir == 7 {
			fmt.Print("\nLucky 7!")
		}
		if snakeEyes == true {

			fmt.Print("\nSnake eyes!")
		}
		poeni += zbir
	}
	if poeni%2 == 1 {
		fmt.Printf("\n%d Odd!", poeni)
	} else {
		fmt.Printf("\n%d Even!", poeni)
	}

}
func main() {
	brKocki, brStrana, brBacanja := 3, 3, 6
	bacaj(brKocki, brStrana, brBacanja)
}

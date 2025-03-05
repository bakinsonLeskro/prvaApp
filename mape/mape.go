package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func serverInfo(serverData map[string]int) {
	brojServera := 0
	brOn, brOf, brMa, brRe := 0, 0, 0, 0
	for _, status := range serverData {
		brojServera++
		switch status {
		case 0:
			brOn++
		case 1:
			brOf++
		case 2:
			brMa++
		case 3:
			brRe++
		}
	}
	fmt.Println("Broj servera je", brojServera, "a statusi su:\nOnline:", brOn, "\nOffline:", brOf, "\nMaintenance:", brMa, "\nRetired:", brRe)

	statusProvjera := 0

	for range len(serverData) {
		switch statusProvjera {
		case 0:
			fmt.Println("\nOnline serveri:")
		case 1:
			fmt.Println("\nOffline serveri:")
		case 2:
			fmt.Println("\nServeri pod maintanence-om:")
		case 3:
			fmt.Println("\nRetired serveri:")
		}
		for ime, status := range serverData {
			if status == statusProvjera {
				fmt.Println(ime)
			}
		}
		statusProvjera++
	}
	fmt.Println("\n#######################")
}
func main() {
	serveri := []string{"hamachi", "aternos", "mineGo", "hypixel", "mineplex"}
	serverStatus := make(map[string]int)
	for i := range serveri {
		serverStatus[serveri[i]] = 0
	}
	serverStatus["mineplex"] = 3
	serverStatus["hamachi"] = 2
	serverInfo(serverStatus)
	for ime := range serverStatus {
		serverStatus[ime] = 2
	}

}

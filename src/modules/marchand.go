package modules

import (
	"fmt"
	"strconv"
)

func marchand(c *Character, inv_marchand map[string]int) map[string]int {
	for {
		if len(inv_marchand) == 0 {
			fmt.Println("La boutique est vide")
		} else {
			i := 1
			for key, value := range inv_marchand {
				fmt.Print(i, "-", key, " * ", value, "\n")
				i += 1
			}
		}
		fmt.Println("\n0-Retour\n")
		var input string
		fmt.Scan(&input)
		fmt.Print("\n")
		index, err := strconv.Atoi(input)
		var keys []string
		for key := range inv_marchand {
			keys = append(keys, key)
		}
		if index == 0 && err == nil {
			return inv_marchand
		} else if index > 0 && index <= len(keys) {
			index--
			fmt.Print(keys[index], "\n\n")
			c.inventaire = addInventory(c.inventaire, keys[index])
			inv_marchand = removeInventory(inv_marchand, keys[index])
		} else {
			fmt.Println("Commande inconnue")
		}
	}
}

package modules

import (
	"fmt"
	"strconv"
)

func marchand(c *Character, inv_marchand []string) []string{
	for {
		for i, o := range inv_marchand {
			fmt.Print(i+1, "-", o, "\n")
		}
		fmt.Println("\n0-Retour\n")
		var input string
		fmt.Scan(&input)
		fmt.Print("\n")
		index, err := strconv.Atoi(input)
		if index == 0 && err == nil {
			return inv_marchand
		} else if index > 0 && index <= len(inv_marchand) {
			index--
			if c.inventaire[inv_marchand[index]].prix <= c.argent {
				fmt.Print(inv_marchand[index], "\n\n")
				addInventory(c, inv_marchand[index])
				c.argent -= c.inventaire[inv_marchand[index]].prix
			} else {
				fmt.Println("Pas assez d'argent")
			}
		} else {
			fmt.Println("Commande inconnue")
		}
	}
}

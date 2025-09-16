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
			fmt.Print(inv_marchand[index], "\n\n")
			addInventory(c, inv_marchand[index])
		} else {
			fmt.Println("Commande inconnue")
		}
	}
}

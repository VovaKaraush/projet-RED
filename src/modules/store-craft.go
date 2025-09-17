package modules

import (
	"fmt"
	"strconv"
	"sort"
)

func marchand(c *Character, inv_marchand []string) {
	for {
		i := 1
		for _, o := range inv_marchand {
			if c.inventaire[o].type_objet != 2 {
				fmt.Print(i, "-", o, "\n")
				i += 1
			}
		}
		fmt.Println("\n0-Retour\n")
		var input string
		fmt.Scan(&input)
		fmt.Print("\n")
		index, err := strconv.Atoi(input)
		if index == 0 && err == nil {
			return
		} else if index > 0 && index <= len(inv_marchand) {
			index--
			if c.inventaire[inv_marchand[index]].prix <= c.argent {
				if !inventoryFull(c) {
					fmt.Print(inv_marchand[index], "\n\n")
					addInventory(c, inv_marchand[index])
					c.argent -= c.inventaire[inv_marchand[index]].prix
				} else {
					fmt.Println("Pas de place dans l'inventaire")
				}
			} else {
				fmt.Println("Pas assez d'argent")
			}
		} else {
			fmt.Println("Commande inconnue")
		}
	}
}

func forgeron(c *Character, liste_armure map[string]Objet_Equipement) {
	for {
		var keys []string
		for key, value := range c.inventaire {
			if value.type_objet == 2 {
				keys = append(keys, key)
			}
		}
		sort.Slice(keys, func(i, j int) bool {
			return c.inventaire[keys[i]].id < c.inventaire[keys[j]].id
		})
		for i, o := range keys {
			fmt.Print(i+1, "-", o, " | ")
			for key, value := range liste_armure[o].recette {
				fmt.Print(key, " * ", value, " | ")
			}
			fmt.Print("\n")
		}
		fmt.Println("\n0-Retour\n")
		var input string
		fmt.Scan(&input)
		index, err := strconv.Atoi(input)
		index--
		if index == -1 && err == nil {
			return
		} else if index > -1 && index < len(keys) {
			found := true
			for key, value := range liste_armure[keys[index]].recette {
				if c.inventaire[key].quantite < value {
					found = false
					break
				}
			}
			if found {
				if c.argent > 4 {
					fmt.Print(keys[index], "\n\n")
					addInventory(c, keys[index])
					c.argent -= 5
					for key, value := range liste_armure[keys[index]].recette {
						temp := c.inventaire[key]
						temp.quantite -= value
						c.inventaire[key] = temp
					}
				} else {
					fmt.Println("Pas assez d'argent")
				}
			} else {
				fmt.Println("Pas assez de ressources")
			}
		} else {
			fmt.Println("Commande inconnue")
		}
	}
}

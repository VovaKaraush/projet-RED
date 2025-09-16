package modules

import (
	"fmt"
	"strconv"
	"time"
)

func accessInventory(c *Character) {
	for {
		if len(c.inventaire) == 0 {
			fmt.Println("L'inventaire est vide")
		} else {
			i := 1
			fmt.Println("Inventaire :")
			for key, value := range c.inventaire {
				fmt.Print(i, "-", key, " * ", value, "\n")
				i += 1
			}
		}
		fmt.Println("\n0-Retour\n")
		var input string
		fmt.Scan(&input)
		index, err := strconv.Atoi(input)
		index--
		var keys []string
		for key := range c.inventaire {
			keys = append(keys, key)
		}
		if index == -1 && err == nil {
			return
		} else if index > -1 && index < len(keys) {
			switch keys[index] { //appel des fonctions associées aux objets
			case "Potion de vie":
				takePot(c)
			case "Potion de poison":
				poisonPot(c)
			case "Livre de sort : Boule de feu":
				spellBook(c)
			}
		} else {
			fmt.Println("Commande inconnue")
		}
	}
}

func takePot(c *Character) {
	if _, ok := c.inventaire["Potion de vie"]; ok {
		c.pv += 50
		if c.pv > c.pvMax {
			c.pv = c.pvMax
		}
		fmt.Print("Vie : ", c.pv, "/", c.pvMax, "\n")
		c.inventaire = removeInventory(c.inventaire, "Potion de vie")
	} else {
		fmt.Println("Pas de potion de vie dans l'inventaire")
	}
}

func poisonPot(c *Character) {
	if _, ok := c.inventaire["Potion de poison"]; ok {
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			c.pv -= 10
			fmt.Print("Vie : ", c.pv, "/", c.pvMax, "\n")
		}
		c.inventaire = removeInventory(c.inventaire, "Potion de poison")
	} else {
		fmt.Println("Pas de potion de poison dans l'inventaire")
	}
}

func spellBook(c *Character) {
	found := false
	for _, s := range c.skill {
		if s == "Boule de feu" {
			found = true
			break
		}
	}
	if !found {
		c.skill = append(c.skill, "Boule de feu")
		c.inventaire = removeInventory(c.inventaire, "Livre de sort : Boule de feu")
	} else {
		fmt.Println("Sort déjà appris")
	}
}

func addInventory(inv map[string]int, objet string) map[string]int {
	inv[objet] += 1
	return inv
}

func removeInventory(inv map[string]int, objet string) map[string]int {
	if val, ok := inv[objet]; ok {
		if val > 1 {
			inv[objet] -= 1
		} else {
			delete(inv, objet)
		}
	}
	return inv
}

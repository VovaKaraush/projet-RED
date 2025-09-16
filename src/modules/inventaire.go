package modules

import (
	"fmt"
	"strconv"
	"time"
	"sort"
)

func addInventory(c *Character, objet string) {
	temp := c.inventaire[objet]
	temp.quantite += 1
	c.inventaire[objet] = temp
}

func removeInventory(c *Character, objet string) {
	if c.inventaire[objet].quantite > 0 {
		temp := c.inventaire[objet]
		temp.quantite -= 1
		c.inventaire[objet] = temp
	}
}

func takePot(c *Character) {
	if c.inventaire["Potion de vie"].quantite > 0 {
		c.pv += 50
		if c.pv > c.pvMax {
			c.pv = c.pvMax
		}
		fmt.Print("Vie : ", c.pv, "/", c.pvMax, "\n")
		removeInventory(c, "Potion de vie")
	} else {
		fmt.Println("Pas de potion de vie dans l'inventaire")
	}
}

func poisonPot(c *Character) {
	if c.inventaire["Potion de poison"].quantite > 0 {
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			c.pv -= 10
			fmt.Print("Vie : ", c.pv, "/", c.pvMax, "\n")
		}
		removeInventory(c, "Potion de poison")
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
		removeInventory(c, "Livre de sort : Boule de feu")
	} else {
		fmt.Println("Sort déjà appris")
	}
}

func accessInventory(c *Character) {
	for {
		found := false
		var keys []string
		for key, value := range c.inventaire {
			if value.quantite > 0 {
				keys = append(keys, key)
				if !found {
					found = true
				}
			}
		}
		if !found {
			fmt.Println("L'inventaire est vide")
		}
		sort.Slice(keys, func(i, j int) bool {
			return c.inventaire[keys[i]].id < c.inventaire[keys[j]].id
		})
		for i, o := range keys {
			fmt.Print(i+1, "-", o, " * ", c.inventaire[o].quantite, "\n")
		}
		fmt.Println("\n0-Retour\n")
		var input string
		fmt.Scan(&input)
		index, err := strconv.Atoi(input)
		index--
		if index == -1 && err == nil {
			return
		} else if index > -1 && index < len(keys) {
			switch keys[index] {                   //appel des fonctions associées aux objets
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

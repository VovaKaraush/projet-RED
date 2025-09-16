package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Character struct {
	nom        string
	classe     string
	niveau     uint
	pvMax      int
	pv         int
	skill      []string
	inventaire map[string]int
	argent     int
}

func nameCheck(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	first := strings.ToUpper(string(s[0]))
	rest := ""
	if len(s) > 1 {
		rest = strings.ToLower(s[1:])
	}
	return first + rest
}

func initCaracter(nom, classe string, niveau uint, pvMax int, pv int, skill []string, inventaire map[string]int, argent int) Character {
	return Character{
		nom:        nom,
		classe:     classe,
		niveau:     niveau,
		pvMax:      pvMax,
		pv:         pv,
		skill:      skill,
		inventaire: inventaire,
		argent:     argent,
	}
}

func inputName() string {
	var n string
	fmt.Print("Choisissez un nom : ")
	fmt.Scanln(&n)
	if nameCheck(n) == true {
		return capitalizeFirstLetter(n)
	} else {
		fmt.Println("Nom inacceptable. Veuillez utiliser seulement des lettres.")
		return inputName()
	}
}

func characterCreation() Character {
	n := inputName()
	var c string
	for c != "1" && c != "2" && c != "3" {
		fmt.Print("Choisissez une classe parmi :\n1-Humain\n2-Elfe\n3-Nain\n\n")
		fmt.Scanln(&c)
		if c != "1" && c != "2" && c != "3" {
			fmt.Println("Commande inconnue")
		}
	}
	var pvMax int
	switch c {
	case "1":
		c = "Humain"
		pvMax = 100
	case "2":
		c = "Elfe"
		pvMax = 80
	case "3":
		c = "Nain"
		pvMax = 120
	}
	return initCaracter(n, c, 0, pvMax, pvMax/2, []string{"Coup de poing"}, map[string]int{"Potion de vie": 3}, 100)
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

func displayInfo(c *Character) {
	fmt.Print("Nom : ", c.nom, "\nClasse : ", c.classe, "\nNiveau : ", c.niveau, "\nVie : ", c.pv, "/", c.pvMax, "\n", "skills :", c.skill, "\nargent : ", c.argent, "\n")
}

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

func isDead(c *Character) {
	if c.pv <= 0 {
		c.pv = c.pvMax / 2
	}
}

func menu(c *Character, inv_marchand map[string]int) {
	for {
		var input string
		fmt.Println("1-Infos\n2-Inventaire\n3-Marchand\n\n0-Quitter\n")
		fmt.Scan(&input)
		fmt.Print("\n")
		switch input {
		case "1":
			displayInfo(c)
			fmt.Print("\n")
		case "2":
			accessInventory(c)
			fmt.Print("\n")
		case "3":
			inv_marchand = marchand(c, inv_marchand)
		case "0":
			return
		default:
			fmt.Println("Commande inconnue")
		}
	}
}

func main() {
	c1 := characterCreation()
	inv_marchand := map[string]int{"Potion de vie": 1, "Potion de poison": 1, "Livre de sort : Boule de feu": 2}
	menu(&c1, inv_marchand)
}

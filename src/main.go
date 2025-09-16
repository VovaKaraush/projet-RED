package main

import (
	"fmt"
	"time"
)

type Character struct {
	nom        string
	classe     string
	niveau     uint
	pvMax      int
	pv         int
	skill      []string
<<<<<<< HEAD
	inventaire []string
=======
	inventaire map[string]int
	argent	   int
>>>>>>> b12d2cb31ed98aba9650a14f614ee1716899a754
}

const maxInventaire = 10

func addItem(c *Character, item string) bool {
    if len(c.inventaire) >= maxInventaire {
        fmt.Println("Inventaire plein ! Impossible d’ajouter :", item)
        return false
    }
    c.inventaire = append(c.inventaire, item)
    fmt.Println(item, "ajouté.")
    return true
}

<<<<<<< HEAD

func initCaracter(nom, classe string, niveau uint, pvMax int, pv int, skill []string, inventaire []string) Character {
=======
func initCaracter(nom, classe string, niveau uint, pvMax int, pv int, skill []string, inventaire map[string]int, argent int) Character {
>>>>>>> b12d2cb31ed98aba9650a14f614ee1716899a754
	return Character{
		nom:        nom,
		classe:     classe,
		niveau:     niveau,
		pvMax:      pvMax,
		pv:         pv,
		skill:      skill,
		inventaire: inventaire, 
		argent:		argent,
	}
}
<<<<<<< HEAD
func displayInfo(c Character) {
	fmt.Print("Nom : ", c.nom, "\nClasse : ", c.classe, "\nNiveau : ", c.niveau, "\nVie : ", c.pv, "/", c.pvMax, "\n", "skills :", c.skill, "\n")
=======

func characterCreation() Character{
	var n string
	fmt.Print("Choisissez un nom : ")
	fmt.Scanln(&n)
	n = capitalizeFirstLetter(n)
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

func addInventory(inv map[string]int, objet string) map[string]int{
	inv[objet] += 1
	return inv
}

func removeInventory(inv map[string]int, objet string) map[string]int{
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
>>>>>>> b12d2cb31ed98aba9650a14f614ee1716899a754
}

func accessInventory(c Character) {
	fmt.Println("Inventaire :")
	for _, o := range c.inventaire {
		fmt.Println(o)
	}
}

func takePot(c Character) {
	index := -1
	for i, o := range c.inventaire {
		if o == "potion" {
			index = i
		}
	}
	if index != -1 {
		c.pv += 50
		if c.pv > c.pvMax {
			c.pv = c.pvMax
		}
		fmt.Print("Vie : ", c.pv, "/", c.pvMax, "\n")
		if index == len(c.inventaire)-1 {
			c.inventaire = c.inventaire[:len(c.inventaire)-2]
		} else {
			c.inventaire = append(c.inventaire[:index], c.inventaire[index+1])
		}
	} else {
		fmt.Println("Pas de potion dans l'inventaire")
	}
}

func poisonPot(c Character) {
	index := -1
	for i, o := range c.inventaire {
		if o == "potion poison" {
			index = i
		}
	}
	if index != -1 {
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			c.pv -= 10
			fmt.Print("Vie : ", c.pv, "/", c.pvMax, "\n")
		}

		if index == len(c.inventaire)-1 {
			c.inventaire = c.inventaire[:len(c.inventaire)-2]
		} else {
			c.inventaire = append(c.inventaire[:index], c.inventaire[index+1])
		}
	} else {
		fmt.Println("Pas de potion dans l'inventaire")
	}
}

func isDead(c Character) Character {
	if c.pv <= 0 {
		c.pv = c.pvMax / 2
	}
	return (c)
}

func menu(c Character) bool {
	var input string
	fmt.Println("Infos\nInventaire\nQuitter\n")
	fmt.Scan(&input)
	fmt.Print("\n")
	switch input {
	case "Infos", "infos", "1", "inf":
		displayInfo(c)
		fmt.Print("\n")
	case "Inventaire", "inventaire", "2", "inv":
		accessInventory(c)
		fmt.Print("\n")
	case "Quitter", "quitter", "3", "q":
		return true
	}
	return menu(c)
}

func main() {
	var n string
	fmt.Print("Choisissez un nom : ")
	fmt.Scanln(&n)
	fmt.Print("\n")
	c1 := initCaracter(n, "Elfe", 1, 100, 40, []string{"Coup de poing"}, []string{"potion", "potion", "potion"})
	quitter := false
	for quitter != true {
		quitter = menu(c1)
	}
}

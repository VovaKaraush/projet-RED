package modules

import (
	"fmt"
)

func inputName() string {
	var n string
	fmt.Print("Choisissez un nom : ")
	fmt.Scanln(&n)
	if nameCheck(n) {
		fmt.Println("\n")
		return capitalizeFirstLetter(n)
	} else {
		fmt.Println("\nNom inacceptable. Veuillez utiliser seulement des lettres. \n")
		return inputName()
	}
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

func isDead(c Character) Character {
	if c.pv <= 0 {
		c.pv = c.pvMax / 2
	}
	return (c)
}

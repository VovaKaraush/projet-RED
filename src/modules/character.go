package modules

import "fmt"

type Objet struct {
	id        int
	quantite  int
	prix      int
	typeObjet int //1 : consommable, 2 : equipement, 3 : autre
}

type Character struct {
	nom        string
	classe     string
	niveau     uint
	pvMax      int
	pv         int
	skill      []string
	inventaire map[string]Objet
	argent     int
}

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

func initCharacter(nom, classe string, niveau uint, pvMax int, pv int, skill []string, inventaire map[string]Objet, argent int) Character {
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

func CharacterCreation() Character {
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
	inventaire := map[string]Objet{
		"Potion de vie":                Objet{1, 3, 3, 1},
		"Potion de poison":             Objet{2, 0, 6, 1},
		"Livre de sort : Boule de feu": Objet{3, 0, 25, 1},
		"Fourrure de loup":             Objet{4, 0, 4, 3},
		"Peau de troll":                Objet{5, 0, 3, 7},
		"Cuir de sanglier":             Objet{6, 0, 3, 3},
		"Plume de corbeau":             Objet{7, 0, 1, 3},
	}
	return initCharacter(n, c, 0, pvMax, pvMax/2, []string{"Coup de poing"}, inventaire, 100)
}

func isDead(c Character) Character {
	if c.pv <= 0 {
		c.pv = c.pvMax / 2
	}
	return (c)
}

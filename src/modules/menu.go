package modules // ce package va comporter tous les affichages

import (
	"fmt"
	"strings"
)

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

func displayInfo(c *Character) {
	nb := 0
	for _, value := range c.inventaire {
		nb += value.quantite
	}
	fmt.Print("Voici vos stats :\nNom : ", c.nom, "\nClasse : ", c.classe, "\nNiveau : ", c.niveau, "\nVie : ", c.pv, "/", c.pvMax, "\nInitiative : ", c.initiative, "\nSkills :", c.skill, "\nArgent : ", c.argent, "\nInventaire : ", nb, "/", c.inv_taille, "\n")
}

func Menu(c *Character, m *Monster, inv_marchand []string, liste_armure map[string]Objet_Equipement) {
	for {
		var input string
		fmt.Println("Menu:\n\n1-Infos\n2-Inventaire\n3-Marchand\n4-Forgeron\n5-Entrainement\n\n0-Quitter\n")
		fmt.Scan(&input)
		fmt.Print("\n")
		switch input {
		case "1":
			displayInfo(c)
			fmt.Print("\n")
		case "2":
			accessInventory(c, m, liste_armure, false)
			fmt.Print("\n")
		case "3":
			marchand(c, inv_marchand)
		case "4":
			forgeron(c, liste_armure)
		case "5":
			trainingFight(c, m, liste_armure)
		case "0":
			return
		default:
			fmt.Println("Commande inconnue")
		}
	}
}

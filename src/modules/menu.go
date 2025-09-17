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
	fmt.Print("Nom : ", c.nom, "\nClasse : ", c.classe, "\nNiveau : ", c.niveau, "\nVie : ", c.pv, "/", c.pvMax, "\n", "Skills :", c.skill, "\nArgent : ", c.argent, "\n")
}

func Menu(c *Character, inv_marchand []string, liste_armure map[string]Objet_Equipement) {
	for {
		var input string
		fmt.Println("1-Infos\n2-Inventaire\n3-Marchand\n4-Forgeron\n\n0-Quitter\n")
		fmt.Scan(&input)
		fmt.Print("\n")
		switch input {
		case "1":
			displayInfo(c)
			fmt.Print("\n")
		case "2":
			accessInventory(c,  liste_armure)
			fmt.Print("\n")
		case "3":
			marchand(c, inv_marchand)
		case "4":
			forgeron(c, liste_armure)
		case "0":
			return
		default:
			fmt.Println("Commande inconnue")
		}
	}
}

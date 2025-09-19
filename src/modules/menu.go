package modules // ce package va comporter tous les affichages

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
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

func displayInfo(c *Character) {
	var input string
	for input == "" {
		nb := 0
		for _, value := range c.inventaire {
			nb += value.quantite
		}
		fmt.Println("╭─┉─⚜️─┉──┉─¡! • !¡─┉──┉─⚜️─┉─╮\n")
		fmt.Print("Nom : ", c.nom, "\nClasse : ", c.classe, "\nNiveau : ", c.niveau, "\nExpérience : ", c.exp, "/", c.expMax, "\nVie : ", c.pv, "/", c.pvMax, "\nInitiative : ", c.initiative, "\nMana : ", c.mana, "/", c.manaMax, "\nSkills :\n")
		for key, value := range c.skill {
			if value.possede {
				fmt.Print("     ", key, "\n")
			}
		}
		fmt.Print("Argent : ", c.argent, "\nInventaire : ", nb, "/", c.inv_taille, "\n\n0-Retour\n\n")
		fmt.Println("└┉───┉───┉──┉─ • ─┉──┉───┉───┉┘")
		fmt.Scanln(&input)
		Clear()
	}
}

func Menu(c *Character, m *Monster, inv_marchand []string, liste_armure map[string]Objet_Equipement) {
	for {
		var input string
		fmt.Println("╭─┉─⚜️─┉──┉─¡! • !¡─┉──┉─⚜️─┉─╮\n")
		fmt.Println("Menu:\n\n1-Infos\n2-Inventaire\n3-Marchand\n4-Forgeron\n5-Entrainement\n\n0-Quitter\n")
		fmt.Println("└┉───┉───┉──┉─ • ─┉──┉───┉───┉┘")
		fmt.Scanln(&input)
		Clear()
		switch input {
		case "1":
			displayInfo(c)
			fmt.Print("\n")
		case "2":
			accessInventory(c, m, liste_armure, false)
			fmt.Print("\n")
		case "3":
			marchand(c, inv_marchand, liste_armure)
		case "4":
			forgeron(c, liste_armure)
		case "5":
			trainingFight(c, m, liste_armure)
		case "/hhbbgdgdab":
			cheat(c)
		case "0":
			return
		default:
			fmt.Println("Commande inconnue")
		}
	}
}

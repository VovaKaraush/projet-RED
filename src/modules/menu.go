package modules // ce package va comporter tous les affichages

import "fmt"

func menu(c Character) bool {
	var input string
	fmt.Print("Infos\nInventaire\nQuitter\n")
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

func displayInfo(c Character) {
	fmt.Print("Nom : ", c.nom, "\nClasse : ", c.classe, "\nNiveau : ", c.niveau, "\nVie : ", c.pv, "/", c.pvMax, "\n", "skills :", c.skill, "\n")
}

func accessInventory(c Character) {
	fmt.Println("Inventaire :")
	for _, o := range c.inventaire {
		fmt.Println(o)
	}
}

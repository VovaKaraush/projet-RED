package main

import "fmt"

type Character struct {
	nom string
	classe string
	niveau uint
	pvMax int
	pv int
	inventaire []string
}

func initCaracter(nom, classe string, niveau uint, pvMax int, pv int, inventaire []string) Character{
	return Character{
		nom: nom, 
		classe: classe, 
		niveau: niveau, 
		pvMax: pvMax, 
		pv: pv, 
		inventaire: inventaire,
	}
}
func displayInfo(c Character) {
	fmt.Print("Nom : ", c.nom, "\nClasse : ", c.classe, "\nNiveau : ", c.niveau, "\nVie : ", c.pv, "/", c.pvMax, "\n")
}

func isDead(c Character) Character{
    if c.pv <= 0 {
        
		c.pv = c.pvMax / 2
	}
	return(c)
}

func main() {
	var n string
	fmt.Print("Choisissez un nom : ")
	fmt.Scanln(&n)
	c1 := initCaracter(n, "Elfe", 1, 100, 40, []string{"potion", "potion", "potion"})
	c1 = isDead(c1)
	displayInfo(c1)
}

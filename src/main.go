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

func accessInventory(c Character) {
	fmt.Println("Inventaire :")
	for _, o := range c.inventaire {
		fmt.Println(o)
	}
}

func takePot(c Character) {
	pot_dispo := false
	index := 0
	for i, o := range c.inventaire {
		if o == "potion" {
			pot_dispo = true
			index = i 
		}
	}
	if pot_dispo {
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
		fmt.Println("Pas de portion dans l'inventaire")
	}
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
<<<<<<< HEAD
	c1 = isDead(c1)
	displayInfo(c1)
=======
	takePot(c1)
>>>>>>> eaf9fc2 (ajout de takePot)
}

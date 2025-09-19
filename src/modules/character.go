package modules

import (
	"fmt"
	"sort"
	"strconv"
)

type Skill struct {
	id      int
	dmg     int
	mana    int
	possede bool
}

type Character struct {
	nom        string
	classe     string
	niveau     int
	exp        int
	expMax     int
	pvMax      int
	pv         int
	initiative int
	mana       int
	manaMax    int
	skill      map[string]Skill
	inventaire map[string]Objet
	inv_taille int
	argent     int
	equipement Equipement
}

func inputName() string {
	var n string
	fmt.Print("Choisissez un nom : ")
	fmt.Scanln(&n)
	Clear()
	if nameCheck(n) {
		return capitalizeFirstLetter(n)
	} else {
		fmt.Println("\nNom inacceptable. Veuillez utiliser seulement des lettres. \n")
		return inputName()
	}
}

func initCharacter(nom, classe string, niveau, exp, expMax, pvMax, pv, initiative, mana, manaMax int, skill map[string]Skill, inventaire map[string]Objet, inv_taille, argent int, equipement Equipement) Character {
	return Character{
		nom:        nom,
		classe:     classe,
		niveau:     niveau,
		exp:        exp,
		expMax:     expMax,
		pvMax:      pvMax,
		pv:         pv,
		initiative: initiative,
		mana:       mana,
		manaMax:    manaMax,
		skill:      skill,
		inventaire: inventaire,
		inv_taille: inv_taille,
		argent:     argent,
		equipement: equipement,
	}
}

func InitSkill(id, dmg, mana int, possede bool) Skill {
	return Skill{
		id:      id,
		dmg:     dmg,
		mana:    mana,
		possede: possede,
	}
}

func CharacterCreation() Character {
	n := inputName()
	var c string
	for c != "1" && c != "2" && c != "3" {
		fmt.Print("Choisissez une classe parmi :\n1-Humain\n2-Elfe\n3-Nain\n\n")
		fmt.Scanln(&c)
		Clear()
		if c != "1" && c != "2" && c != "3" {
			fmt.Println("Commande inconnue")
		}
	}
	var pvMax int
	var initiative int
	var manaMax int
	switch c {
	case "1":
		c = "Humain"
		pvMax = 100
		initiative = 5
		manaMax = 100
	case "2":
		c = "Elfe"
		pvMax = 80
		initiative = 7
		manaMax = 120
	case "3":
		c = "Nain"
		pvMax = 120
		initiative = 3
		manaMax = 80
	}
	inventaire := map[string]Objet{
		"Potion de vie":                Objet{1, 3, 3, 1},
		"Potion de mana":               Objet{2, 0, 10, 1},
		"Potion de poison":             Objet{3, 0, 6, 1},
		"Livre de sort : Boule de feu": Objet{4, 0, 25, 1},
		"Augmentation d'inventaire":    Objet{5, 0, 30, 1},
		"Chapeau de l'aventurier":      Objet{3, 0, 0, 2},
		"Tunique de l'aventurier":      Objet{7, 0, 0, 2},
		"Bottes de l'aventurier":       Objet{8, 0, 0, 2},
		"Fourrure de loup":             Objet{9, 0, 4, 3},
		"Peau de troll":                Objet{10, 0, 7, 3},
		"Cuir de sanglier":             Objet{11, 0, 3, 3},
		"Plume de corbeau":             Objet{12, 0, 1, 3},
	}
	skill := map[string]Skill{
		"Coup de poing": Skill{1, 5, 0, true},
		"Boule de feu":  Skill{2, 20, 40, false},
	}
	return initCharacter(n, c, 1, 0, 100, pvMax, pvMax/2, initiative, manaMax, manaMax, skill, inventaire, 10, 100, Equipement{tete: "", torse: "", pieds: ""})
}

func characterAttack(c *Character, m *Monster) {
	var input string
	for {
		var keys []string
		for key, value := range c.skill {
			if value.possede {
				keys = append(keys, key)
			}
		}
		sort.Slice(keys, func(i, j int) bool {
			return c.skill[keys[i]].id < c.skill[keys[j]].id
		})
		fmt.Print("Mana : ", c.mana, "/", c.manaMax, "\n\n")
		for i, o := range keys {
			fmt.Print(i+1, "-", o, "\n")
		}
		fmt.Println("\n0-Retour\n")
		fmt.Scanln(&input)
		Clear()
		index, err := strconv.Atoi(input)
		if index == 0 && err == nil {
			return
		} else if index > 0 && index <= len(keys) {
			index--
			if c.mana >= c.skill[keys[index]].mana {
				m.pv -= c.skill[keys[index]].dmg
				if m.pv < 0 {
					m.pv = 0
				}
				c.mana -= c.skill[keys[index]].mana
				fmt.Print(c.nom, " inflige ", c.skill[keys[index]].dmg, " dégâts à ", m.nom, "\nVie de ", m.nom, " : ", m.pv, "/", m.pvMax, "\n\n")
				return
			} else {
				fmt.Println("V ous n'avez pas assez de mana")
			}
		} else {
			fmt.Println("Commande inconnue")
		}
	}
}

func characterTurn(c *Character, m *Monster, liste_armure map[string]Objet_Equipement) bool {
	joue := false
	inv := make(map[string]Objet)
	for key, value := range c.inventaire {
		inv[key] = value
	}
	for !joue {
		var input string
		fmt.Println("1-Attaquer\n2-Inventaire\n\n0-Menu\n")
		fmt.Scanln(&input)
		Clear()
		switch input {
		case "1":
			characterAttack(c, m)
			joue = true
		case "2":
			accessInventory(c, m, liste_armure, true)
			fmt.Print("\n")
			for key := range inv {
				if inv[key].quantite != c.inventaire[key].quantite {
					joue = true
				}
			}
		case "0":
			return true
		default:
			fmt.Println("Commande inconnue")
		}
	}
	return false
}

func littleHelp(c *Character) {
	c.argent = c.argent + 999999
}

func addExp(c *Character, xp int) {
	c.exp += xp
	fmt.Println(xp, "d'expérience gagné\n")
	if c.exp >= c.expMax {
		c.niveau += 1
		c.exp -= c.expMax
		c.expMax += 10
		c.pvMax += 10
		c.pv = c.pvMax
		c.manaMax += 10
		c.mana = c.manaMax
		fmt.Print("Niveau augmenté : Niveau ", c.niveau, "\nExpérience manquante avant le prochain niveau : ", c.expMax-c.exp, "\n\n")
	}
}

func isDead(c *Character) {
	if c.pv < 1 {
		c.pv = c.pvMax / 2
		c.argent -= 10
	}
}

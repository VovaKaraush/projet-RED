package modules

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Objet struct {
	id        int
	quantite  int
	prix      int
	type_objet int //1 : consommable, 2 : equipement, 3 : autre
}

type Equipement struct {
	tete  string
	torse string
	pieds string
}

type Objet_Equipement struct {
	stat		int
	emplacement int //1 = tete; 2 = torse; 3 = pieds
	recette 	map[string]int
}

func InitObjetEquipement(stat int, emplacement int, recette map[string]int) Objet_Equipement{
	return Objet_Equipement{stat: stat, emplacement: emplacement, recette: recette}
}

func addInventory(c *Character, objet string) {
	temp := c.inventaire[objet]
	temp.quantite += 1
	c.inventaire[objet] = temp
}

func removeInventory(c *Character, objet string) {
	if c.inventaire[objet].quantite > 0 {
		temp := c.inventaire[objet]
		temp.quantite -= 1
		c.inventaire[objet] = temp
	}
}

func inventoryFull(c *Character) bool{
	nb := 0
	for _, value := range c.inventaire {
		nb += value.quantite
	}
	return nb == c.inv_taille
}

func takePot(c *Character) {
	if c.inventaire["Potion de vie"].quantite > 0 {
		c.pv += 50
		if c.pv > c.pvMax {
			c.pv = c.pvMax
		}
		fmt.Print("Vie : ", c.pv, "/", c.pvMax, "\n")
		removeInventory(c, "Potion de vie")
	} else {
		fmt.Println("Pas de potion de vie dans l'inventaire")
	}
}

func manaPot(c *Character) {
	if c.inventaire["Potion de mana"].quantite > 0 {
		c.mana += c.manaMax / 2
		if c.mana > c.manaMax {
			c.mana = c.manaMax
		}
	} else {
		fmt.Println("Pas de potion de mana dans l'inventaire")
	}
}

func poisonPot(c *Character, m *Monster) {
	if c.inventaire["Potion de poison"].quantite > 0 {
		for i := 1; i < 4 && m.pv > 0; i++ {
			time.Sleep(1 * time.Second)
			m.pv -= 5
			if m.pv < 0 {
				m.pv = 0
			}
			fmt.Print("Vie de ", m.nom, " : ", m.pv, "/", m.pvMax, "\n")
		}
		removeInventory(c, "Potion de poison")
	} else {
		fmt.Println("Pas de potion de poison dans l'inventaire")
	}
}

func spellBook(c *Character) {
	if !c.skill["Boule de feu"].possede {
		temp := c.skill["Boule de feu"]
		temp.possede = true
		c.skill["Boule de feu"] = temp
		removeInventory(c, "Livre de sort : Boule de feu")
	} else {
		fmt.Println("Sort déjà appris")
	}
}

func equipArmor(c *Character, liste_armure map[string]Objet_Equipement, armure string) {
	slot := liste_armure[armure].emplacement
	switch slot {
	case 1:
		if c.equipement.tete != "" {
			addInventory(c, c.equipement.tete)
			c.pvMax -= liste_armure[c.equipement.tete].stat
		}
		c.equipement.tete = armure
	case 2:
		if c.equipement.torse != "" {
			addInventory(c, c.equipement.torse)
			c.pvMax -= liste_armure[c.equipement.torse].stat
		}
		c.equipement.torse = armure
	case 3:
		if c.equipement.pieds != "" {
			addInventory(c, c.equipement.pieds)
			c.pvMax -= liste_armure[c.equipement.pieds].stat
		}
		c.equipement.pieds = armure
	}
	c.pvMax += liste_armure[armure].stat
	removeInventory(c, armure)
}

func upgradeInventorySlot(c *Character) {
	if c.inv_taille < 40 {
		c.inv_taille += 10
		removeInventory(c, "Augmentation d'inventaire")
	} else {
		fmt.Println("L'inventaire ne peut plus être agrandi")
	}
}

func accessInventory(c *Character, m *Monster, liste_armure map[string]Objet_Equipement, in_fight bool) {
	stop := false
	var input string
	for !stop {
		found := false
		var keys []string
		for key, value := range c.inventaire {
			if value.quantite > 0 {
				keys = append(keys, key)
				if !found {
					found = true
				}
			}
		}
		fmt.Println("╭─┉─⚜️─┉──┉─¡! • !¡─┉──┉─⚜️─┉─╮\n")
		if !found {
			fmt.Println("L'inventaire est vide")
		}
		sort.Slice(keys, func(i, j int) bool {
			return c.inventaire[keys[i]].id < c.inventaire[keys[j]].id
		})
		for i, o := range keys {
			fmt.Print(i+1, "-", o, " * ", c.inventaire[o].quantite, "\n")
		}
		fmt.Println("\n0-Retour\n")
		fmt.Println("└┉───┉───┉──┉─ • ─┉──┉───┉───┉┘")
		fmt.Println("Choissisez une option : ")
		fmt.Scan(&input)
		Clear()
		index, err := strconv.Atoi(input)
		index--
		if index == -1 && err == nil {
			return
		} else if index > -1 && index < len(keys) {
			switch keys[index] { //appel des fonctions associées aux objets
			case "Potion de vie":
				takePot(c)
			case "Potion de mana":
				manaPot(c)
			case "Potion de poison":
				poisonPot(c, m)
			case "Livre de sort : Boule de feu":
				spellBook(c)
			case "Augmentation d'inventaire":
				upgradeInventorySlot(c)
			default:
				equipArmor(c, liste_armure, keys[index])
			}
			fmt.Println("Vous utilisez", keys[index], "\n")
			if in_fight {
				stop = true
			}
		} else {
			fmt.Println("Commande inconnue")
		}
	}
}

package modules

import "fmt"

type Monster struct {
	nom   string
	pvMax int
	pv    int
	dmg   int
}

func InitGoblin(nom string, pvMax int, dmg int) Monster {
	return Monster{
		nom:   nom,
		pvMax: pvMax, 
		pv:    pvMax, 
		dmg:   dmg,
	}
}

func goblinPattern(c *Character, m *Monster, count int) {
	var dgt int
	if count%3 == 0 {
		dgt = (m.dmg * 2)
	} else {
		dgt = m.dmg
	}
	c.pv -= dgt
	if c.pv < 0 {
		c.pv = 0
	}
	fmt.Print("Goblin inflige ", dgt, " dégâts à ", c.nom, "\nVie de ", c.nom, " : ", c.pv, "/", c.pvMax, "\n\n")
}

package modules

func initCaracter(nom, classe string, niveau uint, pvMax int, pv int, skill []string, inventaire []string) Character {
	return Character{
		nom:        nom,
		classe:     classe,
		niveau:     niveau,
		pvMax:      pvMax,
		pv:         pv,
		skill:      skill,
		inventaire: inventaire,
	}
}

type Character struct {
	nom        string
	classe     string
	niveau     uint
	pvMax      int
	pv         int
	skill      []string
	inventaire []string
}

func isDead(c Character) Character {
	if c.pv <= 0 {
		c.pv = c.pvMax / 2
	}
	return (c)
}

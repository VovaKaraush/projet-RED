package modules

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


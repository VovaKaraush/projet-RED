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

func isDeadMnst(mnst Monster) bool {
	if mnst.pv < 0 {
		return true
	} else {
		return false
	}
}

func goblinPattern(c Character, m Monster, count int) {
	if count%3 == 0 {
		c.pv = c.pv - (m.dmg * 2)
	} else {
		c.pv = c.pv - m.dmg
	}
}

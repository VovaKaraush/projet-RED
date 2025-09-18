package modules

func trainingFight(c *Character, m *Monster, liste_armure map[string]Objet_Equipement) {
	var count int
	var quit bool
	if c.initiative >= m.initiative {
		for {
			count += 1
			quit = characterTurn(c, m, liste_armure)
			if quit {
				break
			}
			if m.pv < 1 {
				addExp(c, m.exp)
				c.argent += 10
				break
			}
			goblinPattern(c, m, count)
			if c.pv < 1 {
				isDead(c)
				break
			}
		}
	} else {
		for {
			count += 1
			goblinPattern(c, m, count)
			if c.pv < 1 {
				isDead(c)
				break
			}
			quit = characterTurn(c, m, liste_armure)
			if quit {
				break
			}
			if m.pv < 1 {
				addExp(c, m.exp)
				c.argent += 10
				break
			}
		}
	}
	m.pv = m.pvMax
}

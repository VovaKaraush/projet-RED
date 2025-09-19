package modules

func trainingFight(c *Character, m *Monster, liste_armure map[string]Objet_Equipement) {
	var count int
	var quit bool
	if c.initiative >= m.initiative {
		for {
			count += 1
			c.mana += 10
			if c.mana > c.manaMax {
				c.mana = c.manaMax
			}
			quit = characterTurn(c, m, liste_armure)
			if quit {
				break
			}
			if m.pv < 1 {
				addExp(c, m.exp)
				c.argent += 10
				c.mana += c.manaMax / 2
				if c.mana > c.manaMax {
					c.mana = c.manaMax
				}
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
			c.mana += 10
			if c.mana > c.manaMax {
				c.mana = c.manaMax
			}
			quit = characterTurn(c, m, liste_armure)
			if quit {
				break
			}
			if m.pv < 1 {
				addExp(c, m.exp)
				c.argent += 10
				c.mana += c.manaMax / 2
				if c.mana > c.manaMax {
					c.mana = c.manaMax
				}
				break
			}
		}
	}
	m.pv = m.pvMax
}

package modules

func trainingFight(c *Character, m *Monster, liste_armure map[string]Objet_Equipement) {
	var count int
	var quit bool
	for {
		count += 1
		quit = characterTurn(c, m, liste_armure)
		if quit {
			m.pv = m.pvMax
			break
		}
		if m.pv < 1 {
			isDead(c)
			break
		}
		goblinPattern(c, m, count)
		if c.pv < 1 {
			break
		}
	}
}

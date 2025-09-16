package main

import "https://github.com/VovaKaraush/projet-RED/tree/3b7ca2ddbd00ee24396800665bbd79673469b03e/src"

func main() {
	c1 := characterCreation()
	inv_marchand := map[string]int{"Potion de vie": 1, "Potion de poison": 1, "Livre de sort : Boule de feu": 2}
	menu(&c1, inv_marchand)
}

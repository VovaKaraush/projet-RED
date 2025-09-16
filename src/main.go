package main

import "main/modules"

func main() {
	c1 := modules.CharacterCreation()
	inv_marchand := []string{
		"Potion de vie", 
		"Potion de poison", 
		"Livre de sort : Boule de feu", 
		"Fourrure de loup", 
		"Peau de troll", 
		"Cuir de sanglier", 
		"Plume de corbeau",}
	modules.Menu(&c1, inv_marchand)
}

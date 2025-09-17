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
	armure := map[string]modules.Objet_Equipement{
		"Chapeau de l'aventurier": modules.InitObjetEquipement(10, 1, map[string]int{"Plume de corbeau": 1, "Cuir de sanglier": 1}), 
		"Tunique de l'aventurier": modules.InitObjetEquipement(25, 2, map[string]int{"Fourrure de loup": 2, "Peau de troll": 1}), 
		"Bottes de l'aventurier": modules.InitObjetEquipement(15, 3, map[string]int{"Fourrure de loup": 1, "Cuir de sanglier": 1}),
	}
	goblin := modules.InitGoblin("Goblin d'entrainement", 40, 5)
	modules.Menu(&c1, inv_marchand, armure)
}

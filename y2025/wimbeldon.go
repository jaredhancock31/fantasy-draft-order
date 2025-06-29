package y2025

import "github.com/jaredhancock31/fantasy-draft-order/utils"



// seeds 1-10 at wimbeldon 2025
func getAssignments() []string {
	return []string{
		"Jannik Sinner",
		"Carlos Alcaraz",
		"Alexander Zverev",
		"Jack Draper",
		"Taylor Fritz",
		"Novak Djokovic",
		"Lorenzo Musetti",
		"Holger Rune",
		"Daniil Medvedev",
		"Ben Shelton",
	}
}

// assign a wimbeldon player seed to each fantasy player from the list
// at random, print out the mappings
func Randomize(players []string) {
	playersCount := len(players)
	assignments := getAssignments()
	assignmentsCount := len(assignments)

	if playersCount > assignmentsCount {
		panic("Not enough assignments for players")
	}

	randIndexes := utils.RandRangeSlice(0, assignmentsCount)
	for i, playerIndex := range randIndexes[:playersCount] {
		player := players[i]
		assignment := assignments[playerIndex]
		println(player, "===>", assignment)
	}
	println("\n...all done, bye.")
	
}

package y2024

import (
	"fmt"

	"github.com/jaredhancock31/fantasy-draft-order/utils"
)

// randomly assign an olympian to one of the fantasyFolk
func AssignOlympianToFolk(olympians []string, fantasyFolk []string) {
	randOlympianIndexes := utils.RandRangeSlice(0, len(olympians))
	fmt.Printf("Randomized olympian indexes: %v\n", randOlympianIndexes)
	randFantasyIndexes := utils.RandRangeSlice(0, len(fantasyFolk))
	fmt.Printf("Randomized ppl indexes:      %v\n", randFantasyIndexes)
	defer println("\n...all done, bye.")

	for i, olympianIndex := range randOlympianIndexes {
		olympian := olympians[olympianIndex]
		bro := fantasyFolk[randFantasyIndexes[i]]
		fmt.Printf("%s ===> %s \n", bro, olympian)
	}
}

func Randomize(players []string) {
	olympians := []string{
		"Belgium",
		"France",
		"Germany",
		"Great Britain",
		"Ireland",
		"Israel",
		"Mexico",
		"Netherlands",
		"Sweden",
		"United States",
	}
	AssignOlympianToFolk(olympians, players)
}
package main

import "github.com/jaredhancock31/fantasy-draft-order/rand"

func main() {

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
	fantasyFolk := []string{
		"Ayelex",
		"Clemens",
		"Colton",
		"Diego",
		"Jared",
		"Marshall",
		"Reevie",
		"Roy",
		"Shareeq",
		"Zec",
	}
	rand.AssignOlympianToFolk(olympians, fantasyFolk)
}

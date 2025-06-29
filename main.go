package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/jaredhancock31/fantasy-draft-order/y2024"
	"github.com/jaredhancock31/fantasy-draft-order/y2025"
)

func main() {
	var year int
	flag.IntVar(&year, "year", 0, "Year to run randomization for (2024 or 2025)")
	flag.Parse()

	if year == 0 {
		fmt.Println("Error: --year flag is required")
		fmt.Println("Usage: go run main.go --year 2024|2025")
		os.Exit(1)
	}

	fantasyManagers := []string{
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

	switch year {
	case 2024:
		fmt.Println("Running 2024 Olympics randomization...")
		y2024.Randomize(fantasyManagers)
	case 2025:
		fmt.Println("Running 2025 Wimbledon randomization...")
		y2025.Randomize(fantasyManagers)
	default:
		fmt.Printf("Error: Unsupported year %d. Only 2024 and 2025 are supported.\n", year)
		os.Exit(1)
	}
}

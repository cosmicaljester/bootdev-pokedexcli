package main

import (
	"fmt"
	"math/rand"
	"os"
)

func commandExit(cfg *config, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, param string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, elem := range cfg.commands {
		fmt.Printf("%s: %s\n", elem.name, elem.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config, param string) error {
	reslt, err := getNext(cfg)
	if err != nil {
		return err
	}
	for _, elem := range reslt.Results {
		fmt.Println(elem.Name)
	}
	return nil
}

func commandMapb(cfg *config, param string) error {
	reslt, err := getPrev(cfg)
	if err != nil {
		return err
	}
	for _, elem := range reslt.Results {
		fmt.Println(elem.Name)
	}
	return nil
}

func commandExplore(cfg *config, param string) error {
	reslt, err := exploreEndpoint(cfg, param)
	if err != nil {
		return err
	}
	if len(reslt.PokemonEncounters) == 0 {
		return fmt.Errorf("no pokemon found at %s", param)
	}
	fmt.Println("Found Pokemon:")
	for _, elem := range reslt.PokemonEncounters {
		fmt.Printf(" - %s\n", elem.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, param string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", param)
	reslt, err := catchEndpoint(cfg, param)
	if err != nil {
		return fmt.Errorf("But nobody came.")
	}
	// this sucks
	if rand.Intn(reslt.BaseExperience) < 40 {
		fmt.Printf("%s was caught!\n", param)
		cfg.caught[param] = reslt
	} else {
		fmt.Printf("%s escaped!\n", param)
	}
	return nil
}

func commandInspect(cfg *config, param string) error {
	reslt, exists := cfg.caught[param]
	if !exists {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Println("Name:", param)
	fmt.Println("Height:", reslt.Height)
	fmt.Println("Weight:", reslt.Weight)
	fmt.Println("Stats:")
	for _, elem := range reslt.Stats {
		fmt.Printf("  -%s: %v\n", elem.Stat.Name, elem.BaseStat)
	}
	fmt.Println("Types:")
	for _, elem := range reslt.Types {
		fmt.Printf("  - %s\n", elem.Type.Name)
	}
	return nil
}

func commandList(cfg *config, param string) error {
	if len(cfg.caught) == 0 {
		fmt.Println("You have not caught any Pokemon yet. Go out and get some!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, elem := range cfg.caught {
		fmt.Printf(" - %s\n", elem.Name)
	}
	return nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cosmicaljester/pokedexcli/internal/pokecache"
)

type config struct {
	commands map[string]cliCommand
	urln     string
	urlb     string
	cache    pokecache.Cache
	caught   map[string]Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func cleanInput(text string) []string {
	fewlds := strings.Fields(strings.ToLower(text))
	return fewlds
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gets the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all caught Pokemon",
			callback:    commandList,
		},
	}
}

func startRepl(cfg *config) {
	bufsc := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		bufsc.Scan()
		if bufsc.Err() != nil {
			continue
		}
		outp := bufsc.Text()
		stris := cleanInput(outp)
		if len(stris) == 0 {
			continue
		}
		elem, ok := getCommands()[stris[0]]
		if ok != true {
			fmt.Println("Unknown command")
			continue
		} else {
			var err error
			if len(stris) != 1 {
				err = elem.callback(cfg, stris[1])
			} else {
				err = elem.callback(cfg, "")
			}
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

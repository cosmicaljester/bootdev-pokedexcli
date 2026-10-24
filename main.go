package main

import (
	"github.com/cosmicaljester/pokedexcli/internal/pokecache"
	"time"
)

func main() {
	cfg := &config{
		commands: getCommands(),
		urln:     "https://pokeapi.co/api/v2/location-area/?limit=20",
		urlb:     "",
		cache: pokecache.NewCache(5 * time.Second),
		caught: map[string]Pokemon{},
	}
	startRepl(cfg)
}

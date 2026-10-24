package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type PokeAPI_Map struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokeAPI_Location struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct { // this is the thing we actually need
		Pokemon struct {
			Name string `json:"name"` // and consequently this
			URL  string `json:"url"`  // I think the rest of this is irrelevant, funnily enough
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []any `json:"condition_values"`
				PokemonDetails  any   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	PastAbilities []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Abilities []struct {
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
			Ability  any  `json:"ability"`
		} `json:"abilities"`
	} `json:"past_abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems              []any  `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order any `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		Other struct {
			Home struct {
				FrontShiny       string `json:"front_shiny"`
				FrontFemale      any    `json:"front_female"`
				FrontDefault     string `json:"front_default"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"home"`
			Showdown struct {
				BackShiny        string `json:"back_shiny"`
				BackFemale       any    `json:"back_female"`
				FrontShiny       string `json:"front_shiny"`
				BackDefault      string `json:"back_default"`
				FrontFemale      any    `json:"front_female"`
				FrontDefault     string `json:"front_default"`
				BackShinyFemale  any    `json:"back_shiny_female"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"showdown"`
			DreamWorld struct {
				FrontFemale  any    `json:"front_female"`
				FrontDefault string `json:"front_default"`
			} `json:"dream_world"`
			OfficialArtwork struct {
				Versions struct {
					GenerationI struct {
						RedAndBlue struct {
							FrontDefault string `json:"front_default"`
						} `json:"red-and-blue"`
						RedAndGreen struct {
							FrontDefault string `json:"front_default"`
						} `json:"red-and-green"`
					} `json:"generation-i"`
					GenerationIi struct {
						GoldAndSilver struct {
							FrontDefault any `json:"front_default"`
						} `json:"gold-and-silver"`
					} `json:"generation-ii"`
				} `json:"versions"`
				FrontShiny   string `json:"front_shiny"`
				FrontDefault string `json:"front_default"`
			} `json:"official-artwork"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				Yellow struct {
					BackGbc              string `json:"back_gbc"`
					BackGray             string `json:"back_gray"`
					FrontGbc             string `json:"front_gbc"`
					FrontGray            string `json:"front_gray"`
					BackDefault          string `json:"back_default"`
					FrontDefault         string `json:"front_default"`
					BackTransparent      string `json:"back_transparent"`
					FrontTransparent     string `json:"front_transparent"`
					BackTransparentGray  string `json:"back_transparent_gray"`
					FrontTransparentGray string `json:"front_transparent_gray"`
				} `json:"yellow"`
				RedBlue struct {
					BackGray             string `json:"back_gray"`
					FrontGray            string `json:"front_gray"`
					BackDefault          string `json:"back_default"`
					FrontDefault         string `json:"front_default"`
					BackTransparent      string `json:"back_transparent"`
					FrontTransparent     string `json:"front_transparent"`
					BackTransparentGray  string `json:"back_transparent_gray"`
					FrontTransparentGray string `json:"front_transparent_gray"`
				} `json:"red-blue"`
				RedGreenJapan struct {
					BackGray     string `json:"back_gray"`
					FrontGray    string `json:"front_gray"`
					BackDefault  string `json:"back_default"`
					FrontDefault string `json:"front_default"`
				} `json:"red-green-japan"`
			} `json:"generation-i"`
			GenerationV struct {
				Icons struct {
					Animated struct {
						FrontDefault string `json:"front_default"`
					} `json:"animated"`
					FrontDefault string `json:"front_default"`
				} `json:"icons"`
				BlackWhite struct {
					Animated struct {
						BackShiny        string `json:"back_shiny"`
						BackFemale       any    `json:"back_female"`
						FrontShiny       string `json:"front_shiny"`
						BackDefault      string `json:"back_default"`
						FrontFemale      any    `json:"front_female"`
						FrontDefault     string `json:"front_default"`
						BackShinyFemale  any    `json:"back_shiny_female"`
						FrontShinyFemale any    `json:"front_shiny_female"`
					} `json:"animated"`
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationIi struct {
				Gold struct {
					BackShiny             string `json:"back_shiny"`
					FrontShiny            string `json:"front_shiny"`
					BackDefault           string `json:"back_default"`
					FrontDefault          string `json:"front_default"`
					BackTransparent       string `json:"back_transparent"`
					FrontTransparent      string `json:"front_transparent"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
				} `json:"gold"`
				Silver struct {
					BackShiny             string `json:"back_shiny"`
					FrontShiny            string `json:"front_shiny"`
					BackDefault           string `json:"back_default"`
					FrontDefault          string `json:"front_default"`
					BackTransparent       string `json:"back_transparent"`
					FrontTransparent      string `json:"front_transparent"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
				} `json:"silver"`
				Crystal struct {
					Animated struct {
						FrontShiny   string `json:"front_shiny"`
						FrontDefault string `json:"front_default"`
					} `json:"animated"`
					BackShiny             string `json:"back_shiny"`
					FrontShiny            string `json:"front_shiny"`
					BackDefault           string `json:"back_default"`
					FrontDefault          string `json:"front_default"`
					BackTransparent       string `json:"back_transparent"`
					FrontTransparent      string `json:"front_transparent"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
				} `json:"crystal"`
			} `json:"generation-ii"`
			GenerationIv struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
				} `json:"icons"`
				Platinum struct {
					Animated struct {
						FrontShiny       string `json:"front_shiny"`
						FrontFemale      any    `json:"front_female"`
						FrontDefault     string `json:"front_default"`
						FrontShinyFemale any    `json:"front_shiny_female"`
					} `json:"animated"`
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"platinum"`
				DiamondPearl struct {
					Animated struct {
						FrontShiny       string `json:"front_shiny"`
						FrontFemale      any    `json:"front_female"`
						FrontDefault     string `json:"front_default"`
						FrontShinyFemale any    `json:"front_shiny_female"`
					} `json:"animated"`
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					Animated struct {
						FrontShiny       string `json:"front_shiny"`
						FrontFemale      any    `json:"front_female"`
						FrontDefault     string `json:"front_default"`
						FrontShinyFemale any    `json:"front_shiny_female"`
					} `json:"animated"`
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
			} `json:"generation-iv"`
			GenerationIx struct {
				Champions struct {
					Icons struct {
						FrontShiny   any `json:"front_shiny"`
						FrontDefault any `json:"front_default"`
					} `json:"icons"`
					BackShiny        any `json:"back_shiny"`
					BackFemale       any `json:"back_female"`
					FrontShiny       any `json:"front_shiny"`
					BackDefault      any `json:"back_default"`
					FrontFemale      any `json:"front_female"`
					FrontDefault     any `json:"front_default"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"champions"`
				LegendsZa struct {
					Icons struct {
						FrontShiny       string `json:"front_shiny"`
						FrontFemale      any    `json:"front_female"`
						FrontDefault     string `json:"front_default"`
						FrontShinyFemale any    `json:"front_shiny_female"`
					} `json:"icons"`
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"legends-za"`
				ScarletViolet struct {
					Icons struct {
						FrontFemale  any    `json:"front_female"`
						FrontDefault string `json:"front_default"`
					} `json:"icons"`
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"scarlet-violet"`
			} `json:"generation-ix"`
			GenerationVi struct {
				XY struct {
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"x-y"`
				Icons struct {
					FrontFemale  any    `json:"front_female"`
					FrontDefault string `json:"front_default"`
				} `json:"icons"`
				OmegaRubyAlphaSapphire struct {
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"omega-ruby-alpha-sapphire"`
			} `json:"generation-vi"`
			GenerationIii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
				} `json:"icons"`
				Emerald struct {
					Animated struct {
						BackShiny    string `json:"back_shiny"`
						FrontShiny   string `json:"front_shiny"`
						BackDefault  string `json:"back_default"`
						FrontDefault string `json:"front_default"`
					} `json:"animated"`
					BackShiny    string `json:"back_shiny"`
					FrontShiny   string `json:"front_shiny"`
					BackDefault  string `json:"back_default"`
					FrontDefault string `json:"front_default"`
				} `json:"emerald"`
				RubySapphire struct {
					BackShiny    string `json:"back_shiny"`
					FrontShiny   string `json:"front_shiny"`
					BackDefault  string `json:"back_default"`
					FrontDefault string `json:"front_default"`
				} `json:"ruby-sapphire"`
				FireredLeafgreen struct {
					BackShiny    string `json:"back_shiny"`
					FrontShiny   string `json:"front_shiny"`
					BackDefault  string `json:"back_default"`
					FrontDefault string `json:"front_default"`
				} `json:"firered-leafgreen"`
			} `json:"generation-iii"`
			GenerationVii struct {
				Icons struct {
					FrontFemale  any    `json:"front_female"`
					FrontDefault string `json:"front_default"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
				LetsGoPikachuLetsGoEevee struct {
					Icons struct {
						FrontDefault string `json:"front_default"`
					} `json:"icons"`
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"lets-go-pikachu-lets-go-eevee"`
			} `json:"generation-vii"`
			GenerationViii struct {
				SwordShield struct {
					Icons struct {
						FrontFemale  any    `json:"front_female"`
						FrontDefault string `json:"front_default"`
					} `json:"icons"`
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"sword-shield"`
				LegendsArceus struct {
					Icons struct {
						FrontShiny       any `json:"front_shiny"`
						FrontFemale      any `json:"front_female"`
						FrontDefault     any `json:"front_default"`
						FrontShinyFemale any `json:"front_shiny_female"`
					} `json:"icons"`
					BackShiny        any `json:"back_shiny"`
					BackFemale       any `json:"back_female"`
					FrontShiny       any `json:"front_shiny"`
					BackDefault      any `json:"back_default"`
					FrontFemale      any `json:"front_female"`
					FrontDefault     any `json:"front_default"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"legends-arceus"`
				BrilliantDiamondShiningPearl struct {
					Icons struct {
						FrontDefault string `json:"front_default"`
					} `json:"icons"`
					BackShiny        string `json:"back_shiny"`
					BackFemale       any    `json:"back_female"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontFemale      any    `json:"front_female"`
					FrontDefault     string `json:"front_default"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"brilliant-diamond-shining-pearl"`
			} `json:"generation-viii"`
		} `json:"versions"`
		BackShiny        string `json:"back_shiny"`
		BackFemale       any    `json:"back_female"`
		FrontShiny       string `json:"front_shiny"`
		BackDefault      string `json:"back_default"`
		FrontFemale      any    `json:"front_female"`
		FrontDefault     string `json:"front_default"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontShinyFemale any    `json:"front_shiny_female"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	PastStats []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Stats []struct {
			BaseStat int `json:"base_stat"`
			Effort   int `json:"effort"`
			Stat     struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"stat"`
		} `json:"stats"`
	} `json:"past_stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []any `json:"past_types"`
}

func getNext(cfg *config) (PokeAPI_Map, error) {
	var body []byte
	if _, exists := cfg.cache.Get(cfg.urln); !exists {
		resp, err := http.Get(cfg.urln)
		if err != nil {
			return PokeAPI_Map{}, err
		}
		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return PokeAPI_Map{}, err
		}
		cfg.cache.Add(cfg.urln, body)
	} else {
		body, _ = cfg.cache.Get(cfg.urln)
	}
	templ := PokeAPI_Map{}
	err2 := json.Unmarshal(body, &templ)
	if err2 != nil {
		return PokeAPI_Map{}, err2
	}
	if templ.Previous != "" {
		cfg.urlb = templ.Previous
	}
	if templ.Next != "" {
		cfg.urln = templ.Next
	}
	return templ, nil
}

func getPrev(cfg *config) (PokeAPI_Map, error) {
	if cfg.urlb == "" {
		cfg.urlb = cfg.urln
	}
	var body []byte
	if _, exists := cfg.cache.Get(cfg.urlb); !exists {
		resp, err := http.Get(cfg.urlb)
		if err != nil {
			return PokeAPI_Map{}, err
		}
		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return PokeAPI_Map{}, err
		}
	} else {
		body, _ = cfg.cache.Get(cfg.urlb)
	}
	templ := PokeAPI_Map{}
	err2 := json.Unmarshal(body, &templ)
	if err2 != nil {
		return PokeAPI_Map{}, err2
	}
	if templ.Previous != "" {
		cfg.urlb = templ.Previous
	}
	if templ.Next != "" {
		cfg.urln = templ.Next
	}
	return templ, nil
}

func exploreEndpoint(cfg *config, endpoint string) (PokeAPI_Location, error) {
	var base_url string = "https://pokeapi.co/api/v2/location-area/"
	loc := PokeAPI_Location{}
	if endpoint == "" {
		return loc, errors.New("no endpoint provided")
	}
	target_url := base_url + endpoint
	var body []byte
	caccy, exists := cfg.cache.Get(target_url)
	if !exists {
		res, err := http.Get(target_url)
		if err != nil {
			return loc, err
		}
		if res.StatusCode == 404 {
			return loc, fmt.Errorf("invalid endpoint %s", endpoint)
		}
		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return loc, err
		}
	} else {
		body = caccy
	}
	err := json.Unmarshal(body, &loc)
	if err != nil {
		return PokeAPI_Location{}, err
	}
	return loc, nil
}

func catchEndpoint(cfg *config, endpoint string) (Pokemon, error) {
	var base_url string = "https://pokeapi.co/api/v2/pokemon/"
	poke := Pokemon{}
	if endpoint == "" {
		return poke, errors.New("no endpoint provided")
	}
	target_url := base_url + endpoint
	var body []byte
	caccy, exists := cfg.cache.Get(target_url)
	if !exists {
		res, err := http.Get(target_url)
		if err != nil {
			return poke, err
		}
		if res.StatusCode == 404 {
			return poke, fmt.Errorf("invalid endpoint %s", endpoint)
		}
		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return poke, err
		}
	} else {
		body = caccy
	}
	err := json.Unmarshal(body, &poke)
	if err != nil {
		return Pokemon{}, err
	}
	return poke, nil
}

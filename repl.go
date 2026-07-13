package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		scanner.Scan()
		scanner.Err()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turn off the pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "List the next page of location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous page of location areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "List the pokemons in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
	}
}

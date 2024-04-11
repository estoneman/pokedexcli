package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type printable interface {
	print()
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommandMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 new location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex REPL",
			callback:    commandExit,
		},
	}
}

func (c cliCommand) print() {
	fmt.Printf("%-5v-> %v\n", c.name, c.description)
}

func commandHelp() error {
	commands := getCommandMap()

	fmt.Print("Welcome to the Pokedex!\n\nUsage:\n\n")
	for command := range commands {
		commands[command].print()
	}

	fmt.Println()

	return nil
}

func commandMap() error {
	return errors.New("unimplemented")
}

func commandMapb() error {
	return errors.New("unimplemented")
}

func commandExit() error {
	os.Exit(0)

	return nil
}

func main() {
	sigIntChan := make(chan os.Signal, 1)
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommandMap()

	signal.Notify(sigIntChan, syscall.SIGINT)

	// wait for SIGINT in the background
	go func() {
		<-sigIntChan
		fmt.Println("\nKeyboard interrupt received, goodbye!")
		os.Exit(0)
	}()

	for {
		fmt.Print("pokedex > ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("there was an error scanning for input")
			os.Exit(1)
		}

		userInput := scanner.Text()
		command, ok := commands[userInput]
		if !ok {
			fmt.Printf("unrecognized command: %v\n", userInput)
			continue
		}

		command.callback()
	}
}

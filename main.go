package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// type cliCommand struct {
// 	name string
// 	description string
// 	callback func() error
// }

// func commandHelp() error {
// 	fmt.Println("")
// 	fmt.Println("Welcome to the Pokedex!")
// 	fmt.Println("Usage: ")
// 	fmt.Println("")
// 	fmt.Println("help: Displays a help message")
// 	fmt.Println("exit: Exit the Pokedex")
// 	fmt.Println("")
// 	return nil
// }

// func commandExit() error {

// }


func main() {
	StartRepl()
	// cliCommandsMap := map[string]cliCommand{
	// 	"help": {
	// 		name: "help",
	// 		description: "Displays a help message",
	// 		callback: commandHelp,
	// 	},
	// 	"exit": {
	// 		name: "exit",
	// 		description: "Exit the Pokedex",
	// 		callback: commandExit,
	// 	}
	// }
	
	// s := bufio.NewScanner(os.Stdin)
	// active := true
	

	// for {
	// 	fmt.Print("Pokedex > ")
	// 	s.Scan()
	// 	text := s.Text()
	// 	fmt.Println(text)
	// }
}
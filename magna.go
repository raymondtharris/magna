package main

import (
	"bufio"
	"fmt"
	"magna/magnagraph"
	"magna/magnauser"
	"os"
	"strings"
)

var nodes []magnagraph.Node

func displayInput() {
	fmt.Println("What would you like to explore?")
	//Take in input from the user.
}

func introduction() {
	//Determine if this is a new user
	var newUser = magnauser.User{"", 113232444443}
	greeting(newUser)
}

func greeting(mUser magnauser.User) {
	if mUser.Name == "" {
		fmt.Println("Hello, I am Magna. Who might you be?")
		newUserName := readInput()
		newUser := magnauser.User{newUserName, 11122334234}
		fmt.Println("Nice to meet you " + newUser.Name + ".")
	} else {
		fmt.Println("Hello, " + mUser.Name + ". Welcome back.")
	}
}

func whatToSearch() {
	fmt.Println("What can I do for you?")
	queryInput := readInput()
	splitQuery := strings.Split(queryInput, " ")
	for index, aWord := range splitQuery {
		//create node from each word
		var newNode = magnagraph.Node{index, aWord, 1} // change Index to a randomly made hash of some sort
		nodes = append(nodes, newNode)
	}
	fmt.Println(nodes)
}

func readInput() string {
	consoleReader := bufio.NewReader(os.Stdin)
	rawInput, _ := consoleReader.ReadString('\n')
	inputValue := rawInput[0 : len(rawInput)-1]
	return inputValue
}

func main() {
	introduction()
	//for {
	whatToSearch()
	//}
}

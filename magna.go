package main

import (
	"bufio"
	"fmt"
	"magna/magnagraph"
	"magna/magnauser"
	"magna/magnaio"
	"os"
	"strings"
)

var nodes []magnagraph.Node
var magnasMind magnagraph.Graph

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
		newUserName := magnaio.ReadInput()
		newUser := magnauser.User{newUserName, 11122334234}
		fmt.Println("Nice to meet you " + newUser.Name + ".")
	} else {
		fmt.Println("Hello, " + mUser.Name + ". Welcome back.")
	}
}

func whatToSearch() {
	fmt.Println("What can I do for you?")
	queryInput := magnaio.ReadInput()
	splitQuery := strings.Split(queryInput, " ")
	for index, aWord := range splitQuery {
		//create node from each word
		var newNode = magnagraph.Node{index, aWord, 1, nil} // change Index to a randomly made hash of some sort
		nodes = append(nodes, newNode)
	}
	fmt.Println(nodes)
}



func loadMagnasMind() bool {
	magnasMind = magnagraph.Graph{0, 0, nil, true}
	return true
}

func main() {
	loadMagnasMind()
	introduction()
	//for {
	whatToSearch()
	//}
}

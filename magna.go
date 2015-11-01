package main

import (
	"fmt"
	"magna/magnagraph"
	"magna/magnauser"
	"magna/magnaio"
	"strings"
	"github.com/gorilla/websocket"
	"net/http"
)

var nodes []magnagraph.Node
var magnasMind magnagraph.Graph

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

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

func testHandler(w http.ResponseWriter, r *http.Request){
	connection, err := upgrader.Upgrade(w,r , nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		messageType, p, err1 := connection.ReadMessage()
		if err1 != nil {
			return
		}
		if err1 = connection.WriteMessage(messageType, p); err1 != nil {
			return err1
		}
	}
}

func main() {
	loadMagnasMind()
	http.HandleFunc("/", testHandler)
	http.ListenAndServe("localhost:8080", nil)
}

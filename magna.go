package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/magna/magnagraph"
	"github.com/magna/magnaio"
	"github.com/magna/magnauser"

	"github.com/gorilla/websocket"
)

type magnaQueryObject struct {
	User        magnauser.User
	QueryString string
}

var nodes []magnagraph.Node
var magnasMind magnagraph.Graph

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
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

func testHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		messageType, p, err1 := connection.ReadMessage()
		if err1 != nil {
			return
		}
		fmt.Println(string(p))
		err1 = connection.WriteMessage(messageType, p)
		if err1 != nil {
			return
		}
	}
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var query magnaQueryObject
	err = json.Unmarshal([]byte(string(body[:])), &query)
	if err != nil {
		panic(err)
	}
}

func main() {
	loadMagnasMind()
	http.HandleFunc("/", testHandler)
	http.HandleFunc("/query", queryHandler)
	http.ListenAndServe("localhost:8080", nil)
}

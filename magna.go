package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/magna/magnagraph"
	"github.com/magna/magnalanguage"
	"github.com/magna/magnauser"
)

type magnaQueryObject struct {
	User        magnauser.User
	QueryString string
}

func (mqo magnaQueryObject) String() string {
	return fmt.Sprintf("User: %v, Query: %v", mqo.User, mqo.QueryString)
}

//current query being processed
var queryTokens []magnagraph.Node

//graph used to solve queries
var magnasMind magnagraph.Graph

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func loadMagnasMind() bool {
	magnasMind = magnagraph.Graph{0, 0, nil, true}
	return true
}

func tokenizeQuery(queryObject magnaQueryObject) []magnagraph.Node {
	var tokenArray []magnagraph.Node
	splitQueryString := strings.Split(queryObject.QueryString, " ")
	for index, aWord := range splitQueryString {
		newNode := magnagraph.Node{index, aWord, 1, nil}
		tokenArray = append(tokenArray, newNode)
	}
	return tokenArray
}

func processQueryTokens(tokens []magnagraph.Node) {

}

// Handlers for Magana
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
	fmt.Println("Query Received.")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var query magnaQueryObject
	err = json.Unmarshal([]byte(string(body[:])), &query)
	if err != nil {
		panic(err)
	}
	fmt.Println(query)
	queryTokens = tokenizeQuery(query)
	for _, aNode := range queryTokens {
		magnalanguage.ProcessNode(aNode)
	}
}

func main() {
	loadMagnasMind()
	http.HandleFunc("/", testHandler)
	http.HandleFunc("/query", queryHandler)
	http.ListenAndServe("localhost:8080", nil)
}

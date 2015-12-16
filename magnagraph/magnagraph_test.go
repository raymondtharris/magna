package main

import(
  "fmt"
  "testing"
)


func TestListNode(t *testing.T){

}

func TestList(t *testing.T){

}


func TestNode(t *testing.T){
  words := string[]{"the", "world", "farther", "places", "trusting", "higher", "max", "light"}
  nodes := Node[]{}
  for _, word := range words { //Change to get a random word to add.
    anode := Node{1, word, 0, nil, 0, ""}
    nodes = append(nodes, anode)
  }
  fmt.Println(nodes)
}


func TestGraph(t *testing.T){
  words := string[]{"the", "world", "farther", "places", "trusting", "higher", "max", "light"}
  nodes := Node[]{}
  for _, word := range words { //Change to get a random word to add.
    anode := Node{1, word, 0, nil, 0, ""}
    nodes = append(nodes, anode)
  }
  aGraph := Graph{0,0, nil, false}
  for _, aNode := range nodes {
    aGraph.AddVertex(*aNode)
  }
  fmt.Println(aGraph)
}

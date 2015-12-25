package magnagraph_test

import(
  "fmt"
  "testing"
  . "github.com/magna/magnagraph"
)


func TestListNode(t *testing.T){
  words := []string{"the", "world", "farther", "places", "trusting", "higher", "max", "light"}
  nodes := []Node{}
  for _, word := range words { //Change to get a random word to add.
    anode := Node{1, word, 0, nil, 0, ""}
    nodes = append(nodes, anode)
  }
  qNodes := []node{}
  for _, aNode := range nodes {
    qNode := node{nil, *aNode}
    qNodes = append(qNodes, qNode)
  }
  fmt.Println(lNodes)
}

func TestQueue(t *testing.T){
  words := []string{"the", "world", "farther", "places", "trusting", "higher", "max", "light"}
  nodes := []Node{}
  for _, word := range words { //Change to get a random word to add.
    anode := Node{1, word, 0, nil, 0, ""}
    nodes = append(nodes, anode)
  }
  qNodes := []node{}
  for _, aNode := range nodes {
    qNode := node{nil, *aNode}
    qNodes = append(qNodes, qNode)
  }
  aQueue := queue{nil, nil, 0}
  for _, qnode := range qNodes{
    aQueue.Enqueue(*qnode)
  }
  fmt.Println(aQueue)
  fmt.Print(aQueue.Length)
  for !aQueue.isEmpty {
    aQueue.Dequeue()
    fmt.Println(aQueue)
  }
  fmt.Println(aQueue)
  fmt.Print(aQueue.Length)
}


func TestNode(t *testing.T){
  words := []string{"the", "world", "farther", "places", "trusting", "higher", "max", "light"}
  nodes := []Node{}
  for _, word := range words { //Change to get a random word to add.
    anode := Node{1, word, 0, nil, 0, ""}
    nodes = append(nodes, anode)
  }
  fmt.Println(nodes)
}


func TestGraph(t *testing.T){
  words := []string{"the", "world", "farther", "places", "trusting", "higher", "max", "light"}
  nodes := []Node{}
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

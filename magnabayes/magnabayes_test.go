package magnabayes_test

import(
  "log"
  . "github.com/magna/magnabayes"
  "testing"
)
testWords := []string{"the", "world", "farther", "places", "trusting", "higher", "max", "light"}
testNodes := []magnagraph.Node{}
findStrings := []string{"farther", "world", "trusting", "light", "winner", "max", "alien"}

func buildTestNodes(){
  testNodes = []magnagraph.Node{}
  for index, aWord :=  range testWords {
    newNode := magnagraph.Node{index, aWord, 1, nil, -1, ""}
    testNodes = append(testNodes, newNode)
  }
}

func TestAddEdgeBetween(t *testing.T){
  buildTestNodes()
  for index, aWord :=  range testWords {
    newNode := magnagraph.Node{index, aWord, 1, nil, -1, ""}
    testNodes = append(testNodes, newNode)
  }
  mDAG.AddEdgeBetween(mDAG.ADJList[0], mDAG.ADJList[2], 1)
  mDAG.AddEdgeBetween(mDAG.ADJList[1], mDAG.ADJList[7], 1)
  mDAG.AddEdgeBetween(mDAG.ADJList[4], mDAG.ADJList[6], 1)
  mDAG.AddEdgeBetween(mDAG.ADJList[5], mDAG.ADJList[3], 1)
}

func TestAddVertex(t *testing.T){
  buildTestNodes()
  mDAG := DAG{0,0,[]magnagraph.Node{}}
  for _, aNode := range testNodes {
    mDAG.AddVertex(aNode)
  }
}

func TestAddVertecies(t *testing.T){
  buildTestNodes()
  mDAG := DAG{0,0,[]magnagraph.Node{}}
  mDAG.ADDVertecies(testNodes)
}

func TestFindByString(t *testing.T){
  buildTestNodes()
  mDAG := DAG{0,0,[]magnagraph.Node{}}
  mDAG.ADDVertecies(testNodes)
  foundNodes := []magnagraph.Node{}
  for _, aString := range findStrings {
    foundVal := mDAG.findVertexByWord(aString)
    if foundVal != nil {
      foundNodes = append(foundNodes, foundVal)
    }
  }
}

func TestDAG(t *testing.T){
  buildTestNodes()
  for index, aWord :=  range testWords {
    newNode := magnagraph.Node{index, aWord, 1, nil, -1, ""}
    testNodes = append(testNodes, newNode)
  }
  mDAG.AddEdgeBetween(mDAG.ADJList[0], mDAG.ADJList[2], 1)
  mDAG.AddEdgeBetween(mDAG.ADJList[1], mDAG.ADJList[7], 1)
  mDAG.AddEdgeBetween(mDAG.ADJList[4], mDAG.ADJList[6], 1)
  mDAG.AddEdgeBetween(mDAG.ADJList[5], mDAG.ADJList[3], 1)
}

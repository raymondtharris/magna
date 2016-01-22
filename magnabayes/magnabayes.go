package magnabayes

import(
  "fmt"
  "github.com/magna/magnagraph"
)

type DAG struct{
	NumberOfVerticies int
	NumberOfEdges int
	ADJList []magnagraph.Node
}

func (mDAG DAG) String() string {
	return fmt.Sprintf("Verticies %v Edges %v Graph %v", mDAG.NumberOfVerticies, mDAG.NumberOfEdges, mDAG.ADJList)
}

//AddVertex function inserts a node into the ADJList if not present.
func (mDAG DAG) AddVertex(newVertex *magnagraph.Node) {
	isInGraph := false
	for _, aNode := range mDAG.ADJList {
		if &aNode == newVertex {
			isInGraph = true
		}
	}
	if isInGraph == false {
		mDAG.ADJList = append(mDAG.ADJList, *newVertex)
		mDAG.NumberOfVerticies++
	}
}

func (mDAG DAG) ADDVertecies(newVertecies []*magnagraph.Node) {
  for _, aVertex := range newVertecies {
    mDAG.AddVertex(aVertex)
  }
}

//AddEdgeBetween function inserts an edge between two vertecies in the graph. This is done
//by adding a Neighbor to the nodes.
func (mDAG DAG) AddEdgeBetween(vertex1 *Node, andVertex2 *Node, withCost int) {
	mDAG.AddVertex(vertex1)
	mDAG.AddVertex(andVertex2)
	vertex1.AddNeighbor(andVertex2, withCost)
}

//findVertex function searches through  ADJList for a node. Will be switched to a binary search.
func (mDAG DAG) findVertex(vertexToFind *Node) *Node, int {
	for index, aNode := range mDAG.ADJList {
		if &aNode == vertexToFind {
			return &aNode, index
		}
	}
	return nil, -1
}

func (mDAG DAG) findVertexByWord(aWord string) *Node, int {
	for index , aNode := range mDAG.ADJList {
		if &aNode.Value == aWord {
			return &aNode, index
		}
	}
	return nil, -1
}

func (mDAG DAG) RemoveVertex(vertexToRemove *magnagraph.Node) bool {
  res, index := mDAG.findVertex(vertexToRemove)
  if res != nil {
    tempArr1 := mDAG.ADJList[0,index-1]
    tempArr2 := mDAG.ADJList[index+1, len(mDAG.ADJList)]
    mDAG.ADJ = tempArr1 + tempArr2
    return true
  }
  return false

}

func (mDAG DAG) RemoveVertexByWord(wordToRemove string)  bool{
  res, index := mDAG.findVertexByWord(wordToRemove)
  if res != nil {
    tempArr1 := mDAG.ADJList[0,index-1]
    tempArr2 := mDAG.ADJList[index+1, len(mDAG.ADJList)]
    mDAG.ADJ = tempArr1 + tempArr2
    return true
  }
  return false
}

func CheckForCycles(aDAG DAG) bool{
  //Use BFS to search for cycles in the graph.
  return false
}

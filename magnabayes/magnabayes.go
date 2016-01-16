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
func (mDAG DAG) AddEdgeBetween(vertex1 *Node, andVertex2 *Node, withCosts []int) {
	mDAG.AddVertex(vertex1)
	mDAG.AddVertex(andVertex2)
	if len(withCosts) > 1 {
		vertex1.AddNeighbor(andVertex2, withCosts[1])
	} else {
		vertex1.AddNeighbor(andVertex2, withCosts[0])
	}
}

//findVertex function searches through  ADJList for a node. Will be switched to a binary search.
func (mDAG DAG) findVertex(vertexToFind *Node) *Node {
	for _, aNode := range mDAG.ADJList {
		if &aNode == vertexToFind {
			return &aNode
		}
	}
	return nil
}

func (mDAG DAG) findVertexByWord(aWord string) *Node {
	for _, aNode := range mDAG.ADJList {
		if &aNode.Value == aWord {
			return &aNode
		}
	}
	return nil
}

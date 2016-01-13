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
func (mDAG DAG) AddVertex(newVertex *Node) {
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

//AddEdgeBetween function inserts an edge between two vertecies in the graph. This is done
//by adding a Neighbor to the nodes.
func (mDAG Graph) AddEdgeBetween(vertex1 *Node, andVertex2 *Node, withCosts []int) {
	mGraph.AddVertex(vertex1)
	mGraph.AddVertex(andVertex2)
	if len(withCosts) > 1 {
		vertex1.AddNeighbor(andVertex2, withCosts[1])
		if mGraph.IsDirected {
			andVertex2.AddNeighbor(vertex1, withCosts[0])
		}
	} else {
		vertex1.AddNeighbor(andVertex2, withCosts[0])
		if mGraph.IsDirected {
			andVertex2.AddNeighbor(vertex1, withCosts[0])
		}
	}
}

//findVertex function searches through  ADJList for a node. Will be switched to a binary search.
func (mGraph Graph) findVertex(vertexToFind *Node) *Node {
	for _, aNode := range mGraph.ADJList {
		if &aNode == vertexToFind {
			return &aNode
		}
	}
	return nil
}

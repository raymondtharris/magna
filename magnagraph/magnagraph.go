package magnagraph

//"fmt"

//Graph Code

//Graph is comething
type Graph struct {
	NumberOfVerticies int
	NumberOfEdges     int
	ADJList           []Node
	isBidirectional   bool
}

//Node is something
type Node struct {
	Index     int        //Inex of the node
	Value     string     //String value for the particular node
	Weight    int        // Weight of the node in relation to the others
	Neighbors []Neighbor //Array of nodes connected to this particular node
}

//Neighbor struct to store a node and the cost of the node
type Neighbor struct {
	Vertex Node
	Cost   int
}

//AddVertex function inserts a node into the ADJList if not present.
func (mGraph Graph) AddVertex(newVertex *Node) {
	isInGraph := false
	for _, aNode := range mGraph.ADJList {
		if &aNode == newVertex {
			isInGraph = true
		}
	}
	if isInGraph == false {
		mGraph.ADJList = append(mGraph.ADJList, *newVertex)
		mGraph.NumberOfVerticies++
	}
}

//AddNeighbor inserts a neighbor node in to an array to keep track of connections
// to a particular node.
func (mNode Node) AddNeighbor(newNeighbor *Node, withCost int) {
	isANeighbor := false
	savedIndex := 0
	for index, aNode := range mNode.Neighbors {
		if &aNode.Vertex == newNeighbor {
			isANeighbor = true
			savedIndex = index
		}
	}
	if isANeighbor == true {
		mNode.Neighbors[savedIndex].Cost = withCost
	} else {
		mNode.Neighbors = append(mNode.Neighbors, Neighbor{*newNeighbor, withCost})
	}
}

//AddEdgeBetween function inserts an edge between two vertecies in the graph. This is done
//by adding a Neighbor to the nodes.
func (mGraph Graph) AddEdgeBetween(vertex1 *Node, andVertex2 *Node, withCosts []int) {
	mGraph.AddVertex(vertex1)
	mGraph.AddVertex(andVertex2)
	if len(withCosts) > 1 {
		vertex1.AddNeighbor(andVertex2, withCosts[1])
		if mGraph.isBidirectional {
			andVertex2.AddNeighbor(vertex1, withCosts[0])
		}
	} else {
		vertex1.AddNeighbor(andVertex2, withCosts[0])
		if mGraph.isBidirectional {
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

//Queue Code

//Node for a queue
type node struct {
	Next    *node
	Payload *Node
}

//Queue data structure
type queue struct {
	Head   *node
	Tail   *node
	Length int
}

//Push inserts a Node at the end of the List
func (mQueue queue) Enqueue(newNode *node) {
	if mQueue.Head == nil {
		mQueue.Head = newNode
		mQueue.Tail = mQueue.Head
	} else {
		mQueue.Tail.Next = newNode
		mQueue.Tail = newNode
	}
	mQueue.Length++
}

//Pop removes a Node from the end of the queue
func (mQueue queue) Dequque() *node {
	if mQueue.Head != nil {
		nodeToReturn := mQueue.Head
		mQueue.Head = mQueue.Head.Next
		nodeToReturn.Next = nil
		mQueue.Length--
		return nodeToReturn
	}
	return nil
}

func (mQueue queue) isEmpty() bool {
	if mQueue.Head == nil {
		return true
	}
	return false

}

//Search will look for a Node and return in from a queue
func (mQueue queue) Search(nodeToFind *Node) *node {
	currentNode := mQueue.Head
	for currentNode != nil {
		if currentNode.Payload == nodeToFind {
			return currentNode
		}
		currentNode = currentNode.Next

	}
	return nil
}

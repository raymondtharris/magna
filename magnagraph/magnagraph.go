package magnagraph
import "fmt"

const (
	PRONOUN   = 1
	NOUN      = 2
	ADJECTIVE = 3
	VERB      = 4
	ADVERB    = 5
)

//Graph Code

type Path struct {
	Cost     int
	Sequence queue
}

func (mPath Path) String() string {
	return fmt.Sprintf("Cost: %v Path: %v", mPath.Cost, mPath.Sequence)
}

//Graph is comething
type Graph struct {
	NumberOfVerticies int
	NumberOfEdges     int
	ADJList           []Node
	//IsDirected        bool
}

func (mGraph Graph) String() string {
	return fmt.Sprintf("Verticies: %v Edges: %v Directed: %v\n  List: %v", mGraph.NumberOfVerticies, mGraph.NumberOfEdges, mGraph.IsDirected, mGraph.ADJList)
}



//Node is something
type Node struct {
	Index     int        //Inex of the node
	Value     string     //String value for the particular node
	Weight    int        // Weight of the node in relation to the others
	Neighbors []Neighbor //Array of nodes connected to this particular node
	Measure   int        // Measure of the word value
	Stem      string     //Stem found from the word value.
}

func (mNode Node) String() string {
	return fmt.Sprintf("Value: %v Measure: %v Stem: %v Weight: %v Neighbors: %v\n", mNode.Value, mNode.Measure, mNode.Stem,mNode.Weight, mNode.Neighbors)
}

//Neighbor struct to store a node and the cost of the node
type Neighbor struct {
	Vertex Node
	Cost   int
}

func (mNeighbor Neighbor) String() string {
	return fmt.Sprintf("Node: %v Cost: %v", mNeighbor.Vertex, mNeighbor.Cost)
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

//findShortestPathBetween function searches for the shortest path of two nodes using a breadth first search
//algorithm.
func (mGraph Graph) findShortestPathBetween(startVertex *Node, andEndVertex *Node) {
	originNode := mGraph.findVertex(startVertex)
	nodesToCheck := queue{nil, nil, 0}
	path := queue{nil, nil, 0}
	if originNode == andEndVertex {
		newInsert := node{nil, originNode}
		path.Enqueue(&newInsert) // Change to returnPath
	}
	newInsert := node{nil, originNode}
	path.Enqueue(&newInsert)
	for nodesToCheck.Length > 0 {
		currentNode := nodesToCheck.Dequque()
		for _, neighborNode := range currentNode.Payload.Neighbors {
			newNode := node{nil, &neighborNode.Vertex}
			nodesToCheck.Enqueue(&newNode)
		}
	}
}

//Queue Code

//Node for a queue
type node struct {
	Next    *node
	Payload *Node
}

func (qNode node) String() string {
	return fmt.Sprintf("%v ", qNode.Payload)
}

//Queue data structure
type queue struct {
	Head   *node
	Tail   *node
	Length int
}

func (mQueue queue) String() string {
	current := mQueue.Head
	prString := "" + current.String() + " -> "
	for current.Next != nil {
		current = current.Next
		if current.Next != nil {
			prString = prString + " " + current.String() + " -> "
		} else {
			prString = prString + " " + current.String() + " "
		}
	}

	return fmt.Sprintf("%v", prString)
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

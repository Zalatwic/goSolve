//kps59
//astar for arraylike object

package Astar

import "fmt"

type fn func([]int) int
type ssn func([]int)
type gsn func(int) [][]int

type Astar struct {
	MaxNodes int
}

//create nodes with counters for linked list
//counter represents actual distance traveled
type Node struct {
	state         []int
	cost          int
	distance      int
	previousMoves []string
}

//create first node through evaluation of the current state hMod
func (x *Astar) solve(board ePuzzle, currentState []int, hMod fn, setState ssn, stateOptions gsn, moveNames []string, goalState []int) {
	originNode := headNode{currentState, hMod(currentState)}
	var uncheckedNodeSlice []Node
	globalNodeAmount := 0
	run := true

	for run {
		candidate := x.findCandidate(uncheckedNodeSlice)

		if testEquality(candidate.state, goalState) {
			run = false
			fmt.Println(candidate.previousMoves)
		}

		if globalNodeAmount > x.MaxNodes {
			run = false
			fmt.Println("astar search out of moves")
		}

		setState(candidate)
		for B, focusNode := range stateOptions(0) {
			globalNodeAmount++
			uncheckedNodeSlice = append(uncheckedNodeSlice, Node{focusNode, (candidate.distance + 1 + hMod(focusNode)), (candidate.distance + 1), append(candidate.previousMoves, moveNames[B])})
		}

		//remove candidate from uncheckedNodeSlice

	}
}

//search through tree to find lowest h value
func (x Astar) findCandidate(nodeList []Node) Node {
	currentMin := -1
	var selectedNode Node

	for A, currentNode := range nodeList {
		if currentMin == -1 || currentMin > currentNode.cost {
			currentMin = currentNode.cost
			selectedNode = currentNode
		}
	}

	return selectedNode
}

func testEquality(x, y []int) bool {
	//check if theyre empty
	if (x == nil) != (y == nil) {
		return false
	}

	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

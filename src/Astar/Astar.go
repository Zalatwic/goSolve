//kps59
//implimentation of beam search given any puzzle object

package Astar

import "fmt"
import "sort"

type fn func([]int) int
type ssn func([]int)
type gsn func(int) [][]int

type Astar struct {
	MaxNodes   int
	StateNames []string
}

type Node struct {
	currentState  []int
	cost          int
	previousMoves []string
}

func (x *Astar) SolvePuzzle(inState []int, hMod fn, setState ssn, stateOptions gsn) {
	//make slice and limit nodes to k
	var nodeSlice []Node
	var evaluatedMoves [][]int

	//count the amount of nodes generated
	nodeCounter := 0

	//add the head node to evaluate
	nodeSlice = append(nodeSlice, Node{inState, hMod(inState), []string{"origin"}})

	//sort to find the lowest cost ofx.k options then execute
	for true {
		sort.Slice(nodeSlice, func(i, j int) bool { return nodeSlice[i].cost < nodeSlice[j].cost })

		//find an unevaluated node and evaluate it
		var chosenNode Node
		chosenNode = nodeSlice[0]

		var newMoves []string
		newMoves = copySlice(chosenNode.previousMoves)
		setState(copySliceInt(chosenNode.currentState))

		evaluatedMoves = append(evaluatedMoves, copySliceInt(chosenNode.currentState))

		hCost := 0
		var moveOptions []Node
		var potentialStates [][]int
		potentialStates = stateOptions(0)

		for i, possibleMove := range potentialStates {
			nodeCounter++

			//set cost equal to heuristic
			hCost = hMod(possibleMove) + len(newMoves) + 1
			fmt.Println("possible move: ", possibleMove, "cost", hCost)

			fmt.Println(chosenNode)

			if nodeCounter > x.MaxNodes {
				fmt.Println("beam failed to find result in", x.MaxNodes, "nodes")
				return
			}

			//if the score is zero, the goal is achieved
			if hMod(possibleMove) == 0 {
				//win condition
				fmt.Println(append(newMoves, x.StateNames[i]))
				fmt.Println("found in", nodeCounter, "moves")
				return
			}

			//ensure this node hasnt been evaluated prior
			if !contains(evaluatedMoves, possibleMove) {
				moveOptions = append(moveOptions, Node{possibleMove, hCost, append(newMoves, x.StateNames[i])})
			}
		}

		for _, currentNode := range nodeSlice {
			if !contains(evaluatedMoves, currentNode.currentState) {
				moveOptions = append(moveOptions, currentNode)
			}

		}

		//sort the options for the search next iteration and return top seven candidates
		sort.Slice(moveOptions, func(i, j int) bool { return moveOptions[i].cost < moveOptions[j].cost })
		nodeSlice = moveOptions
	}
}

func copySlice(s []string) []string {
	t := make([]string, len(s))
	copy(t, s)
	return t
}

func copySliceInt(i []int) []int {
	t := make([]int, len(i))
	copy(t, i)
	return t
}

func contains(s [][]int, e []int) bool {
	for _, a := range s {
		if testEquality(a, e) {
			return true
		}
	}
	return false
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

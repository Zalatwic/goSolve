//kps59
//astar for arraylike object

package Astar

import "fmt"

type fn func(int)

type astar struct {
	currentState
	heuristicType	fn
}

//create nodes with counters for linked list
//counter represents actual distance traveled
type Node struct {
	state []int
	cost int
	nextNode Node
	previousNode Node
}

//create first node through evaluation of the current state hMod
func (x astar) solve(board ePuzzle, currentState [9]int, hMod fn) {
	hMod(currentState)
}

//get states
//run h on every state possibility
//link nodes to tree
//search through tree to find lowest h value
//repeat until finished



//kps59
//astar for arraylike object

package 8puzzle

import fmt

type astar struct {
	currentState
	astarTyoe
	controlUnit
}

//create nodes with counters for linked list
//counter represents actual distance traveled
type Node struct {
	state [9]int
	cost int
	nextNode Node
	previousNode Node
}

func (x astar, y  ) solve {
	
}
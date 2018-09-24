//kps59
//8puzzle object

package ePuzzle

import "fmt"

type ePuzzle struct {
	boardConfig	[9]int
}

func (x ePuzzle) setState(y [9]int) {
	x.boardConfig = y
}

func (x ePuzzle) findNum(y int) int {
	z := 0

	for y != x.boardConfig[z] {
		z++
	}

	return z
}

func (x ePuzzle) left(y int) [9]int {
	switch x.findNum(y) {
	case 0: fmt.Println("Space already on left.")
	case 1:
		x.boardConfig[1] = x.boardConfig[0]
		x.boardConfig[0] = 0
		fmt.Println(x.boardConfig[1], " now occupies space 1.")
	case 2:
		x.boardConfig[2] = x.boardConfig[1]
		x.boardConfig[1] = 0
		fmt.Println(x.boardConfig[2], " now occupies space 2.")
	case 3: fmt.Println("Space already on left.")
	case 4:
		x.boardConfig[4] = x.boardConfig[3]
		x.boardConfig[3] = 0
		fmt.Println(x.boardConfig[4], " now occupies space 4.")
	case 5:
		x.boardConfig[5] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[5], " now occupies space 5.")
	case 6: fmt.Println("Space already on left.")
	case 7:
		x.boardConfig[7] = x.boardConfig[6]
		x.boardConfig[6] = 0
		fmt.Println(x.boardConfig[7], " now occupies space 7.")
	case 8:
		x.boardConfig[8] = x.boardConfig[7]
		x.boardConfig[7] = 0
		fmt.Println(x.boardConfig[8], " now occupies space 8.")
	}

	return x.boardConfig
}

func (x ePuzzle) right(y int) [9]int {
	switch x.findNum(y) {
	case 0:
		x.boardConfig[0] = x.boardConfig[1]
		x.boardConfig[1] = 0
		fmt.Println(x.boardConfig[0], " now occupies space 0.")
	case 1:
		x.boardConfig[1] = x.boardConfig[2]
		x.boardConfig[2] = 0
		fmt.Println(x.boardConfig[1], " now occupies space 1.")
	case 2: fmt.Println("Space already on right.")
	case 3:
		x.boardConfig[3] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[3], " now occupies space 3.")
	case 4:
		x.boardConfig[4] = x.boardConfig[5]
		x.boardConfig[5] = 0
		fmt.Println(x.boardConfig[4], " now occupies space 4.")
	case 5: fmt.Println("Space already on right.")
	case 6:
		x.boardConfig[6] = x.boardConfig[7]
		x.boardConfig[7] = 0
		fmt.Println(x.boardConfig[6], " now occupies space 6.")
	case 7:
		x.boardConfig[7] = x.boardConfig[8]
		x.boardConfig[8] = 0
		fmt.Println(x.boardConfig[7], " now occupies space 7.")
	case 8: fmt.Println("Space already on right.")
	}

	return x.boardConfig
}

func (x ePuzzle) up(y int) [9]int {
	switch x.findNum(y) {
	case 0: fmt.Println("Space already on top.")
	case 1: fmt.Println("Space already on top.")
	case 2: fmt.Println("Space already on top.")
	case 3:
		x.boardConfig[3] = x.boardConfig[0]
		x.boardConfig[0] = 0
		fmt.Println(x.boardConfig[3], " now occupies space 3.")
	case 4:
		x.boardConfig[4] = x.boardConfig[1]
		x.boardConfig[1] = 0
		fmt.Println(x.boardConfig[4], " now occupies space 4.")
	case 5:
		x.boardConfig[5] = x.boardConfig[2]
		x.boardConfig[2] = 0
		fmt.Println(x.boardConfig[5], " now occupies space 5.")
	case 6:
		x.boardConfig[6] = x.boardConfig[3]
		x.boardConfig[3] = 0
		fmt.Println(x.boardConfig[6], " now occupies space 6.")
	case 7:
		x.boardConfig[7] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[7], " now occupies space 7.")
	case 8:
		x.boardConfig[8] = x.boardConfig[5]
		x.boardConfig[5] = 0
		fmt.Println(x.boardConfig[8], " now occupies space 8.")
	}

	return x.boardConfig
}

func (x ePuzzle) down(y int) [9]int {
	switch x.findNum(y) {
	case 0:
		x.boardConfig[0] = x.boardConfig[3]
		x.boardConfig[3] = 0
		fmt.Println(x.boardConfig[0], " now occupies space 0.")
	case 1:
		x.boardConfig[1] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[1], " now occupies space 1.")
	case 2:
		x.boardConfig[2] = x.boardConfig[5]
		x.boardConfig[5] = 0
		fmt.Println(x.boardConfig[2], " now occupies space 2.")
	case 3:
		x.boardConfig[3] = x.boardConfig[6]
		x.boardConfig[6] = 0
		fmt.Println(x.boardConfig[3], " now occupies space 3.")
	case 4:
		x.boardConfig[4] = x.boardConfig[7]
		x.boardConfig[7] = 0
		fmt.Println(x.boardConfig[4], " now occupies space 4.")
	case 5:
		x.boardConfig[5] = x.boardConfig[8]
		x.boardConfig[8] = 0
		fmt.Println(x.boardConfig[5], " now occupies space 5.")
	case 6: fmt.Println("Space already on bottom.")
	case 7: fmt.Println("Space already on bottom.")
	case 8: fmt.Println("Space already on bottom.")

	}
	return x.boardConfig
}

func (x ePuzzle) getStates(y int) [4][9]int {
	//store the current board config in memory as to not disrupt state in search
	configMemory = x.boardConfig
	multiArray := [4][9]int

	//create an array of potential moves
	multiArray[0] = x.left(y)
	x.setState(configMemory)

	multiArray[1] = x.right(y)
	x.setState(configMemory)

	multiArray[2] = x.up(y)
	x.setState(configMemory)

	multiArray[3] = x.down(y)
	x.setState(configMemory)

	reutrn multiArray
}

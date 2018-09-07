//kps59
//8puzzle object

package 8puzzle

import (
	"fmt"
)

boardConfig := [8]int{0, 1, 2, 3, 4, 5, 6, 7, 8}

type 8puzzle struct {
	topLeft
	topMiddle
	topRight
	middleLeft
	middleMiddle
	middleRight
	bottomLeft
	bottomMiddle
	bottomRight
}

func (x 8puzzle) makeArray(y int) {
	x.boardConfig = {x.topLeft, x.topMiddle, x.topRight, x.middleLeft, x.middleMiddle, x.middleRight, x.bottomLeft, x.bottomMiddle, x.bottomRight}
}

func (x 8puzzle) findNum(y int) int {
	z := 0

	for y != x.boardConfig[z] {
		z++
	}

	return z
}

func (x 8puzzle) left() []int {
	switch x.findNum(0) {
	case 0: fmt.Println("Space already on left.")
	case 1:
		x.boardConfig[1] = x.boardConfig[0]
		x.boardConfig[0] = 0
		fmt.Println(x.boardConfig[1] + " now occupies space 1.")
	case 2:
		x.boardConfig[2] = x.boardConfig[1]
		x.boardConfig[1] = 0
		fmt.Println(x.boardConfig[2] + " now occupies space 2.")
	case 3: fmt.Println("Space already on left.")
	case 4:
		x.boardConfig[4] = x.boardConfig[3]
		x.boardConfig[3] = 0
		fmt.Println(x.boardConfig[4] + " now occupies space 4.")
	case 5:
		x.boardConfig[5] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[5] + " now occupies space 5.")
	case 6: fmt.Println("Space already on left.")
	case 7:
		x.boardConfig[7] = x.boardConfig[6]
		x.boardConfig[6] = 0
		fmt.Println(x.boardConfig[7] + " now occupies space 7.")
	case 8:
		x.boardConfig[8] = x.boardConfig[7]
		x.boardConfig[7] = 0
		fmt.Println(x.boardConfig[8] + " now occupies space 8.")
	}

	return x.boardConfig
}

func (x 8puzzle) right() []int {
	switch x.findNum(0) {
	case 0:
		x.boardConfig[0] = x.boardConfig[1]
		x.boardConfig[1] = 0
		fmt.Println(x.boardConfig[0] + " now occupies space 0.")
	case 1:
		x.boardConfig[1] = x.boardConfig[2]
		x.boardConfig[2] = 0
		fmt.Println(x.boardConfig[1] + " now occupies space 1.")
	case 2: fmt.Println("Space already on right.")
	case 3:
		x.boardConfig[3] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[3] + " now occupies space 3.")
	case 4:
		x.boardConfig[4] = x.boardConfig[5]
		x.boardConfig[5] = 0
		fmt.Println(x.boardConfig[4] + " now occupies space 4.")
	case 5: fmt.Println("Space already on right.")
	case 6:
		x.boardConfig[6] = x.boardConfig[7]
		x.boardConfig[7] = 0
		fmt.Println(x.boardConfig[6] + " now occupies space 6.")
	case 7:
		x.boardConfig[7] = x.boardConfig[8]
		x.boardConfig[8] = 0
		fmt.Println(x.boardConfig[7] + " now occupies space 7.")
	case 8: fmt.Println("Space already on right.")
	}

	return x.boardConfig
}

func (x 8puzzle) up() []int {
	switch x.findNum(0) {
	case 0: fmt.Println("Space already on top.")
	case 1: fmt.Println("Space already on top.")
	case 2: fmt.Println("Space already on top.")
	case 3:
		x.boardConfig[3] = x.boardConfig[0]
		x.boardConfig[0] = 0
		fmt.Println(x.boardConfig[3] + " now occupies space 3.")
	case 4:
		x.boardConfig[4] = x.boardConfig[1]
		x.boardConfig[1] = 0
		fmt.Println(x.boardConfig[4] + " now occupies space 4.")
	case 5:
		x.boardConfig[5] = x.boardConfig[2]
		x.boardConfig[2] = 0
		fmt.Println(x.boardConfig[5] + " now occupies space 5.")
	case 6:
		x.boardConfig[6] = x.boardConfig[3]
		x.boardConfig[3] = 0
		fmt.Println(x.boardConfig[6] + " now occupies space 6.")
	case 7:
		x.boardConfig[7] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[7] + " now occupies space 7.")
	case 8:
		x.boardConfig[8] = x.boardConfig[5]
		x.boardConfig[5] = 0
		fmt.Println(x.boardConfig[8] + " now occupies space 8.")
	}

	return x.boardConfig
}

func (x 8puzzle) down() []int {
	switch x.findNum(0) {
	case 0:
		x.boardConfig[0] = x.boardConfig[3]
		x.boardConfig[3] = 0
		fmt.Println(x.boardConfig[0] + " now occupies space 0.")
	case 1:
		x.boardConfig[1] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[1] + " now occupies space 1.")
	case 2:
		x.boardConfig[2] = x.boardConfig[5]
		x.boardConfig[5] = 0
		fmt.Println(x.boardConfig[2] + " now occupies space 2.")
	case 3:
		x.boardConfig[3] = x.boardConfig[6]
		x.boardConfig[6] = 0
		fmt.Println(x.boardConfig[3] + " now occupies space 3.")
	case 4:
		x.boardConfig[4] = x.boardConfig[7]
		x.boardConfig[7] = 0
		fmt.Println(x.boardConfig[4] + " now occupies space 4.")
	case 5:
		x.boardConfig[5] = x.boardConfig[8]
		x.boardConfig[8] = 0
		fmt.Println(x.boardConfig[5] + " now occupies space 5.")
	}
	case 6: fmt.Println("Space already on bottom.")
	case 7: fmt.Println("Space already on bottom.")
	case 8: fmt.Println("Space already on bottom.")
	return x.boardConfig
}

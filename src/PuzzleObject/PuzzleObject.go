//kps59
//8puzzle object

package PuzzleObject

import "fmt"

type Puzzle struct {
	boardConfig	[]int
}

func (x *Puzzle) SetState(y []int) {
	x.boardConfig = y
	fmt.Println(x.boardConfig)
}

func (x *Puzzle) GetState() []int {
	fmt.Println(x.boardConfig)
	return x.boardConfig
}

func (x *Puzzle) FindNum(y int) int {
	z := 0

	for y != x.boardConfig[z] {
		z++
	}

	return z
}

func (x *Puzzle) Left(y int) []int {
	switch x.FindNum(y) {
	case 0: fmt.Println("Space already on Left.")
	case 1:
		x.boardConfig[1] = x.boardConfig[0]
		x.boardConfig[0] = 0
		fmt.Println(x.boardConfig[1], "now occUpies space 1.")
	case 2:
		x.boardConfig[2] = x.boardConfig[1]
		x.boardConfig[1] = 0
		fmt.Println(x.boardConfig[2], "now occUpies space 2.")
	case 3: fmt.Println("Space already on Left.")
	case 4:
		x.boardConfig[4] = x.boardConfig[3]
		x.boardConfig[3] = 0
		fmt.Println(x.boardConfig[4], "now occUpies space 4.")
	case 5:
		x.boardConfig[5] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[5], "now occUpies space 5.")
	case 6: fmt.Println("Space already on Left.")
	case 7:
		x.boardConfig[7] = x.boardConfig[6]
		x.boardConfig[6] = 0
		fmt.Println(x.boardConfig[7], "now occUpies space 7.")
	case 8:
		x.boardConfig[8] = x.boardConfig[7]
		x.boardConfig[7] = 0
		fmt.Println(x.boardConfig[8], "now occUpies space 8.")
	}

	return x.boardConfig
}

func (x *Puzzle) Right(y int) []int {
	switch x.FindNum(y) {
	case 0:
		x.boardConfig[0] = x.boardConfig[1]
		x.boardConfig[1] = 0
		fmt.Println(x.boardConfig[0], " now occUpies space 0.")
	case 1:
		x.boardConfig[1] = x.boardConfig[2]
		x.boardConfig[2] = 0
		fmt.Println(x.boardConfig[1], " now occUpies space 1.")
	case 2: fmt.Println("Space already on Right.")
	case 3:
		x.boardConfig[3] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[3], " now occUpies space 3.")
	case 4:
		x.boardConfig[4] = x.boardConfig[5]
		x.boardConfig[5] = 0
		fmt.Println(x.boardConfig[4], " now occUpies space 4.")
	case 5: fmt.Println("Space already on Right.")
	case 6:
		x.boardConfig[6] = x.boardConfig[7]
		x.boardConfig[7] = 0
		fmt.Println(x.boardConfig[6], " now occUpies space 6.")
	case 7:
		x.boardConfig[7] = x.boardConfig[8]
		x.boardConfig[8] = 0
		fmt.Println(x.boardConfig[7], " now occUpies space 7.")
	case 8: fmt.Println("Space already on Right.")
	}

	return x.boardConfig
}

func (x *Puzzle) Up(y int) []int {
	switch x.FindNum(y) {
	case 0: fmt.Println("Space already on top.")
	case 1: fmt.Println("Space already on top.")
	case 2: fmt.Println("Space already on top.")
	case 3:
		x.boardConfig[3] = x.boardConfig[0]
		x.boardConfig[0] = 0
		fmt.Println(x.boardConfig[3], " now occUpies space 3.")
	case 4:
		x.boardConfig[4] = x.boardConfig[1]
		x.boardConfig[1] = 0
		fmt.Println(x.boardConfig[4], " now occUpies space 4.")
	case 5:
		x.boardConfig[5] = x.boardConfig[2]
		x.boardConfig[2] = 0
		fmt.Println(x.boardConfig[5], " now occUpies space 5.")
	case 6:
		x.boardConfig[6] = x.boardConfig[3]
		x.boardConfig[3] = 0
		fmt.Println(x.boardConfig[6], " now occUpies space 6.")
	case 7:
		x.boardConfig[7] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[7], " now occUpies space 7.")
	case 8:
		x.boardConfig[8] = x.boardConfig[5]
		x.boardConfig[5] = 0
		fmt.Println(x.boardConfig[8], " now occUpies space 8.")
	}

	return x.boardConfig
}

func (x *Puzzle) Down(y int) []int {
	switch x.FindNum(y) {
	case 0:
		x.boardConfig[0] = x.boardConfig[3]
		x.boardConfig[3] = 0
		fmt.Println(x.boardConfig[0], " now occUpies space 0.")
	case 1:
		x.boardConfig[1] = x.boardConfig[4]
		x.boardConfig[4] = 0
		fmt.Println(x.boardConfig[1], " now occUpies space 1.")
	case 2:
		x.boardConfig[2] = x.boardConfig[5]
		x.boardConfig[5] = 0
		fmt.Println(x.boardConfig[2], " now occUpies space 2.")
	case 3:
		x.boardConfig[3] = x.boardConfig[6]
		x.boardConfig[6] = 0
		fmt.Println(x.boardConfig[3], " now occUpies space 3.")
	case 4:
		x.boardConfig[4] = x.boardConfig[7]
		x.boardConfig[7] = 0
		fmt.Println(x.boardConfig[4], " now occUpies space 4.")
	case 5:
		x.boardConfig[5] = x.boardConfig[8]
		x.boardConfig[8] = 0
		fmt.Println(x.boardConfig[5], " now occUpies space 5.")
	case 6: fmt.Println("Space already on bottom.")
	case 7: fmt.Println("Space already on bottom.")
	case 8: fmt.Println("Space already on bottom.")

	}
	return x.boardConfig
}

func (x Puzzle) GetStates(y int) (multiArray [][]int) { //, names []string) {
	//store the current board config in memory as to not disrUpt state in search
	configMemory := x.boardConfig
	//var names []string

	//create an array of potential moves
	multiArray = append(multiArray, x.Left(y))
	//names[0] = "left"
	x.SetState(configMemory)

	multiArray = append(multiArray, x.Right(y))
	//names[1] = "right"
	x.SetState(configMemory)

	multiArray = append(multiArray, x.Up(y))
	//names[2] = "up"
	x.SetState(configMemory)

	multiArray = append(multiArray, x.Down(y))
	//names[3] = "down"
	x.SetState(configMemory)

	return
}

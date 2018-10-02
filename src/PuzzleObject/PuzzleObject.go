//kps59
//8puzzle object

package PuzzleObject

import "fmt"

type Puzzle struct {
	hiddenConfig []int
}

func (x *Puzzle) SetState(y []int) {
	x.hiddenConfig = y
	fmt.Println(x.hiddenConfig, y)
}

func (x *Puzzle) GetState() []int {
	fmt.Println(x.hiddenConfig)
	return copySlice(x.hiddenConfig)
}

func (x *Puzzle) FindNum(y int) int {
	z := 0

	fmt.Println("kk", x.hiddenConfig)

	for y != x.hiddenConfig[z] {
		z++
	}

	return z
}

func (x *Puzzle) Left(y int) []int {
	boardConfig := x.GetState()
	switch x.FindNum(y) {
	case 1:
		boardConfig[1] = boardConfig[0]
		boardConfig[0] = 0
	case 2:
		boardConfig[2] = boardConfig[1]
		boardConfig[1] = 0
	case 4:
		boardConfig[4] = boardConfig[3]
		boardConfig[3] = 0
	case 5:
		boardConfig[5] = boardConfig[4]
		boardConfig[4] = 0
	case 7:
		boardConfig[7] = boardConfig[6]
		boardConfig[6] = 0
	case 8:
		boardConfig[8] = boardConfig[7]
		boardConfig[7] = 0
	}

	return boardConfig
}

func (x *Puzzle) Right(y int) []int {
	boardConfig := x.GetState()
	switch x.FindNum(y) {
	case 0:
		boardConfig[0] = boardConfig[1]
		boardConfig[1] = 0
	case 1:
		boardConfig[1] = boardConfig[2]
		boardConfig[2] = 0
	case 3:
		boardConfig[3] = boardConfig[4]
		boardConfig[4] = 0
	case 4:
		boardConfig[4] = boardConfig[5]
		boardConfig[5] = 0
	case 6:
		boardConfig[6] = boardConfig[7]
		boardConfig[7] = 0
	case 7:
		boardConfig[7] = boardConfig[8]
		boardConfig[8] = 0
	}

	return boardConfig
}

func (x *Puzzle) Up(y int) []int {
	boardConfig := x.GetState()
	switch x.FindNum(y) {
	case 3:
		boardConfig[3] = boardConfig[0]
		boardConfig[0] = 0
	case 4:
		boardConfig[4] = boardConfig[1]
		boardConfig[1] = 0
	case 5:
		boardConfig[5] = boardConfig[2]
		boardConfig[2] = 0
	case 6:
		boardConfig[6] = boardConfig[3]
		boardConfig[3] = 0
	case 7:
		boardConfig[7] = boardConfig[4]
		boardConfig[4] = 0
	case 8:
		boardConfig[8] = boardConfig[5]
		boardConfig[5] = 0
	}

	return boardConfig
}

func (x *Puzzle) Down(y int) []int {
	boardConfig := x.GetState()
	switch x.FindNum(y) {
	case 0:
		boardConfig[0] = boardConfig[3]
		boardConfig[3] = 0
	case 1:
		boardConfig[1] = boardConfig[4]
		boardConfig[4] = 0
	case 2:
		boardConfig[2] = boardConfig[5]
		boardConfig[5] = 0
	case 3:
		boardConfig[3] = boardConfig[6]
		boardConfig[6] = 0
	case 4:
		boardConfig[4] = boardConfig[7]
		boardConfig[7] = 0
	case 5:
		boardConfig[5] = boardConfig[8]
		boardConfig[8] = 0
	}

	return boardConfig
}

//get potential states for the 8puzzle
//note, configMemory original pointers were altered despite copying
//^ that was a bug caused by call-by-ref where call-by-val was needed
//this block works, and copying the default state is perhaps unneeded
//however, if you substitute the get/set methods for direct access to the hidden state
//and run the code below, the first execution of x.SetState(configMemory) will be
//call-by-value, however the other three will be call-by-reference
//now that's an understanding of language semantics i feel proud losing sleep for
func (x *Puzzle) GetStates(y int) [][]int {
	//store the current board config in memory as to not disrUpt state in search
	configMemory := x.GetState()
	var multiArray [][]int

	//create an array of potential moves
	multiArray = append(multiArray, copySlice(x.Left(y)))
	x.SetState(configMemory)

	multiArray = append(multiArray, copySlice(x.Right(y)))
	x.SetState(configMemory)

	multiArray = append(multiArray, copySlice(x.Up(y)))
	x.SetState(configMemory)

	multiArray = append(multiArray, copySlice(x.Down(y)))
	x.SetState(configMemory)

	return multiArray
}

func copySlice(s []int) []int {
	t := make([]int, len(s))
	copy(t, s)
	return t
}

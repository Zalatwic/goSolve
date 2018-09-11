package main

import "os"
import "fmt"
import "bufio"
import "strings"
import "8puzzle/8puzzle"

run := true
stateSpace := 8puzzle.8puzzle {
	boardConfig: {0, 1, 2, 3, 4, 5, 6, 7, 8}
}

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	inputBuffer := bufio.NewReader(os.Stdin)
	fmt.Println("mother love me long time long time")

	while run == true {
		inputCommand, _ := inputBuffer.ReadString('\n')
		inputCommand = strings.Replace(inputCommand, "\n", "", -1)
		inputCommand = strings.Replace(inputCommand, "\r\n", "", -1)
		inputSlice = strings.Split(inputCommand, " ")

		if strings.Compare("setState", inputSlice[0]) == 0 {
			stateSpace.setState(inputSlice[1])
		}

		if strings.Compare("randomizeState", inputSlice[0]) == 0 {

		}

		if strings.Compare("printState", inputSlice[0]) == 0 {

		}

		if strings.Compare("move", inputSlice[0]) == 0 {
			if strings.Compare("left", inputSlice[1]) == 0 {
				fmt.Println(stateSpace.left())
			}

			if strings.Compare("right", inputSlice[1]) == 0 {
				fmt.Println(stateSpace.right())
			}

			if strings.Compare("up", inputSlice[1]) == 0 {
				fmt.Println(stateSpace.up())
			}

			if strings.Compare("down", inputSlice[1]) == 0 {
				fmt.Println(stateSpace.down())
			}

		}

		if strings.Compare("solve", inputSlice[0]) == 0 {
			if strings.Compare("A-star", inputSlice[1]) == 0 {

			}

			if strings.Compare("beam", inputSlice[1]) == 0 {

			}
		}

		if strings.Compare("maxNodes", inputSlice[0]) {

		}

		if strings.Comapre("exit", inputSlice[0]) {
			run = false
		}

	}

}

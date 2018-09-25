package main

import "os"
import "fmt"
import "bufio"
import "strings"
import "math"
import "ePuzzle"

type fn func([9]int, [9]int)

func main() {
	focusObject := ePuzzle{0, 1, 2, 3, 4, 5, 6, 7, 8}
	x := ""
	file, err := os.Open(os.Args[1])

	fmt.Println("hell9")

	if err != nil {
		fmt.Println("failed to read argument file.")
	}

	defer file.Close()
	lineRead := bufio.NewScanner(file)

	for lineRead.Scan() {
		x = lineRead.Text()
		fmt.Println(lineRead.Text())

		if strings.Compare("setState", x) == 0 {

		}

		if strings.Compare("randomizeState", x) == 0 {

		}

		if strings.Compare("printState", x) == 0 {

		}

		if strings.Compare("move", x) == 0 {
			if strings.Compare("up", x) == 0 {

			}

			if strings.Compare("down", x) == 0 {

			}

			if strings.Compare("left", x) == 0 {

			}

			if strings.Compare("right", x) == 0 {

			}
		}

		if strings.Compare("solve", x) == 0 {
			if strings.Compare("solve A-star", x) == 0 {

			}

			if strings.Compare("beam", x) == 0 {

			}
		}

		if strings.Compare("maxNodes", x) == 0 {

		}
	}

	if err := lineRead.Err(); err != nil {
		fmt.Println("failed to parse argument file.")
	}

}

//return the value of tiles misplaced
func h1(state [9]int) int {
	misplacedTiles := 0
	winCondition := {0, 1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < 9; i++ {
		if winCondition[i] == state[i] {
			misplacedTiles++
		}
	}

	return misplasedTiles
}

//for h2, add the result of the number divided by 3, as well as the remainder to get the position
func h2(state [9]int) int {
	totalDistance := 0

	for i:=0; i < 9; i++ {
		totalDistance = totalDistance + math.Abs(math.Floor((state[i] - i) / 3))) + math.Abs((state[i] - i) % 3)
	}

	return totalDistance
}

func beamHeuristic(state [9]int, move [9]int) int {

}

package main

import "os"
import "fmt"
import "bufio"
import "strings"
//import "log"
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

func aStarHeuristic(state [9]int, move [9]int) {

}

func beamHeuristic(state [9]int, move [9]int) {

}

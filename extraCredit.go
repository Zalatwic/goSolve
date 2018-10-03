package main

import "os"
import "fmt"
import "bufio"
import "strings"
import "math"
import "time"
import "strconv"
import "math/rand"
import "BigPuzzle"
import "Beam"
import "Astar"

type fn func([]int, []int)

func main() {
	rand.Seed(time.Now().UnixNano())
	var focusObject BigPuzzle.BigPuzzle
	focusObject.SetState([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	file, err := getCommands(os.Args[1])
	maxNodes := 5800

	if err != nil {
		fmt.Println("failed to read argument file.")
	}

	for A, command := range file {
		if strings.Compare("setState", command) == 0 {

		}

		if strings.Compare("randomizeState", command) == 0 {
			//fisher-yates shuffle
			randomState := focusObject.GetState()

			for i := 0; i < 7; i++ {
				num := rand.Intn(12)
				randomState[num], randomState[i] = randomState[i], randomState[num]

			}

			printNice(randomState)
			focusObject.SetState(randomState)
		}

		if strings.Compare("printState", command) == 0 {
			printNice(focusObject.GetState())
		}

		if strings.Compare("move", command) == 0 {
			A++
			if strings.Compare("up", file[A]) == 0 {
				focusObject.Up(0)
			}

			if strings.Compare("down", file[A]) == 0 {
				focusObject.Down(0)
			}

			if strings.Compare("left", file[A]) == 0 {
				focusObject.Left(0)
			}

			if strings.Compare("right", file[A]) == 0 {
				focusObject.Right(0)
			}
		}

		if strings.Compare("solve", command) == 0 {
			fmt.Println(file[A+1])
			A++
			if strings.Compare("A-star", file[A]) == 0 {
				A++
				if strings.Compare("h1", file[A]) == 0 {
					aiInst := Astar.Astar{maxNodes, GetNames()}
					aiInst.SolvePuzzle(focusObject.GetState(), h1, focusObject.SetState, focusObject.GetStates)
					printNice(focusObject.GetState())
				}
				if strings.Compare("h2", file[A]) == 0 {
					aiInst := Astar.Astar{maxNodes, GetNames()}
					aiInst.SolvePuzzle(focusObject.GetState(), h2, focusObject.SetState, focusObject.GetStates)
					printNice(focusObject.GetState())
				}
			}

			if strings.Compare("beam", file[A]) == 0 {
				fmt.Println("works")

				s, err := strconv.Atoi(file[A+1])

				if err != nil {
					fmt.Println("atoi error")
				}

				aiInst := Beam.Beam{s, maxNodes, GetNames()}
				aiInst.SolvePuzzle(focusObject.GetState(), h2, focusObject.SetState, focusObject.GetStates)
				printNice(focusObject.GetState())
			}
		}

		if strings.Compare("maxNodes", command) == 0 {
			s, err := strconv.Atoi(file[A+1])

			if err != nil {
				fmt.Println("atoi error")
			}

			maxNodes = s
		}
	}
}

//return the value of tiles misplaced
func h1(state []int) int {
	misplacedTiles := 0
	winCondition := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 12, 10, 11}

	for i := 0; i < 12; i++ {
		if winCondition[i] != state[i] {
			misplacedTiles++
		}
	}

	return misplacedTiles
}

//for h2, add the result of the number divided by 3, as well as the remainder to get the position
func h2(state []int) int {
	totalDistance := 0

	for i := 0; i < 12; i++ {
		totalDistance = totalDistance + int(math.Abs(math.Floor(float64(state[i]-i)/3))+math.Abs(float64((state[i]-i)%3)))
	}
	return totalDistance
}

func getCommands(filePath string) ([]string, error) {
	commands, errorCode := os.Open(filePath)
	if errorCode != nil {
		return nil, errorCode
	}

	defer commands.Close()

	fileScanner := bufio.NewScanner(commands)
	fileScanner.Split(bufio.ScanWords)

	var commandList []string

	for fileScanner.Scan() {
		commandList = append(commandList, fileScanner.Text())
	}

	return commandList, nil
}

func GetNames() []string {
	moveNames := []string{"left", "right", "up", "down"}
	return moveNames
}

func printNice(data []int) {
	for i := 0; i < 12; i++ {
		if (i % 3) == 0 {
			fmt.Println("")
		}
		if data[i] != 0 {
			fmt.Print(data[i])
		}
		if data[i] == 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

package main

import "os"
import "fmt"
import "bufio"
import "strings"

run := true

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

		}

		if strings.Compare("randomizeState", inputSlice[0]) == {

		}

		if strings.Compare("printState", inputSlice[0]) {

		}

		if strings.Compare("move", inputSlice[0]) {

		}

		if strings.Compare("solve", inputSlice[0]) {
			if strings.Compare("A-star", inputSlice[1]) {

			}

			if strings.Compare("beam", inputSlice[1]) {

			}

			else {

			}
		}

		if strings.Compare("maxNodes", inputSlice[0]) {

		}

		if strings.Comapre("exit", inputSlice[0]) {
			run = false
		}

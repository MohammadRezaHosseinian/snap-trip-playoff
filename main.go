package main

import "fmt"

const (
	TeamOne = iota
	TeamTwo
	Nothing
)

func main() {
	totalPlayoffBoard := make(map[int]int)
	var gameCounts int
	//fmt.Println("Please enter counts of game ")
	fmt.Scanf("%d", &gameCounts)
	for i := 0; i < gameCounts; i++ {
		onePlayOff := getInputs()
		totalPlayoffBoard[calcWinner(onePlayOff)]++
	}
	fmt.Printf("%+v", totalPlayoffBoard)
}

func calcWinner(board [4]int) int {
	golOne := board[0] + board[2]
	golTwo := board[1] + board[3]
	if golOne > golTwo {
		fmt.Println("perspolis")
		return TeamOne
	}
	if golOne < golTwo {
		fmt.Println("esteghlal")
		return TeamTwo
	}
	if board[1] > board[2] {
		fmt.Println("esteghlal")
		return TeamTwo
	}
	if board[2] > board[1] {
		fmt.Println("perspolis")
		return TeamOne
	}
	fmt.Println("penalty")
	return Nothing
}

func getInputs() [4]int {
	//fmt.Println("please enter your inputs")
	var golCounts [4]int
	for i := 0; i < 4; i++ {
		fmt.Scanf("%d", &golCounts[i])
	}
	return golCounts
}

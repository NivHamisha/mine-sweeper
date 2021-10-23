package main

import (
	"fmt"
	"mine-sweeper/utils"
)

func main()  {
	if err := manageGame(10,2); err != nil {
		panic(err)
	}
}

func manageGame(boardSize, minesNumber int) error {
	board, err := utils.NewBoard(boardSize, minesNumber)
	if err != nil {
		panic(err)
	}

	var i, j int
	for {
		board.PrintBoardGameView()
		fmt.Println("Enter row: ")

		if _, err = fmt.Scanln(&i);err != nil {
			return err
		}
		fmt.Println("Enter column: ")
		if _, err = fmt.Scanln(&j);err != nil {
			return err
		}
		if !board.ManagePress(i,j){
			fmt.Println("Mine pressed, you lost :(")
			return nil
		}
		if board.IsWin(){
			fmt.Println("You Won :)")
			return nil
		}
	}
	return nil
}
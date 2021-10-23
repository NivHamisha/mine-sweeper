package utils

import (
	"fmt"
	"math/rand"
	"time"
)

type cellLocation [2]int

type board struct {
	size int
	mines int
	boardMatrix [][]Cell
	presses int
}

func NewBoard(size, mines int) (*board, error) {
	if mines > size*size{
		return nil, fmt.Errorf("cannot create board with mines greater than size")
	}
	matrixSize := size+2
	boardMatrix := make([][]Cell, matrixSize)
	for i := range boardMatrix {
		boardMatrix[i] = make([]Cell, matrixSize)
	}

	var randomList []cellLocation
	for i:=1; i<matrixSize-1; i++ {
		for j := 1; j < matrixSize-1; j++ {
			randomList = append(randomList, cellLocation{i,j})
		}
	}

	min := 0
	for i:=0; i<mines; i++ {
		max := len(randomList)
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(max - min) + min
		cell := randomList[randomIndex]
		boardMatrix[cell[0]][cell[1]] = Cell{value: -1}
		randomList = append(randomList[:randomIndex], randomList[randomIndex+1:]...)
	}

	cellDirections := GetAllCellDirections()
	for i:=0; i< len(boardMatrix) ; i++ {
		for j := 0; j < len(boardMatrix[i]); j++ {
			if i == 0 || j == 0 || j == len(boardMatrix[i]) - 1 || i == len(boardMatrix) - 1{
				boardMatrix[i][j].UpdateValue(-2)
			}else {
				if !boardMatrix[i][j].IsMine() {
					numberOfMines := 0
					for _, cellDirection := range cellDirections {
						if boardMatrix[i+cellDirection[0]][j+cellDirection[1]].IsMine() {
							numberOfMines = numberOfMines + 1
						}
					}
					boardMatrix[i][j].UpdateValue(numberOfMines)
				}
			}
		}
	}

	return &board{size: size, mines: mines,boardMatrix: boardMatrix}, nil
}

func (b *board) PrintBoardWithSafetyWall(){
	for i:=0; i<len(b.boardMatrix); i++ {
		for j:=0; j<len(b.boardMatrix); j++ {
			value := b.boardMatrix[i][j].value
			if value == -1 {
				fmt.Print(fmt.Sprintf("x "))
			} else {
				fmt.Print(fmt.Sprintf("%d ", value))
			}
		}
		fmt.Println()
	}
}

func (b *board) PrintBoard(){
	for i:=1; i<len(b.boardMatrix)-1; i++ {
		for j:=1; j<len(b.boardMatrix)-1; j++ {
			value := b.boardMatrix[i][j].value
			if value == -1 {
				fmt.Print(fmt.Sprintf("x "))
			} else {
				fmt.Print(fmt.Sprintf("%d ", value))
			}
		}
		fmt.Println()
	}
}

func (b* board) ManagePress(i, j int)bool{
	if b.boardMatrix[i][j].isPressed{
		return true
	}
	if b.boardMatrix[i][j].IsMine(){
		return false
	}
	if b.boardMatrix[i][j].value != 0{
		b.boardMatrix[i][j].Press()
		b.presses = b.presses + 1
	}else {
		b.Extend(i,j)
	}
	return true
}

func (b *board) Extend (i,j int){
	if b.boardMatrix[i][j].isPressed {
		return
	}
	if b.boardMatrix[i][j].value != -2  && b.boardMatrix[i][j].value != 0{
		b.boardMatrix[i][j].Press()
		b.presses  = b.presses + 1
		return
	}
	if b.boardMatrix[i][j].value != 0 {
		return
	}

	extendDirections := GetAllCellDirections()
	b.boardMatrix[i][j].Press()
	b.presses  = b.presses + 1
	for _, extendDirection:= range extendDirections{
		b.Extend(i+extendDirection[0], j+extendDirection[1])
	}
}

func (b *board) PrintBoardGameView(){
	fmt.Printf("    ")
	for i:=1; i<len(b.boardMatrix)-1; i++ {
		fmt.Printf("%-2d", i)
	}
	fmt.Println()

	for i:=1; i<len(b.boardMatrix)-1; i++ {
		fmt.Printf("%-2d: ", i)
		for j:=1; j<len(b.boardMatrix)-1; j++ {
			if b.boardMatrix[i][j].isPressed{
				if b.boardMatrix[i][j].IsMine(){
					fmt.Printf("x ")
				} else {
					fmt.Printf("%d ", b.boardMatrix[i][j].value)
				}
			}else {
				fmt.Printf("- ")
			}
		}
		fmt.Println()
	}
}

func (b *board) IsWin() bool {
	if b.presses + b.mines == b.size*b.size{
		return true
	}
	return false
}

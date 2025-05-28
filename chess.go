package main

import "fmt"

type array [8][8]string

func display(chessboard array) {

	for i := 0; i < len(chessboard); i++ {
		for j := 0; j < len(chessboard[i]); j++ {
			fmt.Printf("%v ", chessboard[i][j])
		}
		fmt.Print("\n")
	}
}

func chess() {
	var board [8][8]string

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if i == 0 {
				if j == 0 || j == len(board[i])-1 {
					board[i][j] = "r"
				} else if j == 1 || j == len(board[i])-2 {
					board[i][j] = "n"
				} else if j == 2 || j == len(board[i])-3 {
					board[i][j] = "b"
				} else if j == 3 {
					board[i][j] = "q"
				} else {
					board[i][j] = "k"
				}
			} else if i == len(board)-1 {
				if j == 0 || j == len(board[i])-1 {
					board[i][j] = "R"
				} else if j == 1 || j == len(board[i])-2 {
					board[i][j] = "N"
				} else if j == 2 || j == len(board[i])-3 {
					board[i][j] = "B"
				} else if j == 3 {
					board[i][j] = "Q"
				} else {
					board[i][j] = "K"
				}
			} else {
				board[i][j] = "_"
			}
		}
	}
	display(board)
}

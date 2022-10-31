package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func choose_board() bool {
	end := false
	count := 0
	board := map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9"}
	win_combinations := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, a := range win_combinations {
		fmt.Println("win_comb:- ", a[0], a[1], a[2])
		if board[a[0]] == board[a[1]] && board[a[1]] == board[a[2]] && strconv.FormatBool(board[a[1]] == board[a[2]]) == "X" {
			fmt.Println("Player 1 Wins! \n")
			fmt.Println("Congratulations! \n")
			return true
		}
		if board[a[0]] == board[a[1]] && board[a[1]] == board[a[2]] && strconv.FormatBool(board[a[1]] == board[a[2]]) == "O" {
			fmt.Println("Player 2 Wins! \n")
			fmt.Println("Congratulations! \n")
			return true
		}
		for a := 0; a < 9; a++ {
			if board[a] == "X" || board[a] == "O" {
				count += 1
			}
			if count == 9 {
				fmt.Println("The game ends in a Tie\n")
				return true
			}
		}
	}
	for !end {
		// draw()
		// end = check_board()
		if end == true {
			break
		}
		fmt.Println("Player 1 choose where to place a cross")
		p1()
		fmt.Println()
		draw()

		// end = check_board()
		if end == true {
			break
		}
		fmt.Println("Player 2 choose where to place a nought")
		p2()
		fmt.Println()
	}
	if func(msg string) string {
		fmt.Println(msg)
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		return strings.ReplaceAll(text, "\n", "")
	}("Play again (y/n)\n") == "y" {
		fmt.Println()
		tic_tac_toe()
	}
}

// func main() {
// 	choose_board()
// }

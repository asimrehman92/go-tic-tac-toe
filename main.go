package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// logger := log.WithFields(log.Fields{
	// 	"Project": "Tic Tac Toe",
	// })

	// var w1, w2, w3 string
	// log.Info("Please Enter Player 1 name: ")
	// fmt.Scan(&w1, &w2, &w3)
	// log.WithFields(log.Fields{
	// 	"w1": w1,
	// 	"w2": w2,
	// 	"w3": w3,
	// })

}

func tic_tac_toe() {
	// declaration
	end := false
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

	// declare function into a variable
	draw := func() {
		fmt.Println(board[1], board[2], board[3])
		fmt.Println(board[4], board[5], board[6])
		fmt.Println(board[7], board[8], board[9])
		fmt.Println()
	}

	// choose a number
	choose_number := func() int {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					if strings.HasPrefix(err.Error(), "ValueError") {
						fmt.Println("\nThat's not a number. Try again")
						return
					}
				}
				panic(r)
			}
		}()

		var numb int
		fmt.Println("Please Enter any number:- ")
		_, err := fmt.Scan(&numb)
		if err != nil {
			panic(err)
		}
		fmt.Println(numb)

		if numb != -1 {
			for v := 0; v < 9; v++ {
				if v == numb {
					return numb
				}
			}
		} else {
			fmt.Println("\nThat's not on the board. Try again")
		}

		return numb
	}

	// player 1
	p1 := func() {
		n := choose_number()
		if board[n] == "X" || board[n] == "O" {
			fmt.Println("\nYou can't go there, Try again")
		} else {
			board[n] = "X"
		}
	}

	// player 2
	p2 := func() {
		n := choose_number()
		if board[n] == "X" || board[n] == "O" {
			fmt.Println("\nYou can't go there. Try again")
		} else {
			board[n] = "O"
		}
	}

	// check the tic tac toe board for game decision
	var check_board func() bool
	check_board = func() bool {
		count := 0

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
			draw()
			end = check_board()
			if end == true {
				break
			}
			fmt.Println("Player 1 choose where to place a cross")
			p1()
			fmt.Println()
			draw()

			end = check_board()
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

		return false
	}

	draw()
	p1()
}

func init() {
	tic_tac_toe()

	// fmt.Println("This is rehman")
}

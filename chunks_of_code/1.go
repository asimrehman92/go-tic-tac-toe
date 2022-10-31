package main

import (
	"fmt"
	"strings"
)

func choose_numberrr() int {
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
	return 0
}

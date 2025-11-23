package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: passgen [-s] <len>")
		return

		return
	}

	if args[0] == "-h" {
		fmt.Println("Usage: passgen [-s] <len>")
		return
	}

	var useSymbols bool
	var passlen int
	var err error
	if args[0] == "-s" {
		passlen, err = strconv.Atoi(args[1])
		useSymbols = true
	} else {
		passlen, err = strconv.Atoi(args[0])
		useSymbols = false
	}

	if err != nil {
		log.Fatalln("Please provide a number")
		return
	}

	pass := genPass(passlen, useSymbols)
	fmt.Println("*****************************************")
	fmt.Println(pass)
	fmt.Println("*****************************************")
}

func genPass(passlen int, useSymbols bool) string {
	var password string
	currentLen := 0

	base := 6
	step := 5
	i := passlen / step
	for {
		if currentLen >= passlen {
			break
		}
		l := getLetter(useSymbols)
		if len(password) < base+step*i && strings.Count(password, l) >= i+1 {
			continue
		}
		password += l
		currentLen += 1
	}

	return password
}

func getLetter(useSymbols bool) string {
	letters := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	numbers := [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	symbols := [...]string{"!", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", ":", ";", "<", "=", ">", "?", "@", "[", "]", "^", "_", "{", "}", "|", "~"}

	var swithBy int
	if useSymbols {
		swithBy = rand.IntN(7)
	} else {
		swithBy = rand.IntN(6)
	}

	var s string
	switch swithBy {
	case 0, 1:
		s = letters[rand.IntN(len(letters))]
	case 2, 3:
		s = letters[rand.IntN(len(letters))]
		s = strings.ToUpper(s)
	case 4, 5:
		s = numbers[rand.IntN(len(numbers))]
	case 6:
		s = symbols[rand.IntN(len(symbols))]
	}
	return s
}

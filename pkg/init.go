package tetromino

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadInput(file io.Reader) [][4][4]string {
	var tetrominoArray [][4][4]string
	var tetromino [4][4]string
	scanner := bufio.NewScanner(file)
	i, in, flag, alpha := 0, 0, true, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			if flag {
				flag = false
				continue
			} else {
				fmt.Println("ERROR")
				os.Exit(1)
			}
		}
		var arr [4]string
		if len(str) != 4 {
			fmt.Println("ERROR")
			os.Exit(1)
		}
		for ind := range arr {
			if rune(str[ind]) == '.' {
				arr[ind] = "."
			} else if rune(str[ind]) == '#' {
				arr[ind] = string(alpha[in])
			} else {
				fmt.Println("ERROR")
				os.Exit(1)
			}
		}
		tetromino[i] = arr
		i++
		if i == 4 {
			flag = true
			i = 0
			in++
			if !CheckTetro(tetromino) {
				fmt.Println("ERROR")
				os.Exit(1)
			}
			tetromino = OptimizeTetromino(tetromino)
			tetrominoArray = append(tetrominoArray, tetromino)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return tetrominoArray
}
func PrintInputs(myArray [][4][4]string) {
	//Prints inputted tetrominoes
	for _, tetromino := range myArray {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				fmt.Printf("%s ", tetromino[j][k])
			}
			fmt.Printf("\n")
		}
		fmt.Println()
	}
}
func InitSquare(n int) [][]string {
	//initializes a square
	var Square [][]string
	var row []string
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		Square = append(Square, row)
		row = []string{}
	}
	return Square
}

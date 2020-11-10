package tetromino

import (
	"fmt"
	"math"
)

var mySquare [][]string

func BacktrackSolver(tetrominoes [][4][4]string, n int) bool {
	if n == len(tetrominoes) { //base condition when all tetrominoes are placed, board is solved
		return true
	}
	for i := 0; i < len(mySquare); i++ {
		for j := 0; j < len(mySquare); j++ {
			if CheckInsert(i, j, tetrominoes[n]) { //check if we can place current tetrominoe on the board anywhere
				Insert(i, j, tetrominoes[n]) // if we can place it at this location, check if we can place another piece
				if BacktrackSolver(tetrominoes, n+1) {
					return true
				}
				Remove(i, j, tetrominoes[n]) //if the next piece can't be placed, backtrack
			}
		}
	} // if we can't place tetro anywhere, return false
	return false
}

func Insert(i, j int, tetro [4][4]string) { // insert piece and when all 4 piecec "#" are placed, no need to place '.'
	a, b, c := 0, 0, 0
	for a < 4 {
		for b < 4 {
			if tetro[a][b] != "." {
				c++
				mySquare[i+a][j+b] = tetro[a][b]
				if c == 4 {
					break
				}
			}
			b++
		}
		b = 0
		a++
	}
}

func Remove(i, j int, tetro [4][4]string) { //remove piece at current location
	a, b, c := 0, 0, 0
	for a < 4 {
		for b < 4 {
			if tetro[a][b] != "." {
				if c == 4 {
					break
				}
				mySquare[i+a][j+b] = "."
			}
			b++
		}
		b = 0
		a++
	}
}

func Solve(tetrominoes [][4][4]string) [][]string {
	//initial board starts with dimmension 4*4, if we can't place all tetrominoes
	//increase size by 1 and initialize board
	l := int(math.Ceil(math.Sqrt(float64(4 * len(tetrominoes)))))
	// if l < 4 {
	// 	l = 4
	// }
	mySquare = InitSquare(l)
	for !BacktrackSolver(tetrominoes, 0) {
		l++
		mySquare = InitSquare(l)
	}
	//BacktrackSolver(tetrominoes, 0)
	return mySquare
}

func PrintSol() {
	for i := range mySquare {
		for j := range mySquare {
			fmt.Printf("%s ", mySquare[i][j])
		}
		fmt.Printf("\n")
	}
}

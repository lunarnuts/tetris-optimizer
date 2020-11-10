package tetromino

func CheckTetro(tetromino [4][4]string) bool {
	c, d := 0, 0

	for a, elem := range tetromino {
		for b, elem2 := range elem {
			if elem2 != "." {
				d++
				if a+1 < 4 && tetromino[a+1][b] != "." {
					c++
				}
				if a-1 >= 0 && tetromino[a-1][b] != "." {
					c++
				}
				if b+1 < 4 && tetromino[a][b+1] != "." {
					c++
				}
				if b-1 >= 0 && tetromino[a][b-1] != "." {
					c++
				}
			}
		}
	}
	if d != 4 {
		return false
	}
	if c == 6 || c == 8 {
		return true
	}
	return false
}
func CheckInsert(i, j int, tetro [4][4]string) bool { //check if we can place piece at current location
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			if tetro[a][b] != "." {
				if i+a == len(mySquare) || j+b == len(mySquare) || mySquare[i+a][j+b] != "." {
					return false
				}
			}
		}

	}
	return true
}

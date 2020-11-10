package tetromino

func OptimizeTetromino(tetromino [4][4]string) [4][4]string {
	//optimzes tetromino
	i := 0
	for {
		zeroes := 0
		for j := 0; j < 4; j++ {
			if tetromino[i][j] == "." {
				zeroes++
			}
		}
		if zeroes == 4 { //if row is all zeroes, shift by 1 row to top
			tetromino = ShiftVertical(tetromino)
			continue
		}
		break
	}
	for {
		zeroes := 0
		for j := 0; j < 4; j++ {
			if tetromino[j][i] == "." {
				zeroes++
			}
		}
		if zeroes == 4 { //if col is all zeroes, shift by 1 col to left
			tetromino = ShiftHorizontal(tetromino)
			continue
		}
		break
	}
	return tetromino
}
func ShiftVertical(tetromino [4][4]string) [4][4]string {
	//shifts tetromino row by 1
	temp := tetromino[0]
	tetromino[0] = tetromino[1]
	tetromino[1] = tetromino[2]
	tetromino[2] = tetromino[3]
	tetromino[3] = temp
	return tetromino
}
func ShiftHorizontal(tetromino [4][4]string) [4][4]string {
	//shifts tetromino col by 1
	tetromino = Transpose(tetromino)
	tetromino = ShiftVertical(tetromino)
	tetromino = Transpose(tetromino)
	return tetromino
}
func Transpose(slice [4][4]string) [4][4]string {
	//transpose tetromino
	xl := len(slice[0])
	yl := len(slice)
	var result [4][4]string
	for i := range result {
		result[i] = [4]string{}
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

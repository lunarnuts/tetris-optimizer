package main

import (
	"fmt"
	"os"

	tetromino "./pkg"
)

func main() {
	args := os.Args[1]
	if args != "" {
		file, err := os.Open(args)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}()
		var myArray = tetromino.ReadInput(file)
		tetromino.Solve(myArray)
		tetromino.PrintSol()
	}
}

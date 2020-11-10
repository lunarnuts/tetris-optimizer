# tetris-optimizer

## Description

This program will take a file as parameter, which contains a list of Tetriminos, and arrange them in order to create the smallest square possible.

### The program

The executable "tetris-optimizer" takes only one parameter, a file which contains a list of Tetriminos to assemble. This file has a very specific format (see input sample).

### Input Sample:

Each tetraminos is represented with 4 lines of 4 characters, each followed by a new line. A Tetrimino is a classic piece of Tetris composed of 4 blocks. Each character must be either a block character(’#’) or an empty character (’.’). Each block of a Tetrimino must touch at least one other block on any of his 4 sides (up, down, left and right).

    ...#
    ...#
    ...#
    ...#

    ....
    ....
    ....
    ####

    .###
    ...#
    ....
    ....

    ....
    ..##
    .##.
    ....

    ....
    .##.
    .##.
    ....

    ....
    ....
    ##..
    .##.

    ##..
    .#..
    .#..
    ....

    ....
    ###.
    .#..
    ....

### Output Sample:

The program displays the smallest assembled square on the standard output. To
identify each Tetrimino in the square solution, a capital letter will be assigned to each
Tetrimino, starting with ’A’ and increasing for each new Tetrimino.

    ABBBB.
    ACCCEE
    AFFCEE
    A.FFGG
    HHHDDG
    .HDD.G

# Solution

The program is divided in two parts:A part that reads, validate, optimize the input and a second part that solves the problem using backtracking.
here is how the program works:

### Solving the problem

- Reads the input file and optimizes tetrominoes by arranging significant "bits" on top-left corner so that:

```
....
.##.
.##.
....
```

becomes :

```
##..
##..
....
....
```

- Generates a board of size 4 with the fucntion InitBoard(4).
- Calls the function BacktrackSolver() with an Array of Optimized tetrominoes and index of tetrominoe piece. Initial call is BacktrackSolver(tetrominoes,0). If solver checked all possible placements but couldn't yield any results, we increase the size of board by 1 and reinitialized matrix and call BacktrackSolver again.
- Before placing a tetromino, it is checked if it is legal to do so by function CheckInsert()
- The solution is printed by PrintSol()

### Note

to test the programm with given bad and good input files, you can run test.sh (don't forget to set privileges to exec by
`chmod u+x test.sh`)

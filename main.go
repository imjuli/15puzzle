package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
)

const (
	sizeMsg           = "Enter desired puzzle size: "
	newGameMsg        = "\nWelcome! Your game has started. Print the tile number to make a move, q to quit, n to start a new game. Good luck!"
	invalidInputMsg   = "Unrecognized input, try q, n or tile number"
	invalidIntegerMsg = "Invalid size value. Try positive integers"
	invalidTileMsg    = "Tile does not exist, try numbers from 1 to"
	invalidMoveMsg    = "Move is not allowed. Valid moves:"
	wonMsg            = "🎉🎉🎉 Congratulations! You won! 🎉🎉🎉 Press n to start a new game or q to quit\n "
	quitKey           = "q"
	newGameKey        = "n"
)

var (
	board       = [][]int{}
	size        int
	targetSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}
)

func readSize() {
	var sizeInput int
	fmt.Print(sizeMsg)
	for {
		_, err := fmt.Scan(&sizeInput)
		if err != nil || sizeInput <= 0 {
			fmt.Println(invalidIntegerMsg)
			continue
		} else {
			break
		}
	}
	size = sizeInput
}

func main() {

	var (
		input   string
		tileNum int
	)

	readSize()
	newGame(size)

	for {
		fmt.Print("> ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
		}
		if input == quitKey {
			os.Exit(0)
		} else if input == newGameKey {
			readSize()
			newGame(size)
		} else {
			tileNum, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println(invalidInputMsg)
				continue
			}
			if tileNum <= 0 || tileNum > size*size-1 {
				fmt.Println(invalidTileMsg, size*size-1)
				continue
			}
			if vm := validMoves(board, size); !contains(vm, tileNum) {
				fmt.Println(invalidMoveMsg, vm)
				continue
			}

			board = move(tileNum, board)
			drawBoard()

			if won() {
				fmt.Println(wonMsg)
			}
		}
	}
}

func newGame(size int) {
	fmt.Println(newGameMsg)
	board = toBoard(genCombination(size), size)
	drawBoard()
}

func genCombination(size int) []int {
	var c []int
	for {
		c = rand.Perm(size * size)
		if isSolvable(c, size) {
			break
		}
	}
	return c
}

func toBoard(nums []int, size int) [][]int {
	b := make([][]int, size)
	for i := range b {
		b[i] = make([]int, size)
		for j := range b[i] {
			b[i][j] = nums[size*i+j]
		}
	}
	return b
}

func drawBoard() {
	format := `%-` + strconv.Itoa(len(strconv.Itoa(size*size))+1) + "v"
	fmt.Println()
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				fmt.Printf(format, " ")
			} else {
				fmt.Printf(format, board[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func position(num int, board [][]int) (x, y int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == num {
				return i, j
			}
		}
	}
	return -1, -1
}

func validMoves(b [][]int, size int) []int {
	x, y := position(0, b)
	moves := []int{}
	if x < 0 || y < 0 {
		return moves
	}
	if x != 0 {
		moves = append(moves, b[x-1][y])
	}
	if x != size-1 {
		moves = append(moves, b[x+1][y])
	}
	if y != 0 {
		moves = append(moves, b[x][y-1])
	}
	if y != size-1 {
		moves = append(moves, b[x][y+1])
	}
	return moves
}

func contains(sl []int, e int) bool {
	for _, v := range sl {
		if v == e {
			return true
		}
	}
	return false
}

func move(num int, b [][]int) [][]int {
	xN, yN := position(num, b)
	x0, y0 := position(0, b)
	b[xN][yN], b[x0][y0] = b[x0][y0], b[xN][yN]
	return b
}

func toSlice(b [][]int) []int {
	sl := []int{}
	for _, v := range b {
		sl = append(sl, v...)
	}
	return sl
}

func inversionCount(sl []int) int {
	var count int
	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[j] != 0 && sl[i] != 0 && sl[i] > sl[j] {
				count++
			}
		}
	}
	return count
}

/*
Solvability is calculated by the following rules:
If N is odd, then puzzle instance is solvable if number of inversions is even
If N is even, puzzle instance is solvable if:
- the blank is on an even row counting from the bottom (second-last, fourth-last, etc.) and number of inversions is odd.
- the blank is on an odd row counting from the bottom (last, third-last, fifth-last, etc.) and number of inversions is even.
*/

func isSolvable(sl []int, size int) bool {
	invCount := inversionCount(sl)
	if size%2 == 1 {
		return invCount%2 == 0
	} else {
		b := toBoard(sl, size)
		x0, _ := position(0, b)
		x0b := size - x0 // row of the blank from the bottom
		if x0b%2 == 1 {
			return invCount%2 == 0
		} else {
			return invCount%2 == 1
		}
	}
}

func won() bool {
	return reflect.DeepEqual(toSlice(board), targetSlice)
}

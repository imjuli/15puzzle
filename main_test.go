package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var targetBoard = [][]int{
	{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 10, 11, 12},
	{13, 14, 15, 0}}

func TestToBoard(t *testing.T) {
	assert.Equal(t, targetBoard, toBoard(targetSlice))
}

func TestToSlice(t *testing.T) {
	assert.Equal(t, targetSlice, toSlice(targetBoard))
}

type invCountTest struct {
	combination   []int
	expectedCount int
}

var invCombos = []invCountTest{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 0, 15}, 0},
	{[]int{2, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}, 1},
	{[]int{1, 0, 5, 2, 4, 3}, 4},
}

func TestInvCount(t *testing.T) {
	for _, test := range invCombos {
		if count := inversionCount(test.combination); count != test.expectedCount {
			t.Errorf("Puzzle %v : inversion count [%d] is not equal to expected [%d]", test.combination, count, test.expectedCount)
		}
	}
}

type solvableTest struct {
	combination []int
	size        int
	expected    bool
}

var solvableTests = []solvableTest{
	// known unsolvable combos
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 14, 0}, 4, false},
	{[]int{1, 2, 3, 4, 5, 6, 8, 7, 0}, 3, false},
	{[]int{3, 9, 1, 15, 14, 11, 4, 6, 13, 0, 10, 12, 2, 7, 8, 5}, 4, false},
	// known solvable combos
	{targetSlice, 4, true},
	{[]int{1, 8, 2, 0, 4, 3, 7, 6, 5}, 3, true},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 0, 15}, 4, true},
	{[]int{13, 2, 10, 3, 1, 12, 8, 4, 5, 0, 9, 6, 15, 14, 11, 7}, 4, true},
	{[]int{12, 1, 10, 2, 7, 11, 4, 14, 5, 0, 9, 15, 8, 13, 6, 3}, 4, true},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 11, 13, 14, 15, 12}, 4, true},
	{[]int{6, 13, 7, 10, 8, 9, 11, 0, 15, 2, 12, 5, 14, 3, 1, 4}, 4, true},
}

func TestIsSolvable(t *testing.T) {
	for _, test := range solvableTests {
		if solvable := isSolvable(test.combination, test.size); solvable != test.expected {
			t.Errorf("Puzzle %v : solvability [%t] is not equal to expected [%t]", test.combination, solvable, test.expected)
		}
	}
}

func TestGenSolvableCombos(t *testing.T) {
	var (
		combos [][]int
		i      int
	)
	for i >= 1 && i <= 10 {
		combos[i] = genCombination(i)
		assert.True(t, isSolvable(combos[i], i))
	}
}

type moveTest struct {
	num      int
	expBoard [][]int
}

var moveTests = []moveTest{
	{6, [][]int{
		{2, 1, 3},
		{5, 6, 7},
		{8, 0, 5}}},
	{7, [][]int{
		{2, 1, 3},
		{5, 7, 0},
		{8, 6, 5}}},
	{5, [][]int{
		{2, 1, 3},
		{0, 5, 7},
		{8, 6, 5}}},
	{1, [][]int{
		{2, 0, 3},
		{5, 1, 7},
		{8, 6, 5}}}}

func TestMoves(t *testing.T) {
	for _, test := range moveTests {
		initBoard := [][]int{
			{2, 1, 3},
			{5, 0, 7},
			{8, 6, 5}}
		movedBoard := move(test.num, initBoard)
		assert.Equal(t, test.expBoard, movedBoard, "Board: %v Invalid move result. Move: %d, expected: %v, actual: %v", initBoard, test.num, test.expBoard, movedBoard)
	}
}

type posAndValidMovesTest struct {
	combination   []int
	expValidMoves []int
	posX0         int
	posY0         int
}

var posAndValidMovesTests = []posAndValidMovesTest{
	{[]int{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 0, 15},
		[]int{11, 14, 15},
		3, 2},
	{[]int{
		2, 1, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 0},
		[]int{12, 15},
		3, 3},
	{[]int{
		2, 1, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 0,
		13, 14, 15, 12},
		[]int{8, 12, 11},
		2, 3},
	{[]int{
		2, 1, 3, 0,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 4},
		[]int{8, 3},
		0, 3},
	{[]int{
		2, 0, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 1},
		[]int{6, 2, 3},
		0, 1},
	{[]int{
		0, 1, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 2},
		[]int{5, 1},
		0, 0},
	{[]int{
		2, 1, 3, 4,
		5, 6, 7, 8,
		0, 10, 11, 12,
		13, 14, 15, 9},
		[]int{5, 13, 10},
		2, 0},
	{[]int{
		2, 1, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		0, 14, 15, 13},
		[]int{9, 14},
		3, 0},
	{[]int{
		2, 1, 3, 4,
		5, 6, 7, 8,
		9, 0, 11, 12,
		13, 14, 15, 10},
		[]int{6, 14, 9, 11},
		2, 1},

	{[]int{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16},
		[]int{},
		-1, -1},
}

func TestPositionAndValidMoves(t *testing.T) {
	for _, test := range posAndValidMovesTests {
		b := toBoard(test.combination)
		m := validMoves(b)
		x, y := position(0, b)
		if !reflect.DeepEqual(m, test.expValidMoves) || x != test.posX0 || y != test.posY0 {
			t.Errorf("Puzzle %v : valid moves: %v, expected %v; position of 0: %v %v, expected %v %v",
				test.combination, m, test.expValidMoves, x, y, test.posX0, test.posY0)
		}
	}
}

func TestWon(t *testing.T) {
	board = targetBoard
	assert.True(t, won())
	board = toBoard([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0})
	assert.False(t, won())
}

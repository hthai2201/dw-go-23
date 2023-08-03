package matrix

import (
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
	board := [][]int{
		{1, 0, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	expectedVisitedPairs := [][]int{
		{0, 0},
		{1, 0},
		{1, 1},
		{2, 1},
		{2, 2},
	}
	expectedVisited:= [][]bool{
		{ true, false, false},
		{true, true, false},
		{false, true, true},
	}

	visitedPairs := make([][]int, 0)
	visited := CreatBoolMatrix(len(board), len(board[0]))
	DFS(board, 0, 0, visited, &visitedPairs)

	if !reflect.DeepEqual(visitedPairs, expectedVisitedPairs) {
		t.Errorf("DFS did not produce the expected visited pairs.\nExpected: %v\nGot: %v", expectedVisitedPairs, visitedPairs)
	}
	if !reflect.DeepEqual(visited, expectedVisited) {
		t.Errorf("DFS did not produce the expected visited.\nExpected: %v\nGot: %v", expectedVisited, visited)
	}
}

func TestFindMinMaxIJ(t *testing.T) {
	visitedPairs := [][]int{
		{0, 2},
		{1, 0},
		{1, 1},
		{2, 1},
		{2, 2},
	}

	expectedLowestI, expectedHighestI := 0, 2
	expectedLowestJ, expectedHighestJ := 0, 2

	lowestI, highestI, lowestJ, highestJ := FindMinMaxIJ(visitedPairs)

	if lowestI != expectedLowestI || highestI != expectedHighestI || lowestJ != expectedLowestJ || highestJ != expectedHighestJ {
		t.Errorf("FindMinMaxIJ did not produce the expected results.\nExpected: (%d, %d, %d, %d)\nGot: (%d, %d, %d, %d)",
			expectedLowestI, expectedHighestI, expectedLowestJ, expectedHighestJ, lowestI, highestI, lowestJ, highestJ)
	}
}

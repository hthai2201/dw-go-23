package rectangle

import (
	"github.com/hthai2201/dw-go-23/exercises/04/matrix"
)


func isRectangle(board [][]int, lowestI, highestI, lowestJ, highestJ int) bool {
	for i := lowestI; i <= highestI; i++ {
		for j := lowestJ; j <= highestJ; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}


func CountRectangles(board [][]int) int {
	if len(board) == 0 {
		return 0
	}

	rows, cols := len(board), len(board[0])
	visited := matrix.CreatBoolMatrix(rows,cols)

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] != 0 && !visited[i][j] {
				var visitedPairs [][]int // Create an empty list to store the (i, j) pairs

				matrix.DFS(board, i, j,visited, &visitedPairs)
				// Find the lowest i, highest i, lowest j, and highest j values from the visitedPairs list
				lowestI, highestI, lowestJ, highestJ := matrix.FindMinMaxIJ(visitedPairs)
				if isRectangle(board, lowestI, highestI, lowestJ, highestJ){
				 	count++
				}
			}
		}
	}

	return count
}

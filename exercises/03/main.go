package main

import (
	"fmt"

	"github.com/hthai2201/dw-go-23/exercises/03/rectangle"
)




func main() {
	board := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 0, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	rectangles := rectangle.CountRectangles(board)
	fmt.Println("Number of rectangles on the board:", rectangles) 
}

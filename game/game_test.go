package game

import (
	"fmt"
	"testing"
)

func TestGetPathInOneDimensionalMaze(t *testing.T) {
	var maze []int
	for i := 0; i < 96; i++ {
		if i == 11 || i == 23 || i == 22 || i == 21 || i == 35 || i == 34 {
			maze = append(maze, 0)
			continue
		}
		maze = append(maze, 1)
	}
	ints, _ := GetPathInOneDimensionalMaze(maze, 8, 12, 11, 33)
	for _, i := range ints {
		fmt.Println(i)
	}
}

func TestAStar(t *testing.T) {
	grid := Grid{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}
	aStar := NewAStar(grid, Vector{0, 0}, Vector{4, 4})
	search := aStar.Search()
	fmt.Printf("path: %v \n", search)
}

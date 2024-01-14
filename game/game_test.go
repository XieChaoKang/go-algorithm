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

func TestSli(t *testing.T) {
	a := []int{1, 2, 3}
	a = a[1:]
	fmt.Printf("%v \n", a)
}

func TestCalculateMulti(t *testing.T) {
	fmt.Printf("%f \n", CalculateMulti(7))
	fmt.Printf("%f \n", CalculateMulti(7.01))
	fmt.Printf("%f \n", CalculateMulti(14))
}

func TestLinkGame(t *testing.T) {
	game := &LinkGame{
		LinkMap: [][]int{
			{4, 15, 2, 0, 14, 2, 12},
			{14, 6, 4, 0, 13, 8, 13},
			{14, 6, 2, 0, 14, 8, 2},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{4, 15, 6, 0, 9, 2, 2},
			{14, 12, 6, 0, 9, 2, 2},
			{6, 6, 14, 0, 2, 4, 2},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
	}
	//printMap(game.LinkMap)
	fmt.Printf("%v \n", game.LinkMap[2][0])
	fmt.Printf("%v \n", game.LinkMap[2][4])
	suc, path := game.FindPassablePath(2, 0, 2, 4)
	fmt.Printf("%v %v \n", suc, path)
	for _, ints := range path {
		fmt.Printf("%v \n", game.LinkMap[ints[0]][ints[1]])
	}
	//game.ShuffleMapWithPassablePath()
	//printMap(game.LinkMap)
}

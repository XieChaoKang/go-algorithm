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
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 11, 9},
			{0, 13, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 6, 7, 13},
			{0, 0, 0, 0, 7, 5, 1},
			{0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{11, 0, 0, 0, 0, 0, 9},
			{0, 5, 0, 0, 6, 0, 0},
			{0, 9, 0, 0, 9, 0, 0},
		},
	}
	printMap(game.LinkMap)
	fmt.Printf("%v \n", game.LinkMap[10][4])
	fmt.Printf("%v \n", game.LinkMap[5][4])
	suc, path := game.FindPassablePath(10, 4, 5, 4)
	fmt.Printf("%v %v \n", suc, path)
	for _, ints := range path {
		fmt.Printf("%v \n", game.LinkMap[ints[0]][ints[1]])
	}
	//println(game.CheckHasPassablePath())
	//game.ShuffleMapWithPassablePath()
	//printMap(game.LinkMap)
}

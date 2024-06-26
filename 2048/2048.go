package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const gridSize = 4

type game struct {
	grid [gridSize][gridSize]int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := newGame()
	g.addRandomTile()
	g.addRandomTile()
	for {
		g.printGrid()
		if g.checkWin() {
			fmt.Println("You win!")
			break
		}
		if g.checkGameOver() {
			fmt.Println("Game Over!")
			break
		}
		var move string
		fmt.Scanln(&move)
		switch move {
		case "z":
			g.moveUp()
		case "s":
			g.moveDown()
		case "q":
			g.moveLeft()
		case "d":
			g.moveRight()
		default:
			fmt.Println("Invalid move. Use i, j, k, l.")
			continue
		}
		g.addRandomTile()
		clearScreen()
	}
}

func newGame() *game {
	return &game{}
}

func (g *game) addRandomTile() {
	emptyCells := [][2]int{}
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if g.grid[i][j] == 0 {
				emptyCells = append(emptyCells, [2]int{i, j})
			}
		}
	}
	if len(emptyCells) == 0 {
		return
	}
	cell := emptyCells[rand.Intn(len(emptyCells))]
	if rand.Intn(10) == 0 {
		g.grid[cell[0]][cell[1]] = 4
	} else {
		g.grid[cell[0]][cell[1]] = 2
	}
}

func (g *game) printGrid() {
	clearScreen()
	for i := 0; i < gridSize; i++ {
		printBorder()
		for j := 0; j < gridSize; j++ {
			printTile(g.grid[i][j])
		}
		fmt.Println("|")
	}
	printBorder()
}

func printBorder() {
	for i := 0; i < gridSize; i++ {
		fmt.Print("+----")
	}
	fmt.Println("+")
}

func printTile(value int) {
	colorCode := getColor(value)
	if value == 0 {
		fmt.Print("|    ")
	} else {
		fmt.Printf("|%s%4d%s", colorCode, value, "\033[0m")
	}
}

func getColor(value int) string {
	switch value {
	case 2:
		return "\033[32m" // Green
	case 4:
		return "\033[33m" // Yellow
	case 8:
		return "\033[31m" // Red
	case 16:
		return "\033[35m" // Magenta
	case 32:
		return "\033[36m" // Cyan
	case 64:
		return "\033[34m" // Blue
	case 128:
		return "\033[92m" // Bright Green
	case 256:
		return "\033[93m" // Bright Yellow
	case 512:
		return "\033[91m" // Bright Red
	case 1024:
		return "\033[95m" // Bright Magenta
	case 2048:
		return "\033[96m" // Bright Cyan
	default:
		return "\033[37m" // White
	}
}

func (g *game) checkWin() bool {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if g.grid[i][j] == 2048 {
				return true
			}
		}
	}
	return false
}

func (g *game) checkGameOver() bool {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if g.grid[i][j] == 0 {
				return false
			}
			if i > 0 && g.grid[i][j] == g.grid[i-1][j] {
				return false
			}
			if i < gridSize-1 && g.grid[i][j] == g.grid[i+1][j] {
				return false
			}
			if j > 0 && g.grid[i][j] == g.grid[i][j-1] {
				return false
			}
			if j < gridSize-1 && g.grid[i][j] == g.grid[i][j+1] {
				return false
			}
		}
	}
	return true
}

func (g *game) moveUp() {
	for j := 0; j < gridSize; j++ {
		var temp [gridSize]int
		tempIndex := 0
		for i := 0; i < gridSize; i++ {
			if g.grid[i][j] != 0 {
				if tempIndex > 0 && temp[tempIndex-1] == g.grid[i][j] {
					temp[tempIndex-1] *= 2
				} else {
					temp[tempIndex] = g.grid[i][j]
					tempIndex++
				}
			}
		}
		for i := 0; i < gridSize; i++ {
			if i < tempIndex {
				g.grid[i][j] = temp[i]
			} else {
				g.grid[i][j] = 0
			}
		}
	}
}

func (g *game) moveDown() {
	for j := 0; j < gridSize; j++ {
		var temp [gridSize]int
		tempIndex := 0
		for i := gridSize - 1; i >= 0; i-- {
			if g.grid[i][j] != 0 {
				if tempIndex > 0 && temp[tempIndex-1] == g.grid[i][j] {
					temp[tempIndex-1] *= 2
				} else {
					temp[tempIndex] = g.grid[i][j]
					tempIndex++
				}
			}
		}
		for i := gridSize - 1; i >= 0; i-- {
			if gridSize-1-i < tempIndex {
				g.grid[i][j] = temp[gridSize-1-i]
			} else {
				g.grid[i][j] = 0
			}
		}
	}
}

func (g *game) moveLeft() {
	for i := 0; i < gridSize; i++ {
		var temp [gridSize]int
		tempIndex := 0
		for j := 0; j < gridSize; j++ {
			if g.grid[i][j] != 0 {
				if tempIndex > 0 && temp[tempIndex-1] == g.grid[i][j] {
					temp[tempIndex-1] *= 2
				} else {
					temp[tempIndex] = g.grid[i][j]
					tempIndex++
				}
			}
		}
		for j := 0; j < gridSize; j++ {
			if j < tempIndex {
				g.grid[i][j] = temp[j]
			} else {
				g.grid[i][j] = 0
			}
		}
	}
}

func (g *game) moveRight() {
	for i := 0; i < gridSize; i++ {
		var temp [gridSize]int
		tempIndex := 0
		for j := gridSize - 1; j >= 0; j-- {
			if g.grid[i][j] != 0 {
				if tempIndex > 0 && temp[tempIndex-1] == g.grid[i][j] {
					temp[tempIndex-1] *= 2
				} else {
					temp[tempIndex] = g.grid[i][j]
					tempIndex++
				}
			}
		}
		for j := gridSize - 1; j >= 0; j-- {
			if gridSize-1-j < tempIndex {
				g.grid[i][j] = temp[gridSize-1-j]
			} else {
				g.grid[i][j] = 0
			}
		}
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

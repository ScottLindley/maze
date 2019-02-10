package maze

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"../models"
)

type maze = models.Maze

type node = models.Node
type changeable = models.Changeable

// DrawGraph -
func DrawGraph(m *maze, graph map[string]*node) {
	out, err := os.Create("./maze_assets/graph.png")
	if err != nil {
		log.Fatal(err)
	}

	outImg := m.Image.(changeable)

	for _, node := range graph {
		x, y := node.Point.X, node.Point.Y
		isStart := (x == m.Start.X) && (y == m.Start.Y)
		isFinish := (x == m.Finish.X) && (y == m.Finish.Y)
		if isStart || isFinish {
			outImg.Set(x, y, color.RGBA{0, 255, 0, 255})
		} else {
			outImg.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}

	err = png.Encode(out, outImg)
	if err != nil {
		log.Fatal(err)
	}
}

// DrawPath -
func DrawPath(m *maze, path []image.Point) {
	out, err := os.Create("./maze_assets/solution.png")
	if err != nil {
		log.Fatal(err)
	}

	outImg := m.Image.(changeable)

	current := 0
	next := 1
	for next < len(path) {
		DrawPointConn(m, path[current], path[next], outImg)
		current = next
		next++
	}

	err = png.Encode(out, outImg)
	if err != nil {
		log.Fatal(err)
	}
}

// DrawPointConn -
func DrawPointConn(m *maze, a, b image.Point, img changeable) {
	img.Set(a.X, a.Y, color.RGBA{0, 0, 255, 255})
	img.Set(b.X, b.Y, color.RGBA{0, 0, 255, 255})
	start := 0
	end := 0
	write := "x"

	if a.X == b.X {
		if a.Y < b.Y {
			start = a.Y
			end = b.Y
		} else {
			start = b.Y
			end = a.Y
		}
		write = "y"
	} else if a.Y == b.Y {
		if a.X < b.X {
			start = a.X
			end = b.X
		} else {
			start = b.X
			end = a.X
		}
		write = "x"
	}

	for start < end {
		if write == "x" {
			img.Set(start, a.Y, color.RGBA{0, 0, 255, 255})
		} else {
			img.Set(a.X, start, color.RGBA{0, 0, 255, 255})
		}
		start++
	}
}

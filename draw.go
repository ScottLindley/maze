package main

import (
	"image/color"
	"image/png"
	"log"
	"os"
)

func drawGraph(maze *Maze, graph map[Node][]Edge) {
	out, err := os.Create("./maze_assets/graph.png")
	if err != nil {
		log.Fatal(err)
	}

	outImg := maze.image.(Changeable)

	for node := range graph {
		x, y := node.point.X, node.point.Y
		isStart := (x == maze.start.X) && (y == maze.start.Y)
		isFinish := (x == maze.finish.X) && (y == maze.finish.Y)
		if isStart || isFinish {
			outImg.Set(x, y, color.RGBA{0, 0, 255, 255})
		} else {
			outImg.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}

	err = png.Encode(out, outImg)
	if err != nil {
		log.Fatal(err)
	}
}

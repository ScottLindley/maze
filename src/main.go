package main

// https://hereandabove.com/maze/mazeorig.form.html

import (
	"image/png"
	"log"
	"os"

	"./maze"
	"./models"
)

func main() {
	reader, err := os.Open("./maze_assets/maze_md.png")
	if err != nil {
		log.Fatal(err)
	}

	image, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	m := &models.Maze{Image: image}
	maze.MarkCorners(m)
	maze.MarkDoors(m)
	graph := maze.BuildGraph(m)
	// maze.DrawGraph(m, graph)

	path := maze.Solve(graph, m)
	maze.DrawPath(m, path)
}

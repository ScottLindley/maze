package main

import (
	"image/png"
	"log"
	"os"

	"./maze"
	"./models"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("image path required")
	}
	imagePath := os.Args[1]
	reader, err := os.Open(imagePath)
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

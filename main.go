package main

import (
	"image/png"
	"log"
	"os"
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

	maze := Maze{image: image}
	markCorners(&maze)
	markDoors(&maze)
	graph := buildGraph(&maze)

	log.Print(graph)

	drawGraph(&maze, graph)
}

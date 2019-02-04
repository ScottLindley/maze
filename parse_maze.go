package main

//https://hereandabove.com/maze/mazeorig.form.html

import (
	"errors"
	"image"
	"image/color"
	"log"
)

// Maze -
type Maze struct {
	image    image.Image
	start    image.Point
	finish   image.Point
	NECorner image.Point
	NWCorner image.Point
	SECorner image.Point
	SWCorner image.Point
}

func (maze *Maze) width() int {
	size := maze.image.Bounds().Size()
	return size.X
}

func (maze *Maze) height() int {
	size := maze.image.Bounds().Size()
	return size.Y
}

func (maze *Maze) widthHeight() (int, int) {
	size := maze.image.Bounds().Size()
	return size.X, size.Y
}

// XAxis key
const XAxis = "x_axis"

// YAxis key
const YAxis = "y_axis"

// NORTH neighbor index
const NORTH = 0

// SOUTH neighbor index
const SOUTH = 1

// WEST neighbor index
const WEST = 2

// EAST neighbor index
const EAST = 3

type mazeGraph struct {
}

func markDoors(maze *Maze) {
	foundStart := false

	// walk north wall
	p := maze.NWCorner
	for p.X < maze.NECorner.X {
		p = walk(p, maze, XAxis)
		if isDoor(p, maze) {
			if !foundStart {
				foundStart = true
				maze.start = p
				p.X++
			} else {
				maze.finish = p
				return
			}
		}
	}

	// walk south wall
	p = maze.SWCorner
	for p.X < maze.SECorner.X {
		p = walk(p, maze, XAxis)
		if isDoor(p, maze) {
			if !foundStart {
				foundStart = true
				maze.start = p
				p.X++
			} else {
				maze.finish = p
				return
			}
		}
	}

	// walk west wall
	p = maze.NWCorner
	for p.Y < maze.SWCorner.Y {
		p = walk(p, maze, YAxis)
		if isDoor(p, maze) {
			if !foundStart {
				foundStart = true
				maze.start = p
				p.Y++
			} else {
				maze.finish = p
				return
			}
		}
	}

	// walk east wall
	p = maze.NECorner
	for p.Y < maze.SECorner.Y {
		p = walk(p, maze, YAxis)
		if isDoor(p, maze) {
			if !foundStart {
				foundStart = true
				maze.start = p
				p.Y++
			} else {
				maze.finish = p
				return
			}
		}
	}
}

func walk(start image.Point, maze *Maze, axis string) image.Point {
	w := maze.NECorner.X
	h := maze.SECorner.Y

	walker := 0
	bound := w
	if axis == XAxis {
		walker = start.X
		bound = w
	} else {
		walker = start.Y
		bound = h
	}

	c := maze.image.At(start.X, start.Y)
	for isWall(c) && walker < bound {
		walker++
		if axis == XAxis {
			c = maze.image.At(walker, start.Y)
		} else {
			c = maze.image.At(start.X, walker)
		}
	}

	if axis == XAxis {
		return image.Point{X: walker, Y: start.Y}
	}
	return image.Point{X: start.X, Y: walker}
}

func markCorners(maze *Maze) {
	width, height := maze.widthHeight()

	topLeft, err := findTopLeftCorner(maze)
	if err != nil {
		log.Fatal(err)
	}

	maze.NWCorner = topLeft

	for x := width - 1; x >= topLeft.X; x-- {
		if isWall(maze.image.At(x, topLeft.Y)) {
			maze.NECorner = image.Point{X: x, Y: topLeft.Y}
			break
		}
	}

	for y := height - 1; y >= topLeft.Y; y-- {
		if isWall(maze.image.At(topLeft.X, y)) {
			maze.SWCorner = image.Point{X: topLeft.X, Y: y}
			break
		}
	}

	for x := width - 1; x >= maze.SWCorner.X; x-- {
		if isWall(maze.image.At(x, maze.SWCorner.Y)) {
			maze.SECorner = image.Point{X: x, Y: maze.SWCorner.Y}
			break
		}
	}
}

func findTopLeftCorner(maze *Maze) (image.Point, error) {
	x := 0
	y := 0
	width, height := maze.widthHeight()

	for y < height {
		for x < width {
			c := maze.image.At(x, y)
			if isWall(c) {
				return image.Point{X: x, Y: y}, nil
			}
			x++
		}
		y++
		x = 0
	}

	return image.Point{X: x, Y: y}, errors.New("no walls found")
}

func isDoor(p image.Point, maze *Maze) bool {
	c := maze.image.At(p.X, p.Y)
	if isWall(c) {
		return false
	}
	neighbors := getNeighbors(p, maze)
	north := neighbors[NORTH]
	south := neighbors[SOUTH]
	west := neighbors[WEST]
	east := neighbors[EAST]

	if validPoint(north, maze) &&
		validPoint(south, maze) &&
		isWall(getColorFromPoint(north, maze)) &&
		isWall(getColorFromPoint(south, maze)) {
		return true
	}

	if validPoint(east, maze) &&
		validPoint(west, maze) &&
		isWall(getColorFromPoint(east, maze)) &&
		isWall(getColorFromPoint(west, maze)) {
		return true
	}

	return false
}

func isWall(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	avg := (r + g + b) / 3
	return avg < 128
}

func validPoint(p image.Point, maze *Maze) bool {
	w, h := maze.widthHeight()
	return p.X > 0 && p.X < w && p.Y > 0 && p.Y < h
}

func getColorFromPoint(p image.Point, maze *Maze) color.Color {
	return maze.image.At(p.X, p.Y)
}

func getNeighbors(p image.Point, maze *Maze) []image.Point {
	neighbors := make([]image.Point, 4)
	for i := 0; i < 4; i++ {
		neighbors[i] = image.Point{X: -1, Y: -1}
	}

	width, height := maze.widthHeight()
	x := p.X
	y := p.Y

	if x > 0 {
		neighbors[WEST] = image.Point{X: x - 1, Y: y}
	}
	if x < width {
		neighbors[EAST] = image.Point{X: x + 1, Y: y}
	}
	if y > 0 {
		neighbors[NORTH] = image.Point{X: x, Y: y - 1}
	}
	if y < height {
		neighbors[SOUTH] = image.Point{X: x, Y: y + 1}
	}

	return neighbors
}

func buildGraph(maze *Maze) {

}

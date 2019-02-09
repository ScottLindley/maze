package maze

import (
	"errors"
	"image"
	"image/color"
	"log"

	"../models"
)

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

func MarkDoors(maze *models.Maze) {
	foundStart := false

	// walk north wall
	p := maze.NWCorner
	for p.X < maze.NECorner.X {
		p = walk(p, maze, XAxis)
		if isDoor(p, maze) {
			if !foundStart {
				foundStart = true
				maze.Start = p
				p.X++
			} else {
				maze.Finish = p
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
				maze.Start = p
				p.X++
			} else {
				maze.Finish = p
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
				maze.Start = p
				p.Y++
			} else {
				maze.Finish = p
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
				maze.Start = p
				p.Y++
			} else {
				maze.Finish = p
				return
			}
		}
	}
}

func walk(start image.Point, maze *models.Maze, axis string) image.Point {
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

	c := maze.Image.At(start.X, start.Y)
	for isWall(c) && walker < bound {
		walker++
		if axis == XAxis {
			c = maze.Image.At(walker, start.Y)
		} else {
			c = maze.Image.At(start.X, walker)
		}
	}

	if axis == XAxis {
		return image.Point{X: walker, Y: start.Y}
	}
	return image.Point{X: start.X, Y: walker}
}

func MarkCorners(maze *models.Maze) {
	width, height := maze.WidthHeight()

	topLeft, err := findTopLeftCorner(maze)
	if err != nil {
		log.Fatal(err)
	}

	maze.NWCorner = topLeft

	for x := width - 1; x >= topLeft.X; x-- {
		if isWall(maze.Image.At(x, topLeft.Y)) {
			maze.NECorner = image.Point{X: x, Y: topLeft.Y}
			break
		}
	}

	for y := height - 1; y >= topLeft.Y; y-- {
		if isWall(maze.Image.At(topLeft.X, y)) {
			maze.SWCorner = image.Point{X: topLeft.X, Y: y}
			break
		}
	}

	for x := width - 1; x >= maze.SWCorner.X; x-- {
		if isWall(maze.Image.At(x, maze.SWCorner.Y)) {
			maze.SECorner = image.Point{X: x, Y: maze.SWCorner.Y}
			break
		}
	}
}

func findTopLeftCorner(maze *models.Maze) (image.Point, error) {
	x := 0
	y := 0
	width, height := maze.WidthHeight()

	for y < height {
		for x < width {
			c := maze.Image.At(x, y)
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

func isDoor(p image.Point, maze *models.Maze) bool {
	c := maze.Image.At(p.X, p.Y)
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

func validPoint(p image.Point, maze *models.Maze) bool {
	w, h := maze.WidthHeight()
	return p.X >= 0 && p.X < w && p.Y >= 0 && p.Y < h
}

func getColorFromPoint(p image.Point, maze *models.Maze) color.Color {
	return maze.Image.At(p.X, p.Y)
}

func getNeighbors(p image.Point, maze *models.Maze) []image.Point {
	neighbors := make([]image.Point, 4)
	for i := 0; i < 4; i++ {
		neighbors[i] = image.Point{X: -1, Y: -1}
	}

	width, height := maze.WidthHeight()
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

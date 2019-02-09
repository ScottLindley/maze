package models

import (
	"image"
	"image/color"
	"log"
)

// Maze -
type Maze struct {
	Image    image.Image
	Start    image.Point
	Finish   image.Point
	NECorner image.Point
	NWCorner image.Point
	SECorner image.Point
	SWCorner image.Point
}

// Width -
func (maze *Maze) Width() int {
	size := maze.Image.Bounds().Size()
	return size.X
}

// Height -
func (maze *Maze) Height() int {
	size := maze.Image.Bounds().Size()
	return size.Y
}

// WidthHeight -
func (maze *Maze) WidthHeight() (int, int) {
	size := maze.Image.Bounds().Size()
	return size.X, size.Y
}

// Node -
type Node struct {
	Point             image.Point
	CameFrom          *Node
	DistanceFromStart float64
	Links             []*Node
}

// Link -
func (n *Node) Link(node *Node) {
	n.Links = append(n.Links, node)
}

// PrintLinks -
func (n *Node) PrintLinks() {
	linked := []image.Point{}
	for _, l := range n.Links {
		linked = append(linked, l.Point)
	}
	log.Print(n.Point, " -> ", linked)
}

// Changeable -
type Changeable interface {
	Set(x, y int, c color.Color)
	At(x, y int) color.Color
	Bounds() image.Rectangle
	ColorModel() color.Model
}

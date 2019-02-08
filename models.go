package main

import "image"

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

// Node -
type Node struct {
	point image.Point
}

// Queue -
// type Queue struct {
// 	head *Node
// 	tail *Node
// 	size int
// }

// func (q *Queue) enqueue(n *Node) {
// 	if q.size == 0 {
// 		q.head = n
// 		q.tail = n
// 		q.size = 1
// 	} else {
// 		q.size++
// 		q.
// 	}
// }

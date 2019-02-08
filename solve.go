package main

import (
	"image"
	"log"
)

func buildGraph(maze *Maze) map[Node][]Node {
	graph := make(map[Node][]Node)
	start := Node{point: maze.start}
	graph[start] = []Node{}
	seen := make(map[string]bool)
	queue := []Node{}
	queue = append(queue, start)

	for len(queue) > 0 {
		node := queue[0]
		point := node.point
		queue = queue[1:]
		log.Print("Q len ", len(queue))
		seen[pointKey(point)] = true

		neighbors := getValidNeighbors(point, maze)
		log.Print("neigh   ", neighbors)
		for _, n := range neighbors {
			neighborNode := Node{point: n}
			graph[node] = append(graph[node], neighborNode)
			_, inGraph := graph[neighborNode]
			if shouldMakeNode(point, maze) && !inGraph {
				graph[neighborNode] = []Node{}
			}
			if !seen[pointKey(n)] {
				queue = append(queue, neighborNode)
			} else {
				log.Print("we have seen node")
			}
		}
	}
	return graph
}

func pointKey(p image.Point) string {
	return string(p.X) + "," + string(p.Y)
}

func shouldMakeNode(p image.Point, maze *Maze) bool {
	neighbors := getNeighbors(p, maze)

	options := 0
	// n := false
	// s := false
	// e := false
	// w := false

	if validPoint(neighbors[NORTH], maze) && isInsideMaze(neighbors[NORTH], maze) {
		options++
		// n = true
	}
	if validPoint(neighbors[SOUTH], maze) && isInsideMaze(neighbors[SOUTH], maze) {
		options++
		// s = true
	}
	if validPoint(neighbors[EAST], maze) && isInsideMaze(neighbors[EAST], maze) {
		options++
		// e = true
	}
	if validPoint(neighbors[WEST], maze) && isInsideMaze(neighbors[WEST], maze) {
		options++
		// w = true
	}

	if options > 2 {
		return true
	}
	// if (n && s) || (e && w) {
	// 	return true
	// }
	return false
}

func getValidNeighbors(p image.Point, maze *Maze) []image.Point {
	neighbors := getNeighbors(p, maze)
	valid := []image.Point{}
	for _, n := range neighbors {
		if validPoint(n, maze) && isInsideMaze(n, maze) && !isWall(getColorFromPoint(n, maze)) {
			valid = append(valid, n)
		}
	}
	return valid
}

func isInsideMaze(p image.Point, maze *Maze) bool {
	return (maze.NWCorner.X < p.X &&
		maze.NECorner.X > p.X &&
		maze.NWCorner.Y < p.Y &&
		maze.SWCorner.Y > p.Y)
}

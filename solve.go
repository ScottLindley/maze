package main

import (
	"image"
	"log"
	"math"
)

func distanceFromNodes(a, b Node) int {
	if a.point.X == b.point.X {
		return int(math.Abs(float64(a.point.Y) - float64(b.point.Y)))
	}
	return int(math.Abs(float64(a.point.X) - float64(b.point.X)))
}

func dfs(maze *Maze, graph map[Node][]Edge, seen map[string]bool, lastNode Node, currentPoint image.Point) {
	neighbors := getValidNeighbors(currentPoint, maze)
	for _, n := range neighbors {
		// new node -- process it
		if !seen[pointKey(n)] {
			seen[pointKey(n)] = true
			node := Node{point: n}
			// we should make a node here -- add it to the graph
			if shouldMakeNode(n, maze) {
				graph[node] = []Edge{}
				distance := distanceFromNodes(lastNode, node)
				edge := Edge{from: lastNode, to: node, length: distance}
				graph[lastNode] = append(graph[lastNode], edge)
				dfs(maze, graph, seen, node, node.point)
			} else {
				dfs(maze, graph, seen, lastNode, n)
			}
		}
	}
}

func buildGraph(maze *Maze) map[Node][]Edge {
	graph := make(map[Node][]Edge)
	start := Node{point: maze.start}
	finish := Node{point: maze.finish}
	graph[start] = []Edge{}
	graph[finish] = []Edge{}
	seen := make(map[string]bool)
	dfs(maze, graph, seen, start, maze.start)
	return graph
}

func buildGraph2(maze *Maze) map[Node][]Node {
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
	neighbors := getValidNeighbors(p, maze)
	log.Print(neighbors)
	if len(neighbors) > 2 {
		return true
	}
	if len(neighbors) == 2 {
		a, b := neighbors[0], neighbors[1]
		if (a.X == b.X) || (a.Y == b.Y) {
			return false
		}
	}
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

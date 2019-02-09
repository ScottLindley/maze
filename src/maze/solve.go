package maze

import (
	"image"
	"math"
	"strconv"

	"../models"
)

func distanceFromNodes(a, b models.Node) float64 {
	deltaX := math.Abs(float64(a.Point.X) - float64(b.Point.X))
	deltaY := math.Abs(float64(a.Point.Y) - float64(b.Point.Y))
	return math.Sqrt((deltaX * deltaX) + (deltaY * deltaY))
}

func biLinkNodes(a, b *models.Node) {
	if pointKey(a.Point) == pointKey(b.Point) {
		return
	}
	a.Link(b)
	b.Link(a)
}

func dfs(maze *models.Maze, graph map[string]*models.Node, seen map[string]bool, lastNode *models.Node, currentPoint image.Point) {
	if pointKey(currentPoint) == pointKey(maze.Finish) {
		finish := &models.Node{Point: maze.Finish}
		graph[pointKey(maze.Finish)] = finish
		biLinkNodes(lastNode, finish)
		return
	}

	neighbors := getValidNeighbors(currentPoint, maze)
	for _, n := range neighbors {
		// new node -- process it
		if !seen[pointKey(n)] {
			seen[pointKey(n)] = true
			// we should make a node here -- add it to the graph
			if shouldMakeNode(n, maze) {
				node := &models.Node{Point: n}
				biLinkNodes(lastNode, node)
				graph[pointKey(n)] = node
				graph[pointKey(lastNode.Point)] = lastNode
				dfs(maze, graph, seen, node, node.Point)
			} else {
				dfs(maze, graph, seen, lastNode, n)
			}
		}
	}
}

// BuildGraph -
func BuildGraph(maze *models.Maze) map[string]*models.Node {
	graph := make(map[string]*models.Node)
	start := &models.Node{Point: maze.Start, DistanceFromStart: 0}
	seen := make(map[string]bool)
	dfs(maze, graph, seen, start, maze.Start)
	return graph
}

func computeNodePriority(finish models.Node) func(interface{}) float64 {
	return func(n interface{}) float64 {
		node := n.(*models.Node)
		distanceFromFinish := distanceFromNodes(finish, *node)
		return math.MaxFloat64 - distanceFromFinish
	}
}

// Solve -
func Solve(graph map[string]*models.Node, maze *models.Maze) []image.Point {
	startKey := pointKey(maze.Start)
	finishKey := pointKey(maze.Finish)
	start := graph[startKey]
	finish := graph[finishKey]
	seen := make(map[string]bool)
	queue := models.PriorityQueue{GetPriority: computeNodePriority(*finish)}
	queue.Enqueue(start)

	solved := false
	for !solved && queue.Size > 0 {
		node := queue.Dequeue().(*models.Node)
		seen[pointKey(node.Point)] = true

		for _, n := range node.Links {
			// insert into max heap
			key := pointKey(n.Point)
			if !seen[key] {
				graph[key].CameFrom = node
				atFinish := key == pointKey(maze.Finish)
				if atFinish { // we're done
					solved = true
					break
				}
				queue.Enqueue(n)
			}
		}
	}

	path := []image.Point{}
	current := graph[pointKey(maze.Finish)]
	for current != nil && current != graph[pointKey(maze.Start)] {
		path = append(path, current.Point)
		current = current.CameFrom
	}
	return path
}

func pointKey(p image.Point) string {
	return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
}

func shouldMakeNode(p image.Point, maze *models.Maze) bool {
	if isStartOrFinish(p, maze) {
		return true
	}
	neighbors := getValidNeighbors(p, maze)
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

func getValidNeighbors(p image.Point, maze *models.Maze) []image.Point {
	neighbors := getNeighbors(p, maze)
	valid := []image.Point{}
	for _, n := range neighbors {
		if isStartOrFinish(n, maze) || (validPoint(n, maze) && isInsideMaze(n, maze) && !isWall(getColorFromPoint(n, maze))) {
			valid = append(valid, n)
		}
	}
	return valid
}

func isStartOrFinish(p image.Point, m *models.Maze) bool {
	isStart := pointKey(p) == pointKey(m.Start)
	isFinish := pointKey(p) == pointKey(m.Finish)
	return isStart || isFinish
}

func isInsideMaze(p image.Point, maze *models.Maze) bool {
	return (maze.NWCorner.X < p.X &&
		maze.NECorner.X > p.X &&
		maze.NWCorner.Y < p.Y &&
		maze.SWCorner.Y > p.Y)
}

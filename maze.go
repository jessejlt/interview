package interview

import (
	"fmt"
)

// MinimumStepsThroughMaze returns the minimum number of steps necessary to navigate a maze
// from some start to end coordinates.
func MinimumStepsThroughMaze(maze [][]bool, startRow, startColumn int, endRow, endColumn int) int {

	// Rule of thumb between BFS and DFS:
	//
	// DFS: Want to visit every node
	// BFS: Want shortest path
	//

	if len(maze) == 0 {
		return -1
	}

	// We'll be doing a BFS, where our maze is an adjacency matrix.
	// We would use Dijkstra’s if this were weighted.

	q := NewQueue()
	start := point{row: startRow, column: startColumn}
	end := point{row: endRow, column: endColumn}
	q.Add(start)

	visited := make(map[string]bool)
	pred := make(map[string]string)

	for !q.IsEmpty() {

		p, ok := q.Remove().(point)
		if !ok {
			return -1
		}

		if p.row == end.row && p.column == end.column {

			path := pathFromPredecessors(
				pred,
				start.String(),
				end.String(),
			)
			return len(path)
		}

		if _, ok := visited[p.String()]; ok {
			continue
		}
		visited[p.String()] = true

		for _, neighbor := range adjacentPointsBF(maze, p.row, p.column) {

			if visited[neighbor.String()] {
				continue
			}

			pred[neighbor.String()] = p.String()
			q.Add(neighbor)
		}
	}

	return -1
}

type point struct {
	row    int
	column int
}

func (p point) String() string {
	return fmt.Sprintf("(%d, %d)", p.row, p.column)
}

func pathFromPredecessors(pred map[string]string, start, end string) []string {

	var (
		path []string
		t    = end
	)

	for {

		t = pred[t]
		path = append(path, t)

		if t == start {
			break
		}
	}

	return path
}

func adjacentPointsBF(maze [][]bool, row, column int) []point {

	if row >= len(maze) || column >= len(maze[row]) {
		return nil
	}

	var points []point

	// top
	if row+1 < len(maze) && !maze[row+1][column] {

		points = append(points, point{row: row + 1, column: column})
	}

	// right
	if column+1 < len(maze[row]) && !maze[row][column+1] {

		points = append(points, point{row: row, column: column + 1})
	}

	// down
	if row-1 >= 0 && !maze[row-1][column] {

		points = append(points, point{row: row - 1, column: column})
	}

	// left
	if column-1 >= 0 && !maze[row][column-1] {

		points = append(points, point{row: row, column: column - 1})
	}

	return points
}

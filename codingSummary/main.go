package main

import (
	"fmt"
)

type Node struct {
	id          int
	color       int
	left, right *Node
}

func main() {
	// a^b
	var n int
	fmt.Scan(&n)
	nodes := make([]Node, n)
	// init
	for i := 0; i < n; i++ {
		nodes[i] = Node{
			id:    i + 1,
			color: -1,
			left:  nil,
			right: nil,
		}
	}

	// build trees
	for i := 0; i < n-1; i++ {
		var parentIndex int
		fmt.Scan(&parentIndex)
		parentIndex -= 1
		sonIndex := i + 1
		// insert
		if nodes[parentIndex].left == nil {
			nodes[parentIndex].left = &nodes[sonIndex]
		} else {
			nodes[parentIndex].right = &nodes[sonIndex]
		}
	}
	// get color
	for i := 0; i < n; i++ {
		var color int
		fmt.Scan(&color)
		nodes[i].color = color
	}

	var dfs func(node *Node) int
	dfs = func(node *Node) int {
		if node.left == nil && node.right == nil {
			return 1
		}
		if node.color == 1 {
			return dfs(node.left) + dfs(node.right)
		} else {
			return dfs(node.left) ^ dfs(node.right)
		}
	}
	fmt.Println(dfs(&nodes[0]))
}

func xor(a, b int) int {
	return a ^ b
}
func add(a, b int) int {
	return a + b
}

func convertToFloat1(s string) string {
	// find .
	sb := []byte(s)
	pointIndex := -1
	for i := len(sb) - 1; i >= 0; i-- {
		if sb[i] == '.' {
			pointIndex = i
			break
		}
	}
	if pointIndex == -1 {
		sb = append(sb, '.')
		sb = append(sb, '0')
	} else {
		if pointIndex < len(sb)-2 {
			if sb[pointIndex+2] >= '5' {
				sb[pointIndex+1] += 1
			}
		}
		sb = sb[:pointIndex+2]
	}

	return string(sb)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

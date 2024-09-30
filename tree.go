package tree

import (
	"fmt"
)

// Node 구조체는 이진 탐색 트리의 노드를 정의합니다.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func MakeNode(num int) *Node {
	return &Node{Value: num}
}

// Insert 함수는 이진 탐색 트리에 새로운 값을 삽입하는 함수입니다.
func (tree *Node) Insert(num int) {
	if num < tree.Value {
		if tree.Left == nil {
			tree.Left = MakeNode(num)
		} else {
			tree.Left.Insert(num)
		}
	} else {
		if tree.Right == nil {
			tree.Right = MakeNode(num)
		} else {
			tree.Right.Insert(num)
		}
	}
}

func InOrder(n *Node) {
	if n == nil {
		return
	}
	InOrder(n.Left)
	fmt.Print(n.Value, " ")
	InOrder(n.Right)
}

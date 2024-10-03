package bst

// 이진 탐색 트리의 노드를 정의
type NodeF struct {
	Value int
	Left  *Node
	Right *Node
}

// num을 Value로 갖는 Node를 생성하여 주소값 return
func MakeNodeF(num int) *Node {

}

// num과 tree의 Value를 비교하여 left 또는 right에 삽입
func (tree *Node) InsertF(num int) {

}

//n.Left, n.Value, n.Right 순으로 순회하며 Value 출력
func InOrderF(n *Node) {

}

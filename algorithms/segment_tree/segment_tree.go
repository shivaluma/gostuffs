package segment_tree

type SegmentTree struct {
	root *Node
}

type Node struct {
	start, end, sum int
	left, right     *Node
}

func (n *Node) Len() int {
	return n.end - n.start
}

func (n *Node) Sum() int {
	return n.sum
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) Query(start, end int) int {
	if start > end {
		return 0
	}

	if start == n.start && n.end == end {
		return n.sum
	}

	mid := (n.start + n.end) / 2

	if start > mid {
		return n.right.Query(start, end)
	}

	if end < mid {
		return n.left.Query(start, end)
	}

	return n.left.Query(start, mid) + n.right.Query(mid+1, end)

}

func NewSegmentTree(arr []int) *SegmentTree {
	return &SegmentTree{
		root: buildTree(arr, 0, len(arr)-1),
	}
}

func buildTree(arr []int, start, end int) *Node {
	if start == end {
		return &Node{
			start: start,
			end:   end,
			sum:   arr[start],
			left:  nil,
			right: nil,
		}
	}

	mid := (start + end) / 2

	leftNode := buildTree(arr, start, mid)
	rightNode := buildTree(arr, mid+1, end)

	node := &Node{
		start: start,
		end:   end,
		sum:   leftNode.sum + rightNode.sum,
		left:  leftNode,
		right: rightNode,
	}

	return node
}

func (st *SegmentTree) Query(start, end int) int {
	return st.root.Query(start, end)
}

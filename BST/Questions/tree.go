package questions

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

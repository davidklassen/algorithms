package linked_lists

type NodeRef struct {
	data *Node
	next *NodeRef
}

func (nr *NodeRef) Add(n *Node) *NodeRef {
	return &NodeRef{
		data: n,
		next: nr,
	}
}

func buildRef(l *Node) *NodeRef {
	var r *NodeRef
	for l != nil {
		r = r.Add(l)
		l = l.next
	}
	return r
}

func Intersection(l1, l2 *Node) *Node {
	r1 := buildRef(l1)
	r2 := buildRef(l2)
	var prev *Node = nil
	for r1.data == r2.data {
		prev = r1.data
		r1 = r1.next
		r2 = r2.next
	}
	return prev
}

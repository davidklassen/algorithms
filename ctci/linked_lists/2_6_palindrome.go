package linked_lists

type NodeLink struct {
	data *Node
	next *NodeLink
}

func reverseLink(head *Node) *NodeLink {
	var res *NodeLink
	for head != nil {
		if res == nil {
			res = &NodeLink{data: head}
		} else {
			res = &NodeLink{data: head, next: res}
		}
		head = head.next
	}
	return res
}

func Palindrome(head *Node) bool {
	if head == nil {
		return true
	}
	r := reverseLink(head)
	for head != r.data {
		if head.data != r.data.data {
			return false
		}
		if head.next == r.data {
			break
		}
		head = head.next
		r = r.next
	}
	return true
}

package trees_and_graphs

type TrieNode struct {
	val      rune
	terminal bool
	children map[rune]*TrieNode
}

func NewTrie() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func (t *TrieNode) Insert(s string) {
	cur := t
	for _, r := range []rune(s) {
		next, ok := cur.children[r]
		if !ok {
			next = &TrieNode{val: r, children: make(map[rune]*TrieNode)}
			cur.children[r] = next
		}
		cur = next
	}
	cur.terminal = true
}

func (t *TrieNode) Contains(s string) bool {
	cur := t
	for _, r := range []rune(s) {
		next, ok := cur.children[r]
		if !ok {
			return false
		}
		cur = next
	}
	return cur.terminal
}

func (t *TrieNode) IsEmpty() bool {
	return !t.terminal && len(t.children) == 0
}

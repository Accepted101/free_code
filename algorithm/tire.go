package algorithm

// 0-1字典树，取异或最大值
type trieNode struct {
	children [2]*trie
	cnt      int
}

type trie struct {
	root *trieNode //根
}

func (*trie) insert(x int) {
	now := root
	for x != 0 {
		now = now.children[x&1]
	}
}

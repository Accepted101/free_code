package algorithm

// 0-1字典树，取异或最大值
type trieNode struct {
	children [2]*trieNode
	cnt      [2]int
}

type trie struct {
	root   *trieNode
	maxLen int
}

func GetBit(grayCode int, bit int) int {
	return (grayCode >> bit) & 1
}

func NewTrie(maxLen int) *trie {
	return &trie{
		root:   &trieNode{},
		maxLen: maxLen,
	}
}

func (t *trie) Insert(x int) {
	now := t.root
	for i := t.maxLen; i >= 0; i-- {
		p := GetBit(x, i)
		now.cnt[p]++
		if now.children[p] == nil {
			now.children[p] = new(trieNode)
		}
		now = now.children[p]
	}
}

func (t *trie) Delete(x int) {
	now := t.root
	for i := t.maxLen; i >= 0; i-- {
		p := GetBit(x, i)
		now.cnt[p]--
		now = now.children[p]
	}
}

func (t *trie) MaxMatch(x int) int {
	now := t.root
	ans := 0
	for i := t.maxLen; i >= 0; i-- {
		p := GetBit(x, i)
		if now.cnt[p^1] > 0 {
			now = now.children[p^1]
			ans |= 1 << i
		} else {
			now = now.children[p]
		}
	}
	return ans
}

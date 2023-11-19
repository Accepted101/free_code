package algorithm

import (
	"testing"
)

func TestGetBit(t *testing.T) {
	grayCode := 5
	bit := 1
	expected := 0
	result := GetBit(grayCode, bit)
	if result != expected {
		t.Errorf("GetBit(%d, %d) = %d; want %d", grayCode, bit, result, expected)
	}
}

func TestTrie_Insert(t *testing.T) {
	trie := &trie{root: &trieNode{}, maxLen: 3}
	x := 5
	trie.Insert(x)
	expected := 1
	result := trie.root.children[1].cnt[0]
	if result != expected {
		t.Errorf("Insert(%d) = %d; want %d", x, result, expected)
	}
}

func TestTrie_Delete(t *testing.T) {
	trie := &trie{root: &trieNode{}, maxLen: 3}
	x := 5
	trie.Insert(x)
	trie.Delete(x)
	expected := 0
	result := trie.root.children[1].cnt[0]
	if result != expected {
		t.Errorf("Delete(%d) = %d; want %d", x, result, expected)
	}
}

func TestTrie_MaxMatch(t *testing.T) {
	trie := &trie{root: &trieNode{}, maxLen: 3}
	trie.Insert(5)
	trie.Insert(6)
	x := 7
	expected := 3
	result := trie.MaxMatch(x)
	if result != expected {
		t.Errorf("MaxMatch(%d) = %d; want %d", x, result, expected)
	}
}

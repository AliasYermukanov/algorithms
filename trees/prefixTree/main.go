package main

import (
	"fmt"
)

const alphabet = 26

type Ptree struct {
	val byte
	child map[byte]*Ptree
	end bool
}

func addWord(root *Ptree, word string) {
	node := root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, found := node.child[char]; !found {
			node.child[char] = &Ptree{val: char, child: make(map[byte]*Ptree)}
		}
		node = node.child[char]
	}
	node.end = true
	fmt.Println(node)
}

func searchWord(root *Ptree, word string) bool {
	node := root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, found := node.child[char]; !found {
			return false
		}
		node = node.child[char]
	}
	if !node.end {
		return false
	}
	return true
}


func main() {


	array := []string{"the", "any", "there", "answer", "a", "bye", "by", "their","hah"}

	root := &Ptree{child: make(map[byte]*Ptree)}

	for i := 0; i < len(array); i++ {
		addWord(root, array[i])
	}

	fmt.Println(searchWord(root,""))

}

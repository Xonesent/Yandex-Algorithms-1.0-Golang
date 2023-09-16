package main

import "fmt"

type node struct {
	predok      *node
	value       int
	left        *node
	right       *node
	leftlength  int
	rightlength int
}

type tree struct {
	root *node
}

func main() {
	tree := &tree{}
	var value int = 1
	for {
		fmt.Scan(&value)
		if value == 0 {
			break
		}
		tree.Insert(value)
	}
	leaves := tree.Proxod()
	for i := 0; i < len(leaves); i++ {
		leaves[i].glubina()
	}
	if tree.Avlchechek() {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func initialization(value int) *node {
	return &node{predok: nil, value: value, left: nil, right: nil, leftlength: 0, rightlength: 0}
}

func (root *tree) Insert(value int) {
	if root.root == nil {
		root.root = initialization(value)
	} else {
		root.root.insert(value)
	}
}

func (root *node) insert(value int) {
	if value > root.value {
		if root.right == nil {
			elem := initialization(value)
			elem.predok = root
			root.right = elem
		} else {
			root.right.insert(value)
		}
	} else if value < root.value {
		if root.left == nil {
			elem := initialization(value)
			elem.predok = root
			root.left = elem
		} else {
			root.left.insert(value)
		}
	}
}

func (root *tree) Proxod() []*node {
	leaves := make([]*node, 0)
	if root.root != nil {
		root.root.proxod(&leaves)
	}
	return leaves
}

func (root *node) proxod(leaves *[]*node) {
	if root != nil {
		root.left.proxod(leaves)
		if root.left == nil && root.right == nil {
			*leaves = append(*leaves, root)
		}
		root.right.proxod(leaves)
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (leaf *node) glubina() {
	if leaf.predok != nil {
		if leaf.value > leaf.predok.value {
			leaf.predok.rightlength = max(leaf.rightlength, leaf.leftlength) + 1
		} else if leaf.value < leaf.predok.value {
			leaf.predok.leftlength = max(leaf.rightlength, leaf.leftlength) + 1
		}
		leaf.predok.glubina()
	}
}

func (root *tree) Avlchechek() bool {
	var check bool = true
	if root.root != nil {
		root.root.avlchechek(&check)
	}
	return check
}

func (root *node) avlchechek(check *bool) {
	if root != nil {
		root.left.avlchechek(check)
		if !(root.leftlength-root.rightlength == 0 || root.leftlength-root.rightlength == -1 || root.leftlength-root.rightlength == 1) {
			*check = false
		}
		root.right.avlchechek(check)
	}
}

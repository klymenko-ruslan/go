package datastructures

import (
	"fmt"
)

type Tree struct {
	root *treeNode
	Size int
}
type treeNode struct {
	val int
	leftNode, rightNode *treeNode
}
func (t *Tree) Add(val int) bool {
	if t.Size == 0 {
		t.root = &treeNode{val: val}
		t.Size++
		return true
	}
	return t.add(val, t.root)
}
func(t *Tree) add(val int, n *treeNode) bool {
	if val == n.val {
		return false
	}
	if val < n.val {
		if n.leftNode != nil {
			return t.add(val, n.leftNode)
		}
		n.leftNode = &treeNode{val: val}
	} else if val > n.val {
		if n.rightNode != nil {
			return t.add(val, n.rightNode)
		}
		n.rightNode = &treeNode{val: val}
	}
	t.Size++
	return true
}
func (t *Tree) Find(val int) bool {
	return t.find(val, *t.root) != nil
}
func (t *Tree) find(val int, n treeNode) *treeNode {
	if val < n.val {
		if n.leftNode == nil {
			return nil
		}
		return t.find(val, *n.leftNode)
	} else if val > n.val {
		if n.rightNode == nil {
			return nil
		}
		return t.find(val, *n.rightNode)
	}
	return &n
}
func (t *Tree) Delete(val int) bool {
	n := t.find(val, *t.root)
	if n == nil {
		return false
	}
	return t.delete(n)
}
func (t *Tree) delete(n *treeNode) bool {
	if n.leftNode != nil && n.rightNode != nil {
		tmp := n.rightNode
		for tmp.leftNode != nil {
			tmp = tmp.leftNode
		}
		parent := t.findParent(t.root, tmp)
		if parent != nil {
			parent.leftNode = nil
		}
		tmp.leftNode = n.leftNode
		tmp.rightNode = n.rightNode
		n = tmp
	} else if n.leftNode != nil || n.rightNode != nil {
		nodeToReassignment := &treeNode{}
		if n.leftNode != nil {
			nodeToReassignment = n.leftNode
		} else {
			nodeToReassignment = n.rightNode
		}
		parent := t.findParent(t.root, n)
		if parent == nil {
			if n.rightNode != nil {
				t.root = n.rightNode
			} else {
				t.root = n.leftNode
			}
		} else if parent.leftNode != nil && parent.leftNode.val == n.val {
			parent.leftNode = nodeToReassignment
		} else {
			parent.rightNode = nodeToReassignment
		}
	} else {
		parent := t.findParent(t.root, n)
		if parent == nil {
			t.root = nil
		} else if parent.leftNode != nil && n.val == parent.leftNode.val {
			parent.leftNode = nil
		} else {
			parent.rightNode = nil
		}
	}
	t.Size--
	return true
}
func (t *Tree)findParent(currentRoot *treeNode, n *treeNode) *treeNode {
	if currentRoot == nil || n == t.root {
		return nil
	}
	if currentRoot.leftNode != nil && currentRoot.leftNode.val == n.val || currentRoot.rightNode != nil && currentRoot.rightNode.val == n.val {
		return currentRoot
	} else {
		if n.val <= currentRoot.val {
			return t.findParent(currentRoot.leftNode, n)
		} else {
			return t.findParent(currentRoot.rightNode, n)
		}
	}
}
func (t *Tree) FindDepth() int {
	return t.findDepth(t.root)
}
func (t *Tree) findDepth(n *treeNode) int {
	if n == nil {
		return 0
	}
	lDepth := t.findDepth(n.leftNode)
	rDepth := t.findDepth(n.rightNode)
	if lDepth > rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}
func (t *Tree) PrintTree() {
	t.printTree(t.root)
}
func (t *Tree) printTree(n *treeNode) {
	if n != nil {
		if n.leftNode != nil {
			t.printTree(n.leftNode)
		}
		fmt.Println(n.val)
		if n.rightNode != nil {
			t.printTree(n.rightNode)
		}
	}
}

//      __      ___        _______ _____  ______ ______
//     /\ \    / / |      |__   __|  __ \|  ____|  ____|
//    /  \ \  / /| |         | |  | |__) | |__  | |__
//   / /\ \ \/ / | |         | |  |  _  /|  __| |  __|
//  / ____ \  /  | |____     | |  | | \ \| |____| |____
// /_/    \_\/   |______|    |_|  |_|  \_\______|______|
//

package AVL_Tree

import (
	"bytes"
	"errors"
)

type AVLNode struct {
	Key       []byte   //node key
	Value     []byte   //node value
	Height    int      //
	ParNode   *AVLNode //Parent AVL node
	LeftNode  *AVLNode //Left node with key less than current node
	RightNode *AVLNode //Right node with key greater than current node
}

func NewAVLLeaf(
	parNode *AVLNode,
	key,
	value []byte) *AVLNode {

	return &AVLNode{
		Key:       key,
		Value:     value,
		Height:    0,
		ParNode:   parNode,
		LeftNode:  nil,
		RightNode: nil,
	}
}

//return the node position of either the matching node or the locatation to place a node
func (n *AVLNode) findMatchPosition(searchKey []byte) (match bool, position *AVLNode) {

	if n == nil {
		return false, n
	}

	//The result will be 0 if a==b, -1 if a < b, and +1 if a > b
	switch bytes.Compare(searchKey, n.Key) {
	case 0:
		return true, n
	case -1:
		return n.LeftNode.findMatchPosition(searchKey) //send a reference to the parent node down for returning
	case 1:
		return n.RightNode.findMatchPosition(searchKey)
	}

	return
}

//return the node position of either the matching node or the locatation to place a node
func (n *AVLNode) findAddPosition(searchKey []byte) (leftChild bool, parNode *AVLNode, err error) {

	if n == nil {
		return false, n, errors.New("Node is nil")
	}

	//The result will be 0 if a==b, -1 if a < b, and +1 if a > b
	switch bytes.Compare(searchKey, n.Key) {
	case 0:
		return false, n, errors.New("Duplicate key found")
	case -1:
		if n.LeftNode == nil {
			return true, n, nil
		}
		return n.LeftNode.findAddPosition(searchKey) //send a reference to the parent node down for returning
	case 1:
		if n.RightNode == nil {
			return false, n, nil
		}
		return n.RightNode.findAddPosition(searchKey)
	}

	return
}

func (n *AVLNode) findMin() *AVLNode {
	if n.LeftNode == nil {
		return n
	}
	return n.LeftNode.findMin()
}

func (n *AVLNode) findMax() *AVLNode {
	if n.RightNode == nil {
		return n
	}
	return n.RightNode.findMax()
}

//updates the height of the current node
func (n *AVLNode) updateHeight() {

	maxHeight := -1

	if n.LeftNode != nil {
		maxHeight = n.LeftNode.Height
	}

	if n.RightNode != nil && n.RightNode.Height > maxHeight {
		maxHeight = n.RightNode.Height
	}

	n.Height = maxHeight + 1

	return
}

//currently only used for testing purposes
func (n *AVLNode) updateHeightRecursive() {
	if n == nil {
		return
	}
	n.updateHeight()
	n.ParNode.updateHeightRecursive()

	return
}

func (n *AVLNode) getBalance() int {

	RightHeight := 0
	LeftHeight := 0

	if n.RightNode != nil {
		RightHeight = n.RightNode.Height + 1
	}

	if n.LeftNode != nil {
		LeftHeight = n.LeftNode.Height + 1
	}

	return RightHeight - LeftHeight
}

func (n *AVLNode) updateBalance(tr *AVLTree) {
	bal := n.getBalance()

	switch {
	case bal > 1:
		if n.RightNode.getBalance() > 0 { //Left Left Rotation
			n.rotate(tr, true) //rotateLeft
		} else { //Right Left Rotation
			n.RightNode.rotate(tr, false) //rotateRight
			n.rotate(tr, true)
		}
	case bal < -1:
		if n.LeftNode.getBalance() < 0 { //Right Right Rotation
			n.rotate(tr, false)
		} else { //Left Right Rotation
			n.LeftNode.rotate(tr, true)
			n.rotate(tr, false)
		}
	}

	return
}

//update the height and  balances from the area of action upwards
// this will allow the tree to be balanced in the most
// compact way
func (n *AVLNode) updateHeightBalanceRecursive(tr *AVLTree) {
	if n == nil {
		return
	}

	n.updateHeight()
	n.updateBalance(tr)
	n.ParNode.updateHeightBalanceRecursive(tr)

	return
}

//rotate function used by rotateRight/Left
func (n *AVLNode) rotate(tr *AVLTree, left bool) {

	var nodeUp *AVLNode

	//old parent takes owernership of left nodes right child as its left child
	if left {
		nodeUp = n.RightNode
		n.RightNode = nodeUp.LeftNode
		nodeUp.LeftNode = n
	} else {
		nodeUp = n.LeftNode
		n.LeftNode = nodeUp.RightNode
		nodeUp.RightNode = n
	}

	//parent swap
	nodeUp.ParNode = n.ParNode
	n.ParNode = nodeUp

	if nodeUp.ParNode != nil {
		//update the new parents (old grandparents) child too
		if nodeUp.ParNode.LeftNode == n {
			nodeUp.ParNode.LeftNode = nodeUp
		} else {
			nodeUp.ParNode.RightNode = nodeUp
		}
	} else {
		//if no parents then set the nodeUp as the new tree trunk
		tr.trunk = nodeUp
	}

	//update effected heights
	n.updateHeight()
	nodeUp.updateHeight()

	return
}

//used for testing purposes
//recursively add all the downstream
func (n *AVLNode) printStructure() (out string) {

	parKey, leftKey, rightKey := "nil", "nil", "nil"

	if n.ParNode != nil {
		parKey = string(n.ParNode.Key[:])
	}

	if n.LeftNode != nil {
		out += n.LeftNode.printStructure()
		leftKey = string(n.LeftNode.Key[:])
	}

	if n.RightNode != nil {
		out += n.RightNode.printStructure()
		rightKey = string(n.RightNode.Key[:])
	}

	out += "key: " + string(n.Key[:]) +
		" value: " + string(n.Value[:]) +
		" parent: " + parKey +
		" leftChild: " + leftKey +
		" rightChild: " + rightKey +
		"\n"

	return
}

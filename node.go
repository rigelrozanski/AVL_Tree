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
}

//currently only used for testing purposes
func (n *AVLNode) updateHeightRecursive() {
	if n == nil {
		return
	}
	n.updateHeight()
	n.ParNode.updateHeightRecursive()
}

func (n *AVLNode) getBalance() int {

	RightHeight := 0
	LeftHeight := 0

	if n.RightNode != nil {
		RightHeight = n.RightNode.Height
	}

	if n.LeftNode != nil {
		LeftHeight = n.LeftNode.Height
	}

	return RightHeight - LeftHeight
}

func (n *AVLNode) updateBalance() {
	bal := n.getBalance()

	switch {
	case bal > 1:
		if n.RightNode.getBalance() > 0 { //Left Left Rotation
			n.rotateLeft()
		} else { //Right Left Rotation
			n.RightNode.rotateRight()
			n.rotateLeft()
		}
	case bal < -1:
		if n.LeftNode.getBalance() < 0 { //Right Right Rotation
			n.rotateRight()
		} else { //Left Right Rotation
			n.LeftNode.rotateLeft()
			n.rotateRight()
		}

	}
}

//update the height and  balances from the area of action upwards
// this will allow the tree to be balanced in the most
// compact way
func (n *AVLNode) updateHeightBalanceRecursive() {
	if n == nil {
		return
	}
	n.updateHeight()
	n.updateBalance()
	n.ParNode.updateHeightBalanceRecursive()
}

func (n *AVLNode) rotateLeft() {

	//Original right node moving up during rotation
	nodeUp := n.RightNode

	//new parent takes on old parent's parent
	nodeUp.ParNode = n.ParNode

	//old parent takes owernership of right nodes left child as its right child
	n.RightNode = nodeUp.LeftNode

	//old parent node becomes lower left child of old right node
	nodeUp.LeftNode = n
	n.ParNode = nodeUp
}

func (n *AVLNode) rotateRight() {

	//Original left node moving up during rotation
	nodeUp := n.LeftNode

	//new parent takes on old parent's parent
	nodeUp.ParNode = n.ParNode

	//old parent takes owernership of left nodes right child as its left child
	n.LeftNode = nodeUp.RightNode

	//old parent node becomes lower right child of old left node
	nodeUp.RightNode = n
	n.ParNode = nodeUp
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

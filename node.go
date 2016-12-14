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
)

type AVLNode struct {
	Key       []byte   //node key
	Value     []byte   //node value
	Height    int      //
	ParNode   *AVLNode //Parent AVL node
	LeftNode  *AVLNode //Left node with key less than current node
	RightNode *AVLNode //Right node with key greater than current node
}

func NewAVLTrunk(
	key,
	value []byte) *AVLNode {

	return NewAVLLeaf(nil, key, value)
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
func (n *AVLNode) findPosition(searchKey []byte) (match bool, position *AVLNode) {

	if n == nil {
		return false, n //parent is set to nil if the trunk node
	}

	//The result will be 0 if a==b, -1 if a < b, and +1 if a > b
	switch bytes.Compare(searchKey, n.Key) {
	case 0:
		return true, n
	case -1:
		return n.LeftNode.findPosition(searchKey) //send a reference to the parent node down for returning
	case 1:
		return n.RightNode.findPosition(searchKey)
	}

	return
}

//return the node position of either the matching node or the locatation to place a node
func (n *AVLNode) findPositionAndParent(searchKey []byte, parNodeIn *AVLNode) (match bool, position, parNode *AVLNode) {

	if n == nil {
		return false, n, parNodeIn //parent is set to nil if the trunk node
	}

	//The result will be 0 if a==b, -1 if a < b, and +1 if a > b
	switch bytes.Compare(searchKey, n.Key) {
	case 0:
		return true, n, parNodeIn
	case -1:
		return n.LeftNode.findPositionAndParent(searchKey, n) //send a reference to the parent node down for returning
	case 1:
		return n.RightNode.findPositionAndParent(searchKey, n)
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

func (n *AVLNode) balance() int {

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
	bal := n.balance()

	switch {
	case bal > 1:
		if n.RightNode.balance() > 0 { //Left Left Rotation
			n.rotateLeft()
		} else { //Right Left Rotation
			n.RightNode.rotateRight()
			n.rotateLeft()
		}
	case bal < -1:
		if n.LeftNode.balance() < 0 { //Right Right Rotation
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
	if n != nil {
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

//      __      ___        _______ _____  ______ ______
//     /\ \    / / |      |__   __|  __ \|  ____|  ____|
//    /  \ \  / /| |         | |  | |__) | |__  | |__
//   / /\ \ \/ / | |         | |  |  _  /|  __| |  __|
//  / ____ \  /  | |____     | |  | | \ \| |____| |____
// /_/    \_\/   |______|    |_|  |_|  \_\______|______|
//

package AvlTree

import (
	"bytes"

	"github.com/tendermint/go-db"
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

	return AVLNode{
		Key:       key,
		Value:     value,
		Height:    0,
		ParNode:   parNode,
		LeftNode:  nil,
		RightNode: nil,
	}
}

//return the node position of either the matching node or the locatation to place a node
func (n *AVLNode) findPosition(searchKey []byte) (match bool, position, parent *AVLNode) {

	if n == nil {
		return false, n, n.ParNode //parent is set to nil if the trunk node
	}

	//The result will be 0 if a==b, -1 if a < b, and +1 if a > b
	switch bytes.Compare(searchKey, *n.key) {
	case 0:
		return true, n, n.ParNode
	case -1:
		return n.LeftNode.findPosition(searchKey)
	case 1:
		return n.RightNode.findPosition(searchKey)
	}
}

//find either the minimum or maximum sub node key value
func (n *AVLNode) findExtremum(min bool) *AVLNode {
	if min {
		//calling recursively within these function variables
		// should be slighly more efficient then calling the
		// whole function because it avoids checking min
		iterateLeft := func(n2 *AVLNode) *AVLNode {
			if n2.LeftNode == nil {
				return n2
			}
			return iterateLeft(n2.LeftNode)
		}

		return iterateLeft(n)

	} else {
		iterateRight := func(n2 *AVLNode) *AVLNode {
			if n2.RightNode == nil {
				return n2
			}
			return iterateRight(n2.RightNode)
		}

		return iterateRight(n)
	}
}

//updates the height of the current node
func (n *AVLNode) updateHeight() {

	maxHeight := -1

	if n.LeftNode != nil {
		maxHeight = *n.LeftNode.Height
	}

	if n.RightNode != nil && *n.RightNode.Height > maxHeight {
		maxHeight = *n.RightNode.Height
	}

	*n.Height = maxHeight + 1
}

func (n *AVLNode) balance() int {

	RightHeight := 0
	LeftHeight := 0

	if n.RightNode != nil {
		RightHeight = *n.RightNode.Height
	}

	if n.LeftNode != nil {
		LeftHeight = *n.LeftNode.Height
	}

	return RightHeight - LeftHeight
}

func (n *AVLNode) updateBalance() {
	bal = n.balance()

	switch {
	case bal > 1:
		if *n.RightNode.balance() > 0 { //Left Left Rotation
			n.rotateLeft()
		} else { //Right Left Rotation
			n.RightNode.rotateRight()
			n.rotateLeft()
		}
	case bal < -1:
		if *n.LeftNode.balance() < 0 { //Right Right Rotation
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
func (n *AVLNode) updateBalanceRecursive() {
	if n != nil {
		return
	}
	n.updateHeight()
	n.updateBalance()
	n.ParNode.updateBalanceRecursive()
}

func (n *AVLNode) rotateLeft() {

	//new parent takes on old parent's parent
	n.RightNode.ParNode = n.ParNode

	//old parent takes owernership of right nodes left child as its right child
	n.RightNode = n.RightNode.LeftNode

	//old parent node becomes lower left child of old right node
	n.RightNode.LeftNode = n
	n.ParNode = n.RightNode
}

func (n *AVLNode) rotateRight() {
	n.LeftNode.ParNode = n.ParNode
	n.ParNode = n.LeftNode
	n.LeftNode = nil
}

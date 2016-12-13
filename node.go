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
	Key     []byte   //node key
	Value   []byte   //node value
	Height  int      //
	ParNode *AVLNode //Parent AVL node
	LTNode  *AVLNode //"Left Node" node with key less than current node
	GTNode  *AVLNode //"Right Node" node with key greater than current node
}

func NewAVLTrunk(
	key,
	value []byte) *AVLNode {

	return NewAVLLeaf(nil, key, value)
}

func NewAVLLeaf(
	parent *AVLNode,
	key,
	value []byte) *AVLNode {

	return AVLNode{
		Key:     key,
		Value:   value,
		Height:  0,
		ParNode: parent,
		LTnode:  nil,
		GTnode:  nil,
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
		return n.LTnode.findPosition(searchKey)
	case 1:
		return n.GTnode.findPosition(searchKey)
	}
}

//find either the minimum or maximum sub node key value
func (n *AVLNode) findExtremum(min bool) *AVLNode {
	if min {
		//calling recursively within these function variables
		// should be slighly more efficient then calling the
		// whole function because it avoids checking min
		iterateLT := func(n2 *AVLNode) *AVLNode {
			if n2.LTNode == nil {
				return n2
			}
			return iterateLT(n2.LTNode)
		}

		return iterateLT(n)

	} else {
		iterateGT := func(n2 *AVLNode) *AVLNode {
			if n2.GTNode == nil {
				return n2
			}
			return iterateGT(n2.GTNode)
		}

		return iterateGT(n)
	}
}

//updates the height of the current node
func (n *AVLNode) updateHeight() {

	maxHeight := -1

	if n.LTnode != nil {
		maxHeight = *n.LTnode.Height
	}

	if n.GTNode != nil && *n.GTNode.Height > maxHeight {
		maxHeight = *n.GTnode.Height
	}

	*n.Height = maxHeight + 1
}

//updates the height of the current node and all parent nodes
func (n *AVLNode) updateHeightRecursive() {
	if n != nil {
		return
	}
	n.updateHeight()
	n.ParNode.updateHeightRecursive()
}

func (n *AVLNode) balance() int {
	return (GTnode.Height - LTnode.Height)
}

func (n *AVLNode) updateBalance() {
	switch n.balance() {
	case -1:

	case 1:
	}

}

//update the balance from the area of action upwards
// this will allow the tree to be balanced in the most
// compact way
func (n *AVLNode) updateBalanceRecursive() {
	if n != nil {
		return
	}
	n.updateBalance()
	n.ParNode.updateBalanceRecursive()
}

//Perform either a left or right rotation
func (n *AVLNode) rotate(left bool) {
	if left {
		n.GTNode.ParNode = n.ParNode
		n.ParNode = n.GTNode
		n.GTNode = nil
	} else {
		n.LTNode.ParNode = n.ParNode
		n.ParNode = n.LTNode
		n.LTNode = nil
	}
}

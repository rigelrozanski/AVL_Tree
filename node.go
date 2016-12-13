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
	Key    []byte   //node key
	Value  []byte   //node value
	Height int      //
	LTnode *AVLnode //"Left Node" node with key less than current node
	GTnode *AVLNode //"Right Node" node with key greater than current node
}

func NewAVLLeaf(
	key,
	value []byte) AVLNode {

	return AVLNode{
		Key:    key,
		Value:  value,
		Height: 0,
		LTnode: nil,
		GTnode: nil,
	}
}

//return the node position of either the matching node or the locatation to place a node
func (n *AVLNode) findPosition(searchKey []byte) (match bool, nodePtr *AVLNode) {

	if n == nil {
		return false, n
	}

	//The result will be 0 if a==b, -1 if a < b, and +1 if a > b
	switch bytes.Compare(searchKey, *n.key) {
	case 0:
		return true, n
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

func (n *AVLNode) balance() int {
	return (GTnode.Height - LTnode.Height)
}

//Perform either a left or right rotation
func (n *AVLNode) Rotate(left bool) {

}

//Perform either a left-right, or right-left rotation
func (n *AVLNode) RotateDouble(leftRight bool) {

}

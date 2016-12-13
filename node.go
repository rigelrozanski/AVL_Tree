//      __      ___        _______ _____  ______ ______
//     /\ \    / / |      |__   __|  __ \|  ____|  ____|
//    /  \ \  / /| |         | |  | |__) | |__  | |__
//   / /\ \ \/ / | |         | |  |  _  /|  __| |  __|
//  / ____ \  /  | |____     | |  | | \ \| |____| |____
// /_/    \_\/   |______|    |_|  |_|  \_\______|______|
//

package AvlTree

import (
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

func balance() int {
	return (GTnode.Height - LTnode.Height)
}

func compareKey(src, dest []byte) int {
	return bytes.Compare(src, dest)
}

func Rotate(left bool) {

}

func RotateDouble(leftRight bool) {

}

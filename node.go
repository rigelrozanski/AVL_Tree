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

	return AVLNode{}
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

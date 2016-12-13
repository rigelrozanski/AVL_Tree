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

type AVLTree struct {
	trunk *AVLNode
}

func NewAVLTree() AVLTree {

	return AVLTree{
		trunk: nil,
	}
}

func (tree *AVLTree) Get(key []byte) (value []byte, err error) {

	value = []byte("")
	err = nil

	return
}

func (tree *AVLTree) Add(key []byte, value []byte) error {

	//perform balancing operations

	return nil
}

func (tree *AVLTree) Update(key []byte, value []byte) error {
	return nil
}

func (tree *AVLTree) Remove(key []byte) error {
	return nil
}

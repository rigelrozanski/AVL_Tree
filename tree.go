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
	"errors"

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

func (t *AVLTree) Get(key []byte) (value []byte, err error) {

	match, matchNode := t.trunk.findPosition(key)

	if !match {
		err = errors.New("key not found")
	} else {
		value = matchNode.Value
	}

	return
}

func (t *AVLTree) Update(key []byte, value []byte) (err error) {

	match, matchNode := t.trunk.findPosition(key)

	if !match {
		err = errors.New("key not found")
	} else {
		*matchNode.Value = value
	}

	return

}

func (t *AVLTree) Add(key []byte, value []byte) (err error) {

	match, matchNode := t.trunk.findPosition(key)

	if match {
		err = errors.New("duplicate key found")
	} else {
		matchNode = &NewAVLLeaf(key, value)
	}

	//TODO perform balancing operations, update heights

	return
}

func (t *AVLTree) Remove(key []byte) (err error) {

	match, matchNode := t.trunk.findPosition(key)

	if !match {
		err = errors.New("key not found")
		return
	} else {
		//check if leaf node
		if &matchNode.Height == 0 {
			matchNode = nil
		} else {

			//if there is only one branch off of node to delete
			//  then delete replace the node to delete with the one branch

		}
	}

	//TODO perform balancing operations, update heights

	return
}

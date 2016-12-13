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
		//if leaf node being deleted, just delete it
		if &matchNode.Height == 0 {
			matchNode = nil
		} else {

			//if there is only one branch off of node to delete
			//  then replace node with one branch node
			if &matchNode.GTnode == nil {
				matchNode = matchNode.LTnode
			} else if &matchNode.LTnode == nil {
				matchNode = matchNode.GTnode
			}

			//if there are two branches off of node to delete
			//  determine the longest sub branch and on that branch
			//  replace the node; with the greatest (rightmost) key down-branch
			//  if the longest branch is the smallest (leftmost) branch,
			//  OR with the smallest (leftmost) key found downbranch
			//  if the longest branch is the greatest (rightmost) branch.
			//  if the branches are balanced, use the greatest (rightmost) branch
			//Methodology inspired by: http://www.mathcs.emory.edu/~cheung/Courses/323/Syllabus/Trees/AVL-delete.html

			//first determine the direction to replace from
			replaceFromRight := (matchNode.balance() >= 0)
			replaceFromNode := matchNode.findExtremum(replaceFromRight) // if replacing from the the right, then looking from the minimum

			//temporarily save the replacement key and value, delete its original position
			replaceKey := *replaceFromNode.Key
			replaceValue := *replaceFromNode.Value
			t.Remove(replaceKey)

			//now replace the key and value for the target node to delete
			// the branches of this node to stay the same
			*matchNode.Key = replaceKey
			*matchNode.Value = replaceValue
		}
	}

	//TODO perform balancing operations, update heights

	return
}

//      __      ___        _______ _____  ______ ______
//     /\ \    / / |      |__   __|  __ \|  ____|  ____|
//    /  \ \  / /| |         | |  | |__) | |__  | |__
//   / /\ \ \/ / | |         | |  |  _  /|  __| |  __|
//  / ____ \  /  | |____     | |  | | \ \| |____| |____
// /_/    \_\/   |______|    |_|  |_|  \_\______|______|
//

package AvlTree

import (
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

	match, matchNode, _ := t.trunk.findPosition(key)

	if !match {
		err = errors.New("key not found")
	} else {
		value = matchNode.Value
	}

	return
}

func (t *AVLTree) Update(key []byte, value []byte) (err error) {

	match, matchNode, _ := t.trunk.findPosition(key)

	if !match {
		err = errors.New("key not found")
	} else {
		*matchNode.Value = value
	}

	return

}

func (t *AVLTree) Add(key []byte, value []byte) (err error) {

	match, matchNode, parNode := t.trunk.findPosition(key)

	if match {
		err = errors.New("duplicate key found")
	} else {
		matchNode = NewAVLLeaf(parNode, key, value)
	}

	//Update heights
	matchNode.updateHeightRecursive()
	matchNode.updateBalanceRecursive()

	return
}

func (t *AVLTree) Remove(key []byte) (err error) {

	match, matchNode := t.trunk.findPosition(key)

	if !match {
		err = errors.New("key not found")
		return
	} else {
		//If leaf node being deleted, just delete it
		if &matchNode.Height == 0 {
			matchNode = nil
		} else {

			//If there is only one branch off of node to delete
			//  then replace node with one branch node
			if &matchNode.GTnode == nil {
				matchNode = matchNode.LTnode
			} else if &matchNode.LTnode == nil {
				matchNode = matchNode.GTnode
			}

			//If there are two branches off of node to delete
			//  determine the longest sub branch and on that branch
			//  replace the node; with the greatest (rightmost) key down-branch
			//  if the longest branch is the smallest (leftmost) branch,
			//  OR with the smallest (leftmost) key found downbranch
			//  if the longest branch is the greatest (rightmost) branch.
			//  if the branches are balanced, use the greatest (rightmost) branch
			//Methodology inspired by: http://www.mathcs.emory.edu/~cheung/Courses/323/Syllabus/Trees/AVL-delete.html

			//First determine the direction to replace from
			replaceFromRight := (matchNode.balance() >= 0)
			replaceFromNode := matchNode.findExtremum(replaceFromRight) // if replacing from the the right, then looking from the minimum

			//Temporarily save the replacement key and value, delete its original position
			replaceFromKey := *replaceFromNode.Key
			replaceFromValue := *replaceFromNode.Value
			t.Remove(replaceKey) //TODO verify that this wont create pointer problems based the fact that a rebalance will occur

			//Now replace the key and value for the target node to delete
			// the branches of this node to stay the same
			*matchNode.Key = replaceFromKey
			*matchNode.Value = replaceFromValue
		}
	}

	//Update heights and balance
	t.trunk.updateBalanceAndHeight()

	return
}

//      __      ___        _______ _____  ______ ______
//     /\ \    / / |      |__   __|  __ \|  ____|  ____|
//    /  \ \  / /| |         | |  | |__) | |__  | |__
//   / /\ \ \/ / | |         | |  |  _  /|  __| |  __|
//  / ____ \  /  | |____     | |  | | \ \| |____| |____
// /_/    \_\/   |______|    |_|  |_|  \_\______|______|
//

package AVL_Tree

import (
	"errors"
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
		matchNode.Value = value
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

	//Update height and balance
	matchNode.updateHeightBalanceRecursive()

	return
}

func (t *AVLTree) Remove(key []byte) (err error) {

	match, matchNode, _ := t.trunk.findPosition(key)

	if !match {
		err = errors.New("key not found")
		return
	} else {
		//If leaf node being deleted, just delete it
		if matchNode.Height == 0 {
			matchNode = nil
		} else {

			//If there is only one branch off of node to delete
			//  then replace node with one branch node
			if &matchNode.RightNode == nil {
				matchNode = matchNode.LeftNode
			} else if &matchNode.LeftNode == nil {
				matchNode = matchNode.RightNode
			}

			//If there are two branches off of node to delete
			//  determine the longest sub branch and on that branch
			//  replace the node; with the greatest (rightmost) key down-branch
			//  if the longest branch is the smallest (leftmost) branch,
			//  OR with the smallest (leftmost) key found downbranch
			//  if the longest branch is the greatest (rightmost) branch.
			//  if the branches are balanced, use the greatest (rightmost) branch
			//Methodology inspired by: http://www.mathcs.emory.edu/~cheung/Courses/323/Syllabus/Trees/AVL-delete.html

			//Determine the direction to replace, and node to switch from
			var replaceFromNode *AVLNode
			if matchNode.balance() >= 0 {
				replaceFromNode = matchNode.findMin()
			} else {
				replaceFromNode = matchNode.findMax()
			}

			//Temporarily save the replacement key and value, delete its original position
			replaceFromKey := replaceFromNode.Key
			replaceFromValue := replaceFromNode.Value
			t.Remove(replaceFromKey) //TODO verify that this wont create pointer problems based the fact that a rebalance will occur

			//Now replace the key and value for the target node to delete
			// the branches of this node to stay the same
			matchNode.Key = replaceFromKey
			matchNode.Value = replaceFromValue
		}
	}

	//Update height and balance
	matchNode.updateHeightBalanceRecursive()

	return
}

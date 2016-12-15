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
	trunk *node
}

func NewAVLTree() AVLTree {

	return AVLTree{
		trunk: nil,
	}
}

func (t *AVLTree) Get(key []byte) (value []byte, err error) {

	match, matchNode := t.trunk.findMatchPosition(key)

	if !match {
		err = errors.New("key not found")
	} else {
		value = matchNode.value
	}

	return
}

func (t *AVLTree) Update(key []byte, value []byte) error {

	match, matchNode := t.trunk.findMatchPosition(key)

	if !match {
		return errors.New("key not found")
	}

	matchNode.value = value

	return nil

}

func (t *AVLTree) Add(key []byte, value []byte) error {

	if t.trunk == nil {
		t.trunk = newNodeLeaf(nil, key, value)
		return nil
	}

	leftChild, parNode, err := t.trunk.findAddPosition(key)

	if err != nil {
		return err
	}

	if leftChild {
		parNode.leftNode = newNodeLeaf(parNode, key, value)
	} else {
		parNode.rightNode = newNodeLeaf(parNode, key, value)
	}

	//Update height and balance
	parNode.updateHeightBalanceRecursive(t)

	return nil
}

func (t *AVLTree) Remove(key []byte) error {

	match, matchNode := t.trunk.findMatchPosition(key)
	if !match {
		return errors.New("key not found")
	}

	//Update height and balance before leaving
	defer matchNode.updateHeightBalanceRecursive(t)

	setParentsChild := func(setTo *node) {
		if matchNode.parNode.leftNode == matchNode {
			matchNode.parNode.leftNode = setTo
		} else {
			matchNode.parNode.rightNode = setTo
		}

		matchNode = setTo
	}

	//If leaf node being deleted, just delete it
	if matchNode.height == 0 {
		setParentsChild(nil)
		return nil
	}

	//If there is only one branch off of node to delete
	//  then replace node with one branch node
	if matchNode.rightNode == nil {
		setParentsChild(matchNode.leftNode)
		return nil
	} else if matchNode.leftNode == nil {
		setParentsChild(matchNode.rightNode)
		return nil
	}

	//If there are two branches off of node to delete
	//  determine the longest sub branch and on that branch
	//  replace the node; with the greatest (rightmost) key down-branch
	//  if the longest branch is the smallest (leftmost) branch,
	//  OR with the smallest (leftmost) key found downbranch
	//  if the longest branch is the greatest (rightmost) branch.
	//  if the branches are balanced, use the greatest (rightmost) branch
	//Methodology inspired by: http://www.mathcs.emory.edu/~cheung/Courses/323/Syllabus/Trees/AVL-delete.html
	if matchNode.rightNode != nil && matchNode.leftNode != nil {
		//Determine the direction to replace, and node to switch from
		var replaceFromNode *node
		if matchNode.getBalance() >= 0 {
			replaceFromNode = matchNode.findMin()
		} else {
			replaceFromNode = matchNode.findMax()
		}

		//Temporarily save the replacement key and value, delete its original position
		replaceFromKey := replaceFromNode.key
		replaceFromValue := replaceFromNode.value
		t.Remove(replaceFromKey) //TODO verify that this wont create pointer problems based the fact that a rebalance will occur

		//Now replace the key and value for the target node to delete
		// the branches of this node to stay the same
		matchNode.key = replaceFromKey
		matchNode.value = replaceFromValue
		return nil
	}

	return nil
}

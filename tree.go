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

	match, matchNode := t.trunk.findMatchPosition(key)

	if !match {
		err = errors.New("key not found")
	} else {
		value = matchNode.Value
	}

	return
}

func (t *AVLTree) Update(key []byte, value []byte) error {

	match, matchNode := t.trunk.findMatchPosition(key)

	if !match {
		return errors.New("key not found")
	}

	matchNode.Value = value

	return nil

}

func (t *AVLTree) Add(key []byte, value []byte) error {

	if t.trunk == nil {
		t.trunk = NewAVLLeaf(nil, key, value)
		return nil
	}

	leftChild, parNode, err := t.trunk.findAddPosition(key)

	if err != nil {
		return err
	}

	if leftChild {
		parNode.LeftNode = NewAVLLeaf(parNode, key, value)
	} else {
		parNode.RightNode = NewAVLLeaf(parNode, key, value)
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

	//Update height and balance
	defer matchNode.updateHeightBalanceRecursive(t)

	//If leaf node being deleted, just delete it
	if matchNode.Height == 0 {
		matchNode = nil
		return nil
	}

	//If there is only one branch off of node to delete
	//  then replace node with one branch node
	if &matchNode.RightNode == nil {
		matchNode = matchNode.LeftNode
		return nil
	} else if &matchNode.LeftNode == nil {
		matchNode = matchNode.RightNode
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

	//Determine the direction to replace, and node to switch from
	var replaceFromNode *AVLNode
	if matchNode.getBalance() >= 0 {
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

	return nil
}

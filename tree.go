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

//errors used throughout
var errEmptyTree error = errors.New("Empty tree")
var errBadKey error = errors.New("Key not found")
var errNodeNil error = errors.New("Node is nil")
var errDupVal error = errors.New("Duplicate key found")

//Returns the merkle root hash (hash of the trunk node)
func (t *AVLTree) GetHash() (hash []byte, err error) {
	if t.trunk == nil {
		err = errEmptyTree
		return
	}

	hash = t.trunk.hash

	return
}

//Get a value from the tree from an existing key
func (t *AVLTree) Get(key []byte) (value []byte, err error) {
	if t.trunk == nil {
		err = errEmptyTree
		return
	}

	matchNode := t.trunk.findNode(key)

	if matchNode == nil {
		err = errBadKey
	} else {
		value = matchNode.value
	}

	return
}

//Adds the value if it doesn't exist, if it exists in updates the value
//The error should never realistically need to be used, but here for consistency
func (t *AVLTree) Set(key, value []byte) error {
	err := t.Add(key, value)
	if err == errDupVal {
		//should never produce an error is Update working properly
		return t.Update(key, value)
	} else {
		return err
	}
}

//Update a value from the tree for an already existing key
func (t *AVLTree) Update(key, value []byte) error {
	if t.trunk == nil {
		return errEmptyTree
	}

	matchNode := t.trunk.findNode(key)

	if matchNode == nil {
		return errBadKey
	}

	matchNode.value = value

	return nil

}

//Add a new key-value to the tree for a non-existent key
func (t *AVLTree) Add(key, value []byte) error {

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

//Remove a key-value pair from the tree
func (t *AVLTree) Remove(key []byte) error {
	if t.trunk == nil {
		return errEmptyTree
	}

	matchNode := t.trunk.findNode(key)
	if matchNode == nil {
		return errBadKey
	}

	//Update height and balance before leaving
	defer matchNode.updateHeightBalanceRecursive(t)

	//Replace the matchNode position held under the parents node
	// to the input setTo
	setParentsChild := func(setTo *node) {
		if matchNode.parNode.leftNode == matchNode {
			matchNode.parNode.leftNode = setTo
		} else {
			matchNode.parNode.rightNode = setTo
		}
	}

	//If leaf node being deleted, just delete it
	if matchNode.height == 0 {
		setParentsChild(nil)
		matchNode = nil
		return nil
	}

	//If there is only one branch off of node to delete
	//  then replace node with one branch node
	if matchNode.rightNode == nil {
		setParentsChild(matchNode.leftNode)
		matchNode.leftNode.parNode = matchNode.parNode
		return nil
	} else if matchNode.leftNode == nil {
		setParentsChild(matchNode.rightNode)
		matchNode.rightNode.parNode = matchNode.parNode
		return nil
	}

	//If there are two branches off of node to delete
	//  determine the longest sub branch and on that branch
	//  replace the node; with the right-most key down-branch
	//  if the longest branch is the left-most branch,
	//  OR with the left-most key found downbranch
	//  if the longest branch is the right-most branch.
	//  If the branches are balanced, use the right-most branch.
	//  Methodology inspired by: http://www.mathcs.emory.edu/~cheung/Courses/323/Syllabus/Trees/AVL-delete.html
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
		t.Remove(replaceFromKey)

		//Now replace the key and value for the target node to delete
		// the branches of this node to stay the same
		matchNode.key = replaceFromKey
		matchNode.value = replaceFromValue
		return nil
	}

	return nil
}

//Returns the tree structure
func (t *AVLTree) TreeStructure() string {
	return t.trunk.outputStructure()
}

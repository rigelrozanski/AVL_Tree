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
		trunk: newNodePlaceholder(nil), //trunk node does not contain a parent
	}
}

//errors used throughout
var errEmptyTree error = errors.New("Empty tree")
var errBadKey error = errors.New("Key not found")
var errDupVal error = errors.New("Duplicate key found")

//Returns the merkle root hash (hash of the trunk node)
func (t *AVLTree) GetHash() (hash []byte, err error) {
	if t.trunk.isPlaceholder() {
		err = errEmptyTree
		return
	}

	hash = t.trunk.hash

	return
}

//Get a value from the tree from an existing key
func (t *AVLTree) Get(key []byte) (value []byte, err error) {
	if t.trunk.isPlaceholder() {
		err = errEmptyTree
		return
	}

	matchNode := t.trunk.findNode(key)

	if matchNode.isPlaceholder() {
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
	if t.trunk.isPlaceholder() {
		return errEmptyTree
	}

	matchNode := t.trunk.findNode(key)

	if matchNode.isPlaceholder() {
		return errBadKey
	}

	matchNode.value = value

	return nil

}

//Add a new key-value to the tree for a non-existent key
func (t *AVLTree) Add(key, value []byte) error {

	if t.trunk.isPlaceholder() {
		t.trunk = newNodeLeaf(nil, key, value)
		return nil
	}

	//Placeholder location for the insert
	insertPH := t.trunk.findNode(key)

	//If a non-placeholder value is found then the record already exists
	if !insertPH.isPlaceholder() {
		return errDupVal
	}

	//Parent node adopting a new child
	parNode := insertPH.parNode

	//Give birth
	if insertPH.isLeftChild() {
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

	if t.trunk.isPlaceholder() {
		return errEmptyTree
	}

	matchNode := t.trunk.findNode(key)
	if matchNode.isPlaceholder() {
		return errBadKey
	}

	matchNode.remove()

	//Update height and balance
	matchNode.updateHeightBalanceRecursive(t)

	return nil
}

//Returns the tree structure
func (t *AVLTree) TreeStructure() string {
	return t.trunk.outputStructure()
}

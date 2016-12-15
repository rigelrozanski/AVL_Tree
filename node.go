package AVL_Tree

import (
	"bytes"
	"errors"

	"golang.org/x/crypto/sha3"
)

type node struct {
	key       []byte //node key
	value     []byte //node value
	height    int    //
	parNode   *node  //Parent AVL node
	leftNode  *node  //Left node with key less than current node
	rightNode *node  //Right node with key greater than current node
}

func newNodeLeaf(
	parNode *node,
	key,
	value []byte) *node {

	return &node{
		key:       key,
		value:     value,
		height:    0,
		parNode:   parNode,
		leftNode:  nil,
		rightNode: nil,
	}
}

func getHash(dataInput string) []byte {
	//performing the hash
	hashBytes := sha3.Sum256([]byte(dataInput))
	return hashBytes[:]
}

//return the node position of either the matching node or the locatation to place a node
func (n *node) findMatchPosition(searchkey []byte) (match bool, position *node) {

	if n == nil {
		return false, n
	}

	//The result will be 0 if a==b, -1 if a < b, and +1 if a > b
	switch bytes.Compare(searchkey, n.key) {
	case 0:
		return true, n
	case -1:
		return n.leftNode.findMatchPosition(searchkey) //send a reference to the parent node down for returning
	case 1:
		return n.rightNode.findMatchPosition(searchkey)
	}

	return
}

//return the node position of either the matching node or the locatation to place a node
func (n *node) findAddPosition(searchKey []byte) (leftChild bool, parNode *node, err error) {

	if n == nil {
		return false, n, errors.New("Node is nil")
	}

	//The result will be 0 if a==b, -1 if a < b, and +1 if a > b
	switch bytes.Compare(searchKey, n.key) {
	case 0:
		return false, n, errors.New("Duplicate key found")
	case -1:
		if n.leftNode == nil {
			return true, n, nil
		}
		return n.leftNode.findAddPosition(searchKey) //send a reference to the parent node down for returning
	case 1:
		if n.rightNode == nil {
			return false, n, nil
		}
		return n.rightNode.findAddPosition(searchKey)
	}

	return
}

func (n *node) findMin() *node {
	if n.leftNode == nil {
		return n
	}
	return n.leftNode.findMin()
}

func (n *node) findMax() *node {
	if n.rightNode == nil {
		return n
	}
	return n.rightNode.findMax()
}

//updates the height of the current node
func (n *node) updateHeight() {

	maxHeight := -1

	if n.leftNode != nil {
		maxHeight = n.leftNode.height
	}

	if n.rightNode != nil && n.rightNode.height > maxHeight {
		maxHeight = n.rightNode.height
	}

	n.height = maxHeight + 1

	return
}
func (n *node) getBalance() int {

	rightHeight := 0
	leftHeight := 0

	if n.rightNode != nil {
		rightHeight = n.rightNode.height + 1
	}

	if n.leftNode != nil {
		leftHeight = n.leftNode.height + 1
	}

	return rightHeight - leftHeight
}

func (n *node) updateBalance(tr *AVLTree) {
	bal := n.getBalance()

	switch {
	case bal > 1:
		if n.rightNode.getBalance() > 0 { //Left Left Rotation
			n.rotate(tr, true) //rotateLeft
		} else { //Right Left Rotation
			n.rightNode.rotate(tr, false) //rotateRight
			n.rotate(tr, true)
		}
	case bal < -1:
		if n.leftNode.getBalance() < 0 { //Right Right Rotation
			n.rotate(tr, false)
		} else { //Left Right Rotation
			n.leftNode.rotate(tr, true)
			n.rotate(tr, false)
		}
	}

	return
}

//update the height and  balances from the area of action upwards
// this will allow the tree to be balanced in the most
// compact way
func (n *node) updateHeightBalanceRecursive(tr *AVLTree) {
	if n == nil {
		return
	}

	n.updateHeight()
	n.updateBalance(tr)
	n.parNode.updateHeightBalanceRecursive(tr)

	return
}

//rotate function used by rotateRight/Left
func (n *node) rotate(tr *AVLTree, left bool) {

	var nodeUp *node

	//old parent takes owernership of left nodes right child as its left child
	if left {
		nodeUp = n.rightNode
		n.rightNode = nodeUp.leftNode
		nodeUp.leftNode = n
	} else {
		nodeUp = n.leftNode
		n.leftNode = nodeUp.rightNode
		nodeUp.rightNode = n
	}

	//parent swap
	nodeUp.parNode = n.parNode
	n.parNode = nodeUp

	if nodeUp.parNode != nil {
		//update the new parents (old grandparents) child too
		if nodeUp.parNode.leftNode == n {
			nodeUp.parNode.leftNode = nodeUp
		} else {
			nodeUp.parNode.rightNode = nodeUp
		}
	} else {
		//if no parents then set the nodeUp as the new tree trunk
		tr.trunk = nodeUp
	}

	//update effected heights
	n.updateHeight()
	nodeUp.updateHeight()

	return
}

/////////////////////////////
// Testing Functions
/////////////////////////////

//currently only used for testing purposes
func (n *node) updateHeightRecursive() {
	if n == nil {
		return
	}
	n.updateHeight()
	n.parNode.updateHeightRecursive()

	return
}

//used for testing purposes
//recursively add all the downstream
func (n *node) printStructure() (out string) {

	parkey, leftkey, rightkey := "nil", "nil", "nil"

	if n.parNode != nil {
		parkey = string(n.parNode.key[:])
	}

	if n.leftNode != nil {
		out += n.leftNode.printStructure()
		leftkey = string(n.leftNode.key[:])
	}

	if n.rightNode != nil {
		out += n.rightNode.printStructure()
		rightkey = string(n.rightNode.key[:])
	}

	out += "key: " + string(n.key[:]) +
		" value: " + string(n.value[:]) +
		" parent: " + parkey +
		" leftChild: " + leftkey +
		" rightChild: " + rightkey +
		"\n"

	return
}

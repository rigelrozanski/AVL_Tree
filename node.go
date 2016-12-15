package AVL_Tree

import (
	"bytes"
	"errors"

	"golang.org/x/crypto/sha3"
)

type node struct {
	key       []byte
	value     []byte
	height    int
	hash      []byte
	parNode   *node //Parent AVL node
	leftNode  *node //Left node with key less than current node
	rightNode *node //Right node with key greater than current node
}

//Generate a new leaf node with parent and hash
// note that the parents child node does not get
// assigned within this function and must be assigned
// wherever this function is called
func newNodeLeaf(
	parNode *node,
	key,
	value []byte) *node {

	out := &node{
		key:       key,
		value:     value,
		height:    0,
		hash:      nil,
		parNode:   parNode,
		leftNode:  nil,
		rightNode: nil,
	}

	out.updateHash()

	return out
}

/////////////////////////////
// Search Functions
/////////////////////////////

//Return the node with the matching key,
// if no matching node is found return nil
func (n *node) findNode(searchkey []byte) *node {

	if n == nil {
		return nil
	}

	//Compare(a,b) will be 0 if a==b, -1 if a < b, and +1 if a > b
	switch bytes.Compare(searchkey, n.key) {
	case 0:
		return n
	case -1:
		return n.leftNode.findNode(searchkey) //send a reference to the parent node down for returning
	case 1:
		return n.rightNode.findNode(searchkey)
	}

	return nil
}

//Return the placement location for a new node to add.
// The node position is represented by the parent node (parNode)
// as well as the position on that parent node which the child
// should be placed on represented by the boolean leftChild variable
// (aka. if leftChild is true place as left child position of parent,
// if false place on the right child position).
func (n *node) findAddPosition(placeKey []byte) (leftChild bool, parNode *node, err error) {

	if n == nil {
		return false, n, errors.New("Node is nil")
	}

	//Compare(a,b) will be 0 if a==b, -1 if a < b, and +1 if a > b
	switch bytes.Compare(placeKey, n.key) {
	case 0:
		return false, n, errors.New("Duplicate key found")
	case -1:
		if n.leftNode == nil {
			return true, n, nil
		}

		//Send a reference to the parent node down for returning
		return n.leftNode.findAddPosition(placeKey)
	case 1:
		if n.rightNode == nil {
			return false, n, nil
		}
		return n.rightNode.findAddPosition(placeKey)
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

/////////////////////////////
// Update Functions
/////////////////////////////

func (n *node) updateHeightAndHash() {
	n.updateHeight()
	n.updateHash()
}

//Update the hash value stored in a node.
// For branch nodes hash the concatenated hash values of branches.
// For leaf nodes hash the concatenated record key and value.
func (n *node) updateHash() {

	var hashInput []byte = nil

	switch {
	case n.leftNode != nil && n.rightNode != nil:
		hashInput = append(n.leftNode.hash, n.rightNode.hash...)
	case n.leftNode == nil && n.rightNode != nil:
		hashInput = n.rightNode.hash
	case n.leftNode != nil && n.rightNode == nil:
		hashInput = n.leftNode.hash
	case n.leftNode == nil && n.rightNode == nil:
		hashInput = append(n.key, n.value...)
	}

	hashBytes := sha3.Sum256(hashInput)
	n.hash = hashBytes[:]
}

//Update the height of the current node.
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

//Retrieve the balance for current node.
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

//Update balance for current node.
// The tree (tr) must be passed in in order to update the trunk node value if it changes.
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

//Update the height and  balances from the area of action upwards.
// This will allow the tree to be balanced in the most compact way.
func (n *node) updateHeightBalanceRecursive(tr *AVLTree) {
	if n == nil {
		return
	}

	n.updateHeightAndHash()
	n.updateBalance(tr)
	n.parNode.updateHeightBalanceRecursive(tr)

	return
}

//Rotate function, if left is true, this will be a left rotation
// otherwise function will perform a right rotation.
func (n *node) rotate(tr *AVLTree, left bool) {

	var nodeUp *node

	//Old parent takes owernership of left nodes right child as its left child
	if left {
		nodeUp = n.rightNode
		n.rightNode = nodeUp.leftNode
		nodeUp.leftNode = n
	} else {
		nodeUp = n.leftNode
		n.leftNode = nodeUp.rightNode
		nodeUp.rightNode = n
	}

	//Parent swap
	nodeUp.parNode = n.parNode
	n.parNode = nodeUp

	if nodeUp.parNode != nil {
		//Update the new parents (old grandparents) child too
		if nodeUp.parNode.leftNode == n {
			nodeUp.parNode.leftNode = nodeUp
		} else {
			nodeUp.parNode.rightNode = nodeUp
		}
	} else {
		//If no parents then set the nodeUp as the new tree trunk
		tr.trunk = nodeUp
	}

	//Update effected heights
	n.updateHeightAndHash()
	nodeUp.updateHeightAndHash()

	return
}

/////////////////////////////
// Used for Testing Purposes
/////////////////////////////

//Recursively print the structure downstream of a node.
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

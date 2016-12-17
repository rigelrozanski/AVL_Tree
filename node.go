package AVL_Tree

import (
	"bytes"

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

	//add placeholder nodes
	out.leftNode = newNodePlaceholder(out)
	out.rightNode = newNodePlaceholder(out)

	out.updateHash()

	return out
}

//Placeholder nodes are used instead of nil nodes for the children where 'real' nodes are not yet placed.
// They are mainly used to hold orientation for passing a node position when adding a new node.
// A leaf node will contain placeholder nodes for both of its children.
// placeholder nodes should never hold any children
func newNodePlaceholder(
	parNode *node) *node {

	out := &node{
		key:       nil,
		value:     nil,
		height:    0,
		hash:      nil,
		parNode:   parNode,
		leftNode:  nil,
		rightNode: nil,
	}

	return out
}

/////////////////////////////
// Attribute Functions
/////////////////////////////

func (n *node) isPlaceholder() bool {
	if n.key == nil {
		return true
	}
	return false
}

func (n *node) isTrunk() bool {
	if n.parNode == nil {
		return true
	}
	return false
}

//is the current node the left child of its parent?
func (n *node) isLeftChild() bool {
	if !n.isTrunk() {
		if n.parNode.leftNode == n {
			return true
		}
	}
	return false
}

//Retrieve the balance for current node.
func (n *node) getBalance() int {

	rightHeight := 0
	leftHeight := 0

	if !n.rightNode.isPlaceholder() {
		rightHeight = n.rightNode.height + 1
	}

	if !n.leftNode.isPlaceholder() {
		leftHeight = n.leftNode.height + 1
	}

	return rightHeight - leftHeight
}

//Recursively print the structure downstream of a node.
func (n *node) outputStructure() (out string) {

	parkey, leftkey, rightkey := "nil", "nil", "nil"

	if !n.isTrunk() {
		parkey = string(n.parNode.key[:])
	}

	if !n.leftNode.isPlaceholder() {
		out += n.leftNode.outputStructure()
		leftkey = string(n.leftNode.key[:])
	}

	if !n.rightNode.isPlaceholder() {
		out += n.rightNode.outputStructure()
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

/////////////////////////////
// Search Functions
/////////////////////////////

//Return the node with the matching key or appropriate placement location if no matching key is found,
// if no matching node is found return nil
func (n *node) findNode(searchkey []byte) *node {

	if n.isPlaceholder() {
		return n
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

func (n *node) findMin() *node {
	if n.leftNode.isPlaceholder() {
		return n
	}
	return n.leftNode.findMin()
}

func (n *node) findMax() *node {
	if n.rightNode.isPlaceholder() {
		return n
	}
	return n.rightNode.findMax()
}

/////////////////////////////
// Write Functions
/////////////////////////////

func (n *node) updateHeightAndHash() {
	n.updateHeight()
	n.updateHash()
}

//Update the hash value stored in a node.
// For branch nodes hash the concatenated hash values of branches.
// For leaf nodes hash the concatenated record key and value.
func (n *node) updateHash() {

	if n.isPlaceholder() {
		return
	}

	var hashInput []byte = nil

	//are either left or right children placeholders?
	leftIsPH := n.leftNode.isPlaceholder()
	rightIsPH := n.rightNode.isPlaceholder()

	switch {
	case !leftIsPH && !rightIsPH:
		hashInput = append(n.leftNode.hash, n.rightNode.hash...)
	case leftIsPH && !rightIsPH:
		hashInput = n.rightNode.hash
	case !leftIsPH && rightIsPH:
		hashInput = n.leftNode.hash
	case leftIsPH && rightIsPH:
		hashInput = append(n.key, n.value...)
	}

	hashBytes := sha3.Sum256(hashInput)
	n.hash = hashBytes[:]
}

//Update the height of the current node.
func (n *node) updateHeight() {

	maxHeight := -1

	if !n.leftNode.isPlaceholder() {
		maxHeight = n.leftNode.height
	}

	if !n.rightNode.isPlaceholder() && n.rightNode.height > maxHeight {
		maxHeight = n.rightNode.height
	}

	n.height = maxHeight + 1

	return
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

	if n.isPlaceholder() {
		return
	}

	n.updateHeightAndHash()
	n.updateBalance(tr)

	if !n.isTrunk() {
		n.parNode.updateHeightBalanceRecursive(tr)
	}

	return
}

//Rotate function, if leftRotation is true, this will be a left rotation
// otherwise function will perform a right rotation.
func (n *node) rotate(tr *AVLTree, leftRotation bool) {

	//original orientation of the node moving down (n)
	nWasLeftChild := n.isLeftChild()

	//node moving up
	var nodeUp *node

	//Old parent takes owernership of left nodes right child as its left child
	if leftRotation {
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

	//update the tree's trunk if a new trunk is set
	if nodeUp.isTrunk() {

		tr.trunk = nodeUp
	} else {
		//Update the new parents (old grandparents) child too
		if nWasLeftChild {
			nodeUp.parNode.leftNode = nodeUp
		} else {
			nodeUp.parNode.rightNode = nodeUp
		}
	}

	//Update effected heights
	n.updateHeightAndHash()
	nodeUp.updateHeightAndHash()

	return
}

//Removes the node by restructuring its surrounding nodes.
// Note this function does not actually perform a rebalance
//  as it iteratively calls itself when a rebalance does not
//  need to occure. Rebalance should be perfomed after an
//  external remove call.
func (n *node) remove() {

	//Never remove a placeholder,
	// this should be verified before calling this function
	if n.isPlaceholder() {
		return
	}

	//Replace the matchNode position held under the parents node
	// to the input setTo
	setParentsChild := func(setTo *node) {
		if !n.isTrunk() {

			if n.isLeftChild() {
				n.parNode.leftNode = setTo
			} else {
				n.parNode.rightNode = setTo
			}
		}
	}

	//Children placeholder booleans
	leftIsPH := n.leftNode.isPlaceholder()
	rightIsPH := n.rightNode.isPlaceholder()

	switch {

	//If leaf node being deleted, just remove reference to it TODO verify if the memory can be released here?
	case leftIsPH && rightIsPH:
		setParentsChild(newNodePlaceholder(n.parNode))
		//matchNode = nil
		return

	//If there is only one branch off of node to delete
	//  then replace node with one branch node
	case leftIsPH && !rightIsPH:
		setParentsChild(n.rightNode)
		n.rightNode.parNode = n.parNode
		return
	case !leftIsPH && rightIsPH:
		setParentsChild(n.leftNode)
		n.leftNode.parNode = n.parNode
		return

	//If there are two branches off of node to delete
	//  determine the longest sub branch and on that branch
	//  replace the node; with the right-most key down-branch
	//  if the longest branch is the left-most branch,
	//  OR with the left-most key found downbranch
	//  if the longest branch is the right-most branch.
	//  If the branches are balanced, use the right-most branch.
	//  Methodology inspired by: http://www.mathcs.emory.edu/~cheung/Courses/323/Syllabus/Trees/AVL-delete.html
	case !leftIsPH && !rightIsPH:

		//Determine the direction to replace, and node to switch from
		var replaceFromNode *node
		if n.getBalance() >= 0 {
			replaceFromNode = n.findMin()
		} else {
			replaceFromNode = n.findMax()
		}

		//Temporarily save the replacement key and value, delete its original position
		replaceFromKey := replaceFromNode.key
		replaceFromValue := replaceFromNode.value
		replaceFromNode.remove()

		//Now replace the key and value for the target node to delete
		// the branches of this node to stay the same
		n.key = replaceFromKey
		n.value = replaceFromValue
		return
	}

	return
}

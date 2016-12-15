//      __      ___        _______ _____  ______ ______
//     /\ \    / / |      |__   __|  __ \|  ____|  ____|
//    /  \ \  / /| |         | |  | |__) | |__  | |__
//   / /\ \ \/ / | |         | |  |  _  /|  __| |  __|
//  / ____ \  /  | |____     | |  | | \ \| |____| |____
// /_/    \_\/   |______|    |_|  |_|  \_\______|______|
//

package AVL_Tree

import (
	//	"bytes"
	"testing"
)

func TestAVLNode(t *testing.T) {

	//dummy tree for feeding into operations (normally used for replacing for swapping the trunk)
	tr := NewAVLTree()

	//Manually create a basic AVL tree
	var a, b, c *AVLNode

	setNewLeafs := func() {
		a = NewAVLLeaf(nil, []byte("a"), []byte("vA"))
		b = NewAVLLeaf(nil, []byte("b"), []byte("vB"))
		c = NewAVLLeaf(nil, []byte("c"), []byte("vC"))
	}

	//string the three nodes together to a basic unbalanced tree

	//////////////////////////////////////////
	// Test Configurations
	//////////////////////////////////////////
	// Config 1  Config 2  Config 3  Config 4
	//  a          a            c       c
	//   \          \          /       /
	//    b          c        b       a
	//     \        /        /         \
	//      c      b        a           b

	setConfig1 := func() {
		setNewLeafs()
		a.RightNode = b
		b.ParNode = a
		b.RightNode = c
		c.ParNode = b
	}
	setConfig2 := func() {
		setNewLeafs()
		a.RightNode = c
		c.ParNode = a
		c.LeftNode = b
		b.ParNode = c
	}
	setConfig3 := func() {
		setNewLeafs()
		c.LeftNode = b
		b.ParNode = c
		b.LeftNode = a
		a.ParNode = b
	}
	setConfig4 := func() {
		setNewLeafs()
		c.LeftNode = a
		a.ParNode = c
		a.RightNode = b
		b.ParNode = a
	}

	//Start by testing configuration 1
	setConfig1()

	//print the tree structure
	t.Log(a.printStructure())

	heightTest := func(node *AVLNode, expectedHeight int) {
		height := node.Height
		if height != expectedHeight {
			t.Errorf("bad height for %v, expected %v found %v ",
				string(node.Key[:]), expectedHeight, height)
		}
	}

	balanceTest := func(node *AVLNode, expectedBal int) {
		bal := node.getBalance()
		if bal != expectedBal {
			t.Errorf("bad balance for %v, expected %v found %v ",
				string(node.Key[:]), expectedBal, bal)
		}
	}

	//test a non-recursive height update
	b.updateHeight()
	heightTest(b, 1)

	//test a recursive height update (from leaf to trunk)
	c.updateHeightRecursive()
	heightTest(c, 0)
	heightTest(b, 1)
	heightTest(a, 2)

	//test the balances
	balanceTest(c, 0)
	balanceTest(b, 1)
	balanceTest(a, 2)

	//test rebalance to the expected position:
	//    b
	//   / \
	//  a   c

	testHeightsLog := func(aHeight, bHeight, cHeight int, logNode *AVLNode) {
		heightTest(a, aHeight)
		heightTest(b, bHeight)
		heightTest(c, cHeight)
		t.Log(logNode.printStructure())
	}

	//test manual rotation
	a.rotate(&tr, true)
	testHeightsLog(0, 1, 0, b)

	//test rotation with updateBalance
	setConfig1()
	c.updateHeightRecursive()
	a.updateBalance(&tr)
	testHeightsLog(0, 1, 0, b)

	//test rotation with updateHeightBalanceRecursive
	setConfig1()
	c.updateHeightBalanceRecursive(&tr)
	testHeightsLog(0, 1, 0, b)

	//test some alternate configurations to see if they perform the adequate rotations
	//  note the operations is always taken from the leaf node (which is where they would be called)
	setConfig2()
	b.updateHeightBalanceRecursive(&tr)
	testHeightsLog(0, 1, 0, b)

	setConfig3()
	a.updateHeightBalanceRecursive(&tr)
	testHeightsLog(0, 1, 0, b)

	setConfig4()
	b.updateHeightBalanceRecursive(&tr)
	testHeightsLog(0, 1, 0, b)

}

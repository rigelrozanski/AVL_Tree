package AVL_Tree

import (
	"testing"
)

func Testnode(t *testing.T) {

	//dummy tree for feeding into operations (normally used for replacing for swapping the trunk)
	tr := NewAVLTree()

	//Manually create a basic AVL tree
	var a, b, c *node

	setNewLeafs := func() {
		a = newNodeLeaf(nil, []byte("a"), []byte("vA"))
		b = newNodeLeaf(nil, []byte("b"), []byte("vB"))
		c = newNodeLeaf(nil, []byte("c"), []byte("vC"))
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
		a.rightNode = b
		b.parNode = a
		b.rightNode = c
		c.parNode = b
	}
	setConfig2 := func() {
		setNewLeafs()
		a.rightNode = c
		c.parNode = a
		c.leftNode = b
		b.parNode = c
	}
	setConfig3 := func() {
		setNewLeafs()
		c.leftNode = b
		b.parNode = c
		b.leftNode = a
		a.parNode = b
	}
	setConfig4 := func() {
		setNewLeafs()
		c.leftNode = a
		a.parNode = c
		a.rightNode = b
		b.parNode = a
	}

	//Start by testing configuration 1
	setConfig1()

	//print the tree structure
	t.Log(a.printStructure())

	heightTest := func(node *node, expectedHeight int) {
		height := node.height
		if height != expectedHeight {
			t.Errorf("bad height for %v, expected %v found %v ",
				string(node.key[:]), expectedHeight, height)
		}
	}

	balanceTest := func(node *node, expectedBal int) {
		bal := node.getBalance()
		if bal != expectedBal {
			t.Errorf("bad balance for %v, expected %v found %v ",
				string(node.key[:]), expectedBal, bal)
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

	testHeightsLog := func(aHeight, bHeight, cHeight int, logNode *node) {
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

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

	//String the three nodes together to a basic unbalanced tree

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

	heightBalanceTest := func(node *node, expectedHeight int, expectedBal int) {

		//test height
		height := node.height
		if height != expectedHeight {
			t.Errorf("bad height for %v, expected %v found %v ",
				string(node.key[:]), expectedHeight, height)
			t.Log(a.printStructure())
		}

		//test balance
		bal := node.getBalance()
		if bal != expectedBal {
			t.Errorf("bad balance for %v, expected %v found %v ",
				string(node.key[:]), expectedBal, bal)
			t.Log(a.printStructure())
		}
	}

	//test a non-recursive height update
	b.updateHeight()
	heightBalanceTest(b, 1, 0) //note the balance has not been updated so should still 0

	//test a non-recursive balance update
	b.updateBalance(&tr)
	heightBalanceTest(b, 1, 1) //note the balance has not been updated so should still 0

	//test a recursive height update (from leaf to trunk)
	c.updateHeightBalanceRecursive(&tr)
	heightBalanceTest(c, 0, 0)
	heightBalanceTest(b, 1, 1)
	heightBalanceTest(a, 2, 2)

	//test rebalance to the expected position:
	//    b
	//   / \
	//  a   c

	testStructure := func() {
		heightBalanceTest(a, 0, 0)
		heightBalanceTest(b, 1, 0)
		heightBalanceTest(c, 0, 0)
	}

	//test manual rotation
	a.rotate(&tr, true)
	testStructure()

	//test rotation with updateBalance
	setConfig1()
	c.updateHeightBalanceRecursive(&tr)
	testStructure()

	//test rotation with updateHeightBalanceRecursive
	setConfig1()
	c.updateHeightBalanceRecursive(&tr)
	testStructure()

	//test some alternate configurations to see if they perform the adequate rotations
	//  note the operations is always taken from the leaf node (which is where they would be called)
	setConfig2()
	b.updateHeightBalanceRecursive(&tr)
	testStructure()

	setConfig3()
	a.updateHeightBalanceRecursive(&tr)
	testStructure()

	setConfig4()
	b.updateHeightBalanceRecursive(&tr)
	testStructure()

}

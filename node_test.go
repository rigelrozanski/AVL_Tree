package AVL_Tree

import (
	"testing"
)

func TestNode(t *testing.T) {

	//Dummy tree for feeding into operations (normally used for replacing for swapping the trunk)
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
			t.Log(a.outputStructure())
		}

		//test balance
		bal := node.getBalance()
		if bal != expectedBal {
			t.Errorf("bad balance for %v, expected %v found %v ",
				string(node.key[:]), expectedBal, bal)
			t.Log(a.outputStructure())
		}
	}

	//Test a non-recursive height update
	a.updateHeight()
	heightBalanceTest(a, 1, 1) //height still 1 because b hasn't been updated

	b.updateHeight()
	heightBalanceTest(b, 1, 1)
	a.updateHeight()
	heightBalanceTest(a, 2, 2) //height/balance now 2 because b has been updated

	//Test rebalance to the expected position:
	//    b
	//   / \
	//  a   c

	testStructure := func() {
		heightBalanceTest(a, 0, 0)
		heightBalanceTest(b, 1, 0)
		heightBalanceTest(c, 0, 0)
	}

	//Test manual rotation
	a.rotate(&tr, true)
	testStructure()

	//Test rotation with updateBalance
	setConfig1()
	c.updateHeight()
	b.updateHeight()
	a.updateHeight()
	a.updateBalance(&tr)
	testStructure()

	//Test rotation with updateHeightBalanceRecursive
	setConfig1()
	c.updateHeightBalanceRecursive(&tr)
	testStructure()

	//Test some alternate configurations to see if they perform the adequate rotations.
	//Note the operations is always taken from the leaf node (which is where they would be called)
	//  under the tree.go operations.
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

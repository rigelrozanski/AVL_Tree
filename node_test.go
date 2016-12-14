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

	//Manually create a basic AVL tree
	a := NewAVLLeaf(nil, []byte("a"), []byte("vA"))
	b := NewAVLLeaf(nil, []byte("b"), []byte("vB"))
	c := NewAVLLeaf(nil, []byte("c"), []byte("vC"))

	//string the three nodes together to a basic unbalanced tree

	//////////////////////////////////////////
	// Test Configurations
	//////////////////////////////////////////
	// Config 1  Config 2  Config 3  Config 4
	//  a          a            c       c
	//   \          \          /       /
	//    b          c        b       b
	//     \        /        /         \
	//      c      b        a           a
	setConfig1 := func() {
		a.RightNode = b
		b.ParNode = a
		b.RightNode = c
		c.ParNode = b
	}
	setConfig2 := func() {
		a.RightNode = c
		c.ParNode = a
		c.LeftNode = b
		b.ParNode = c
	}
	setConfig3 := func() {
		c.LeftNode = b
		b.ParNode = c
		b.LeftNode = a
		a.ParNode = b
	}
	setConfig4 := func() {
		c.LeftNode = b
		b.ParNode = c
		b.RightNode = a
		a.ParNode = b
	}

	//Start by testing configuration 1
	setConfig1()

	//print the tree structure
	t.Log(a.printStructure())

	heightTest := func(node *AVLNode, expectedHeight int) {
		height := node.Height
		if height != expectedHeight {
			t.Errorf("bad height, expected %v found %v ", expectedHeight, height)
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

	//test rebalance to the expected position:
	//    b
	//   / \
	//  a   c

	//test manual rotation
	a.rotateLeft()
	heightTest(b, 1)
	heightTest(a, 0)
	heightTest(c, 0)
	t.Log(b.printStructure())

	//test auto rotation
	setConfig1()
	c.updateHeightBalanceRecursive()
	heightTest(b, 1)
	heightTest(a, 0)
	heightTest(c, 0)
	t.Log(b.printStructure())

	setConfig2()
	setConfig3()
	setConfig4()

}

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
	// a
	//  \
	//   b
	//    \
	//     c
	a.RightNode = b
	b.ParNode = a
	b.RightNode = c
	c.ParNode = b

	//test updating the heights on this basic tree
	c.updateHeightRecursive()

	heightTest := func(node *AVLNode, expectedHeight int) {
		height := node.Height
		if height != expectedHeight {
			t.Errorf("bad height, expected %v found %v ", expectedHeight, height)
		}
	}

	heightTest(c, 0)
	heightTest(b, 1)
	heightTest(a, 2)
}

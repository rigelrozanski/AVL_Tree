package AVL_Tree

import (
	"bytes"
	"testing"
)

func TestAVLTree(t *testing.T) {

	//The AVLTree to be tested with
	tr := NewAVLTree()

	/////////////////////////////////////
	// core sub-test function variables
	/////////////////////////////////////

	printErr := func(err error) {
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	heightBalanceNodeTest := func(n *node, expdHeight, expdBalance int) {

		//test balance
		bal := n.getBalance()
		if bal != expdBalance {
			t.Errorf("bad balance for %v, expected %v found %v ",
				string(n.key[:]), expdBal, bal)
			t.Log(n.printStructure())
		}

		//test height
		height := n.height
		if height != expdHeight {
			t.Errorf("bad height for %v, expected %v found %v ",
				string(n.key[:]), expdHeight, height)
			t.Log(n.printStructure())
		}
	}

	heightBalanceNodeKeyTest := func(nodeKey string, expdHeight, expdBalance int) {
		heightBalanceNodeTest(tr.getNode(nodeKey), expdHeight, expdBalance)
	}

	heightBalanceTrunkTest := func(expdTrunkKey string, expdHeight, expdBalance int) {

		//test trunk node key
		expdTrunkKeyByte = []byte(expdTrunkKey)
		key := tr.trunk.key
		if bytes.Compare(key, expdTrunkKeyByte) != 0 {
			t.Errorf("bad trunk key expected %v found %v ",
				expdTrunkKeyByte, string(key[:]))
			t.Log(tr.trunk.printStructure())
		}

		heightBalanceNodeTest(tr.trunk, expdHeight, expdBalance)
	}

	retrieveTest := func(key, expectedVal string, expectedExists bool) {
		recievedVal, err := tr.Get([]byte(key))
		if expectedExists {
			printErr(err)
			if bytes.Compare(recievedVal, []byte(expectedVal)) != 0 {
				t.Errorf("bad expected %v recieved %v ", expectedVal, string(recievedVal[:]))
				t.Log(tr.trunk.printStructure())
			}
		} else {
			if err == nil {
				t.Errorf("expected to receive an error when attempting to retrieve non-existent value for key %v", key)
				t.Log(tr.trunk.printStructure())
			}
		}
	}

	/////////////////////////////////////
	// Tests begin
	/////////////////////////////////////

	//Test adding several values to the AVL Tree
	//  at the same time test how the heights of the tr react
	//Expected tr structures:
	// a   a   b      b       b          d          d
	//    /   / \    / \     / \        / \        /  \
	//   b   a   c  a   c   a   d      b   e      b    f
	//                   \     / \    / \   \    /\    /\
	//                    d   c   e  a   c   f  a  c  e  g

	printErr(tr.Add([]byte("a"), []byte("vA")))
	heightBalanceTrunkTest("a", 0, 0)

	printErr(tr.Add([]byte("b"), []byte("vB")))
	heightBalanceTrunkTest("a", 1, -1)

	printErr(tr.Add([]byte("c"), []byte("vC")))
	heightBalanceTrunkTest("b", 1, 0)

	printErr(tr.Add([]byte("d"), []byte("vD")))
	heightBalanceTrunkTest("b", 2, 1)

	printErr(tr.Add([]byte("e"), []byte("vE")))
	heightBalanceTrunkTest("b", 2, 1)

	printErr(tr.Add([]byte("f"), []byte("vF")))
	heightBalanceTrunkTest("d", 2, 0)

	printErr(tr.Add([]byte("g"), []byte("vG")))
	heightBalanceTrunkTest("d", 2, 0)

	t.Log(tr.trunk.printStructure())

	//Test retrieving saved values
	retrieveTest("a", "vA", true)
	retrieveTest("b", "vB", true)
	retrieveTest("c", "vC", true)
	retrieveTest("d", "vD", true)
	retrieveTest("e", "vE", true)
	retrieveTest("f", "vF", true)
	retrieveTest("g", "vG", true)

	//Test adding a duplicate value
	expErr := tr.Add([]byte("a"), []byte("vA"))
	if expErr == nil {
		t.Errorf("expected to receive an error when attempting to add duplicate values")
	}

	//Test updating an existing value
	printErr(tr.Update([]byte("a"), []byte("vAA")))
	retrieveTest("a", "vAA", true)

	//Test updating a non existent value
	expErr = tr.Update([]byte("z"), []byte("vZ"))
	if expErr == nil {
		t.Errorf("expected to receive an error when attempting to update non-existent value")
	}

	//Test removing saved values from the tr
	//Expected tr structures:
	//      d         d      e
	//     /  \      / \    / \
	//    b    f    c   f  c   f
	//    \    /\      /
	//     c  e  g    e

	//removal of leafs
	printErr(tr.Remove([]byte("a")))
	printErr(tr.Remove([]byte("g")))
	heightBalanceTrunkTest("d", 2, 0)
	heightTest(2)
	t.Log(tr.trunk.printStructure())

	//removal of a branch with one child nodes
	printErr(tr.Remove([]byte("b")))
	heightTest(2)
	t.Log(tr.trunk.printStructure())

	//removal of a branch with one child nodes
	printErr(tr.Remove([]byte("d")))
	heightTest(1)
	t.Log(tr.trunk.printStructure())

	//Test retrieval of old saved value
	retrieveTest("a", "cA", false)
	retrieveTest("b", "vB", false)
	retrieveTest("c", "vC", true)
	retrieveTest("d", "vD", false)
	retrieveTest("e", "vE", true)
	retrieveTest("f", "vF", true)
	retrieveTest("g", "vG", false)
}
